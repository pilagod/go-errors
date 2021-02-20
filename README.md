# go-errors [![Build Status](https://travis-ci.com/pilagod/go-errors.svg?branch=master)](https://travis-ci.com/pilagod/go-errors) [![Coverage Status](https://coveralls.io/repos/github/pilagod/go-errors/badge.svg?branch=master)](https://coveralls.io/github/pilagod/go-errors?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/pilagod/go-errors)](https://goreportcard.com/report/github.com/pilagod/go-errors)

Error util for Go

## Installation

```shell
$ go get -u github.com/pilagod/go-errors
```

## Usage

Import `errors` from `go-errors`:

```go
import (
    "github.com/pilagod/go-errors/errors"
)
```

Create error factory with specific code:

```go
var (
    AuthenticationError = errors.Create("AUTHENTICATION_ERROR", errors.HTTPStatusCode(401))
)
```

You can use the error factory to create specific error:

```go
func CheckAuthentication(token string) error {
    if !isValid(token) {
        return AuthenticationError("Access token invalid.")
    }
    return nil
}
```

Handle error with `Cause` and `StackTrace`:

```go
func ErrorHandler(err error) {
    // use `StackTrace` to log stack trace
    fmt.Printf("%+v", errors.StackTrace(err))

    // use `Cause` to get underlying error
    switch v := errors.Cause(err).(type) {
    case *errors.Error:
        return map[string]interface{}{
            "status": v.HTTPStatusCode(),
            "code": v.Code(),
            "message": v.Message(),
            "data": v.Data()
        }
    default
        return map[string]interface{}{
            "status": 500,
            "code": "INTERNAL_SERVER_ERROR",
            "message": "Internal server error"
        }
    }
}
```

## License

© Cyan Ho (pilagod), 2021-NOW

Released under the [MIT License](https://github.com/pilagod/go-errors/blob/master/LICENSE)
