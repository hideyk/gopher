Since Go 1.11, modules have allowed us to manage our projects outsides of the `$GOPATH/src` folder - a big benefit, especially where a project contains more than just Go source code, as many projects organised as mono-repositories frequently do.

## 4.1 Internal folder
The `internal` folder gets special treatment from the Go tool which uses it to limit access to packages contained within it. Effectively makes code invisible to another package unless it shares a common ancestor.

Hide entire packages which are for *internal* use only. 

## 4.2 Cmd folder
What if you have multiple applications which share much of the same code? For example, you have a HTTP rest-based API and also CLI which both use the same database access routines and wrapper logic.

Simply move the executables (anything with `package main` into their own subfolder within a `cmd` folder).

Using thie approach we can manage multiple executables in one project and provide access to the common libraries they use without needing to put them in a third project.


## Pkg folder
First the name `pkg` is semantic and tells other developers what packages in our module are intended to be reusable outside of the module. 
Most projects will have numerous other folders which reside at root level. Folders for config, build automation, data folders.
Helps locating all Go packages in the `pkg` folder.

```
cmd/
    api/
        main.go
    cli/
        main.go
pkg/
    database/
        *.go
    logic/
        *.go
internal/
    hidden-package/
        *.go
go.mod
go.sum
```