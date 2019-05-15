# ◤◣ prefab

A tool to get prefabricated production ready code as a starter for your next adventure.
---------------------------------------------------------------------------------------

[![Build Status](https://travis-ci.com/yantrashala/prefab.svg?branch=master)](https://travis-ci.com/yantrashala/prefab)

## Prerequisites
* [git client](https://git-scm.com/)
* [Docker 17.05 or later](https://www.docker.com/)

> Optionally for local development
>* [go 1.11 or later](https://golang.org/dl)
>* [node v11 or later](https://nodejs.org)


## Steps to Install
If you want to use prefab as your app generator, simply install the prefab binaries. The prefab binaries have no external dependencies.

Finally, you can install the prefab source code with Go(programming language), build the binaries yourself, and run prefab that way. Building the binaries is an easy task for an experienced go getter.

### Install _prefab_ as Your Site Generator (Binary Install)

Use the installation [instructions in the prefab documentation](https://github.com/yantrashala/prefab/blob/master/README.md#the-prefab-documentation).

TBD: steps to get release binary for different plaform

### Build and Install the Binaries from Source (Advanced Install)

#### Prerequisite Tools

- [git client](https://git-scm.com/)
- [Go (tested with 1.12)](https://goland.org/dl)

#### Fetch from GitHub
Prefab uses the Go Modules support built into Go 1.12 to build. The easiest way is to clone prefab in a directory outside of GOPATH, as in the following example:

```
> mkdir $HOME/src
> cd $HOME/src
> git clone https://github.com/yantrashala/prefab.git
> cd prefab
> go run mage.go install
```

If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.

## The _prefab_ Documentation

TBD: links to documentation and tutorials
Docker Overview https://www.tutorialspoint.com/docker/docker_overview.htm

## Getting started for _prefab_ Developers

```
> git clone https://github.com/yantrashala/prefab.git
```

### Build and run in docker (suggested)

```
> cd prefab
> docker build -t ps/fab .
...
> docker run --rm -it -p9876:9876 ps/fab ./fab server
```

### or if you have optional prereqs installed locally try

```
> cd prefab
> go get ./
> cd ui
> npm install
> npm run build
> cd ..
> go get -d -v
> go test ./...
> go run main.go server
```

### or use the make file locally

```
> cd prefab
> make compile
> make start-server
...
...
> make stop-server
```

### or use [mage](https://github.com/magefile/mage) locally

```
> go get -d github.com/magefile/mage
> go run $GOPATH/src/github.com/magefile/mage/bootstrap.go install
> cd prefab
> mage prefab
>./bin/fab server

```

### or use chokidar for watch and recompile locally

first install chokidar globally (assuming node & npm are already installed)

```
> npm install -g chokidar-cli
```

```
> cd prefab
> make compile
> make watch run="make stop-server go-get go-build start-server"
...
```

## Contributing

To contribute to the prefab source code or documentation, you should fork the prefab GitHub project and clone it to your local machine.

1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/prefab && cd prefab)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request
8. Wait for it to get reviewed and merged

