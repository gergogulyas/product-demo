# ProductDemo

## How to

### Start HTTP server
```
go run main.go
```

### Access

After HTTP server has been started it can be reached on `localhost:8090`.

For example:
`curl --location --request GET 'localhost:8090/products'`

### Run test

```
go test -v ./...
```

#### Prerequisite
Please install dependencies before running tests with `go mod vendor` command.
