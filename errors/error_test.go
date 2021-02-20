package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	customErrorCode = "CUSTOM_ERROR"
)

var (
	CustomError = Create(customErrorCode)
)

func TestError(t *testing.T) {
	suite.Run(t, &testErrorSuite{})
}

type testErrorSuite struct {
	suite.Suite
}

func (s *testErrorSuite) TestErrorData() {
	data := map[string]interface{}{
		"a": 123,
	}

	err, _ := Cause(
		CustomError("", Data(data)),
	).(*Error)

	s.Equal(data, err.Data())
}

func (s *testErrorSuite) TestErrorMessage() {
	err := CustomError("Hello World")

	s.Equal(
		err.Error(),
		fmt.Sprintf("[%s] Hello World", customErrorCode),
	)
	e := Cause(err).(*Error)

	s.Equal("Hello World", e.Message())
}

func (s *testErrorSuite) TestErrorHTTPStatusCode() {
	HTTPError := Create("HTTP_ERROR", HTTPStatusCode(401))

	err, _ := Cause(HTTPError("")).(*Error)

	s.Equal(401, err.HTTPStatusCode())
}

func (s *testErrorSuite) TestCheckErrorTypeByCode() {
	err := CustomError("")

	result := Is(err, customErrorCode)

	s.True(result)
}

func (s *testErrorSuite) TestCheckErrorTypeByFactory() {
	err := CustomError("")

	result := Is(err, CustomError)

	s.True(result)
}

func (s *testErrorSuite) TestStackTraceShouldStartFromFileWhichCreatesError() {
	err := CustomError("")

	frames := StackTrace(err)

	frame := fmt.Sprintf("%+v", frames[0])
	s.Contains(frame, "error_test.go", frame)
}
