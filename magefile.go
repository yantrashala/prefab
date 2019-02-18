// +build mage

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	packageName  = "github.com/yantrashala/prefab"
	noGitLdflags = "-X $PACKAGE/common/prefab.buildDate=$BUILD_DATE"
)

var ldflags = "-X $PACKAGE/common/prefab.commitHash=$COMMIT_HASH -X $PACKAGE/common/prefab.buildDate=$BUILD_DATE"

// allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
var goexe = "go"

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}

	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
	// The default is "auto".
	os.Setenv("GO111MODULE", "on")
}

func flagEnv() map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return map[string]string{
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
	}
}

func isGoLatest() bool {
	return strings.Contains(runtime.Version(), "1.11")
}

func buildTags() string {
	// To build the extended Prefab version, build with
	// PRAFAB_BUILD_TAGS=extended mage install etc.
	if envtags := os.Getenv("PREFAB_BUILD_TAGS"); envtags != "" {
		return envtags
	}
	return "none"

}

// Build prefab binary
func Prefab() error {
	mg.Deps(Get)
	return sh.RunWith(flagEnv(), goexe, "build", "-ldflags", ldflags, "-tags", buildTags(), "-o", "./bin/fab", packageName)
}

// Build prefab without git info
func PrefabNoGitInfo() error {
	mg.Deps(Get)
	ldflags = noGitLdflags
	return Prefab()
}

var docker = sh.RunCmd("docker")

// Build prefab Docker container
func Docker() error {
	if err := docker("build", "-t", "prefab", "."); err != nil {
		return err
	}
	// yes ignore errors here
	docker("rm", "-f", "prefab-build")
	if err := docker("run", "--name", "prefab-build", "prefab server"); err != nil {
		return err
	}
	if err := docker("cp", "prefab-build:/go/bin/prefab", "."); err != nil {
		return err
	}
	return docker("rm", "prefab-build")
}

// Run tests and linters
func Check() {
	if strings.Contains(runtime.Version(), "1.8") {
		// Go 1.8 doesn't play along with go test ./... and /vendor.
		// We could fix that, but that would take time.
		fmt.Printf("Skip Check on %s\n", runtime.Version())
		return
	}

	//mg.Deps(Test386)

	mg.Deps(Fmt, Vet)

	// don't run two tests in parallel, they saturate the CPUs anyway, and running two
	// causes memory issues in CI.
	mg.Deps(TestRace)
}

// Run all the tests.
func Test() error {
	return sh.Run(goexe, "test", "./...", "-tags", buildTags())
}

// Run tests with race detector
func TestRace() error {
	return sh.Run(goexe, "test", "-race", "./...", "-tags", buildTags())
}

// Generate test coverage report
func TestCoverHTML() error {
	const (
		coverAll = "coverage-all.out"
		cover    = "coverage.out"
	)
	f, err := os.Create(coverAll)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte("mode: count")); err != nil {
		return err
	}
	pkgs, err := prefabPackages()
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		if err := sh.Run(goexe, "test", "-coverprofile="+cover, "-covermode=count", pkg); err != nil {
			return err
		}
		b, err := ioutil.ReadFile(cover)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		idx := bytes.Index(b, []byte{'\n'})
		b = b[idx+1:]
		if _, err := f.Write(b); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return sh.Run(goexe, "tool", "cover", "-html="+coverAll)
}

var (
	pkgPrefixLen = len("github.com/yantrashala/prefab")
	pkgs         []string
	pkgsInit     sync.Once
)

func prefabPackages() ([]string, error) {
	var err error
	pkgsInit.Do(func() {
		var s string
		s, err = sh.Output(goexe, "list", "./...")
		if err != nil {
			return
		}
		pkgs = strings.Split(s, "\n")
		for i := range pkgs {
			pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
		}
	})
	return pkgs, err
}

// Run golint linter
func Lint() error {
	pkgs, err := prefabPackages()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

//  Run go vet linter
func Vet() error {
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running go vet: %v", err)
	}
	return nil
}

// Run gofmt linter
func Fmt() error {
	if !isGoLatest() {
		return nil
	}
	pkgs, err := prefabPackages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Prefab

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Prefab)
	fmt.Println("Installing...")
	return os.Rename("./bin/fab", "/usr/local/bin/fab")
}

// A step to get and install UI
func UIGet() {
	fmt.Println("Installing UI Deps...")
	cmd := exec.Command("npm", "install")
	return cmd.Run()
}

// Get the dependent modules.
func Get() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "-d", "-v")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("bin")
}
