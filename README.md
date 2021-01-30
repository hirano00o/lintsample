# lintsample

This is a sample of making my own golang linter.

# Usage

```console
# for Linux/Mac
make build
go vet -vettool=lintsample testdata/src/a\a.go

# for Windows
make build
go vet -vettool=lintsample.exe testdata\src\a\a.go
```
