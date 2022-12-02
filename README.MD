# Go-Unit-Test-Workshop

## Workshop cover

```
- Testing is a critical part of the development process
- Well tested code is more maintainable, more secure, and overall
  takes less time to write
- Go has a rich set of tools for testing build right in
- We won't use any third party testing package at all (e.g. Testify)
- We'll write unit test admin integration tests
```

## Run test show coverage

```bash
$ go test ./... -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```

![My Image](https://github.com/yuttasakcom/go-unit-test-workshop/blob/master/images/result-test.png)
