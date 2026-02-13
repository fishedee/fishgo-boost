package assert

import (
	"fmt"
	"runtime"
	"strings"
)

type InternalException struct {
	code    int
	message string
	stack   []string
	cause   interface{}
	isCrash bool
}

func NewInternalException(code int, message string, args ...interface{}) *InternalException {
	return newInternalException(2, nil, false, code, message, args...)
}

func newInternalException(stackBegin int, cause interface{}, isCrash bool, code int, message string, args ...interface{}) *InternalException {
	if len(args) != 0 {
		message = fmt.Sprintf(message, args...)
	}
	stack := []string{}
	for i := stackBegin; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, fmt.Sprintf("%s:%d", file, line))
	}

	return &InternalException{
		code:    code,
		message: message,
		stack:   stack,
		cause:   cause,
	}
}

func (this *InternalException) GetCode() int {
	return this.code
}

func (this *InternalException) GetMessage() string {
	return this.message
}

func (this *InternalException) GetCause() interface{} {
	return this.cause
}

func (this *InternalException) IsCrash() bool {
	return this.isCrash
}

func (this *InternalException) GetStackTrace() string {
	return strings.Join(this.stack, "\n")
}

func (this *InternalException) GetStackTraceLine(i int) string {
	return this.stack[i]
}

func (this *InternalException) Error() string {
	return fmt.Sprintf("[Code:%d] [Message:%s] [Stack:%s]", this.GetCode(), this.GetMessage(), this.GetStackTrace())
}

func InternalThrow(code int, message string, args ...interface{}) {
	exception := newInternalException(2, nil, false, code, message, args...)

	panic(exception)
}

func InternalCatchCrash(handler func(InternalException)) {
	err := recover()
	if err != nil {
		exception, isException := err.(*InternalException)
		if isException {
			handler(*exception)
		} else {
			exception := newInternalException(3, err, true, 1, fmt.Sprint(err))
			handler(*exception)
		}
	}
}

func InternalCatch(handler func(InternalException)) {
	err := recover()
	if err != nil {
		exception, isException := err.(*InternalException)
		if isException {
			if exception.IsCrash() == false {
				handler(*exception)
			} else {
				panic(exception)
			}
		} else {
			exception := newInternalException(3, err, true, 1, fmt.Sprint(err))
			panic(exception)
		}
	}
}
