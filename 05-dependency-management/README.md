## 5.1 Modules
Modules introduced as an experimental feature in Go version 1.11, then became defactor approach for dependency management in version 1.13.

A module comprises one or more Go packages managed by a `go.mod` file. This file defines the module's identifier, the minimum version of Go required, and any external modules with versions.

Module's identifier is also the import path for the module. Remote modules need to be resolvable by that path modifier. 

If external dependencies are present, additional `go.sum` file contains checksums of dependencies and is used to fix versions and determine where dependencies have been modified.

### 5.1.1 Direct vs indirect dependencies
Indirect dependency is a dependency requirement of one of the direct dependencies or a dependency listed in `go.mod` which is not currently used in any of the module source files.

Indirect dependencies in `go.mod` are suffixed with `// indirect`

```go.mod
module github.com/my-module

go 1.17

require (
    github.com/spoonboy-io/koan v0.1.0
    github.com/TwiN/go-color v1.1.0 // indirect
    golang.org/x/sys v0,0,0-20210630005230-0f9fa26af87c // indirect
)
```

### 5.1.2 Creating and updating a module
Module creation is straightforward via `go mod init`

```
go mod init *module-identifier*
```

Remote modules included in your code with the `import` keyword, and are not already available on your system, can be downloaded and added to `go.mod` with `go mod tidy`

```
go mod tidy
```

Alternatively, use the traditional `go get` command from within the project workspace to fetch the module. The command is module-aware. 

```
go get *module-identifier*@v1.0.0
```

Use *tag, branch or commit* reference to use non-latest version.

### 5.1.3 Replace directive

The `replace` directive provides a simple mechanism to substitute a required module with another version of that module. This code can either be stored locally on your machine or at an alternative remote URL. 

For example, use a modified code repository over the original buggy codebase. 

```
go mod edit -replace=old[@v]=new[@v]
go mod edit -dropreplace=old[@v]
```

## 5.2 Workspaces
Added in Go version 1.18 and furthered the flexibility of modules and dependency management.

*Workspaces* allow replacements of modules to be made with modules found locally, without making changes to the `go.mod` file. Avoids the problem of having to add and subsequently remove replace directives, when working with local module versions.

Exclude from version control.

```
// create a workspace file
go work init

// or create the file and add a local path to the workspace
go work init <path/to/modules>
```

## 5.3 Vendoring
Vendoring includes a local copy of external module dependencies inside the project itself in a `vendor` folder. Go then uses the vendored modules in builds and testing instead of external equivalents.

Useful on machines with no internet access; or where strict dependency management has been adopted and submitting vendored modules to version control helps with that. 

Build times can be reduced for automated builds in a CI/CD process. 

`go mod vendor`

A `vendor` folder will be created with local copies of all external module dependencies. A `modules.txt` file which lists vendored modules and their versions will be created too. 


