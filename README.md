## got - go test file generator

I found myself always copy-pastaing my common test functions.
So I made this little thing to save me whole seconds of time.

### Installation

Get a [released binary](https://github.com/warehouse-13/got/releases), or

```bash
go install github.com/warehouse-13/got@latest
```

### Usage

```
got ./got -h
Usage of ./got:
  -name string
        The name of the test file (without the '_test'). Default: current package name.
```

`cd` into the target directory where you want a test file.

Examples:

```
$ cd pkg/client
$ got
Dummy test file written to client_test.go.
Run `go mod tidy` to start using.
```

```
$ cd pkg/validation
$ got -name config
Dummy test file written to config_test.go.
Run `go mod tidy` to start using.
```
