# Prefab documentation

## Table of contents

* Introduction
* Getting started
* Tutorials
* Reference

# ◤◣ prefab
A tool to get prefabricated production ready code as a starter for your next adventure.

[![Build Status](https://travis-ci.com/yantrashala/prefab.svg?branch=master)](https://travis-ci.com/yantrashala/prefab)
## Choose How to Install
If you want to use prefab as your app generator, simply install the prefab binaries. The prefab binaries have no external dependencies.

To contribute to the prefab source code or documentation, you should fork the prefab GitHub project and clone it to your local machine.

Finally, you can install the prefab source code with go, build the binaries yourself, and run prefab that way. Building the binaries is an easy task for an experienced go getter.

### Install *prefab* as Your Site Generator (Binary Install)
Use the installation [instructions in the prefab documentation]().
TBD: steps to get release binary for different plaform

### Build and Install the Binaries from Source (Advanced Install)

#### Prerequisite Tools
* [git client](https://git-scm.com/)
* [Go (tested with 1.12)](https://goland.org/dl)

#### Fetch from GitHub
Prefab uses the Go Modules support built into Go 1.12 to build. The easiest is to clone prefab in a directory outside of GOPATH, as in the following example:

```
> mkdir $HOME/src
> cd $HOME/src
> git clone https://github.com/yantrashala/prefab.git
> cd prefab
> go run mage.go install
```
If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.
