## 2.1 Version Information
- `go version`

## 2.2 Environment Information
- `go env`

## 2.3 Format your code
- `go fmt main.go`

## 2.4 Testing
- `go test ./...`

## 2.5 Cleanup
- `go clean` (Removes object files from package source directories)

## 2.6 Downloading packages
- `go get example.com/pkg` (Download package dependencies)

## 2.7 Running a program
- `go run main.go`

## 2.8 Building your program
- `go build -o helloWorld ./...` (-o Specific name for executable)

## 2.9 Building for other operating systems and architectures
- `go tool dist list`
- `env GOOS=linux GOARCH=amd64 gob uild -o helloWorld ./...`

Variables:
- GOPATH (Location of go src code including standard library)
- GOMODCACHE (Location of modules that were downloaded and cached)
- GOPROXY (Proxy from which new packages are downloaded)
- GOOS (Go Operating System)
- GOARCH (Go Architecture)