package assert

import (
	"strings"
	"testing"
)

func getCatchMessage(fun func()) (_last string) {
	defer InternalCatch(func(e InternalException) {
		_last = e.GetMessage()
	})
	fun()
	return ""
}

func getCatchCrashMessage(fun func()) (_last string) {
	defer InternalCatchCrash(func(e InternalException) {
		_last = e.GetMessage()
	})
	fun()
	return ""
}

type errorStruct struct {
}

func (this *errorStruct) Error() string {
	return "m2"
}

func TestCatch(t *testing.T) {
	testCase := []struct {
		origin func()
		target string
	}{
		{func() {
			panic("m1")
		}, "m1"},
		{func() {
			panic(&errorStruct{})
		}, "m2"},
		{func() {
			InternalThrow(1, "m3")
		}, "m3"},
	}

	for singleIndex, singleTestCase := range testCase {
		msg := getCatchCrashMessage(singleTestCase.origin)
		AssertEqual(t, msg, singleTestCase.target, singleIndex)
	}
}

func explode(input string, separator string) []string {
	dataResult := strings.Split(input, separator)
	dataResultNew := make([]string, 0, len(dataResult))
	for _, singleResult := range dataResult {
		singleResult = strings.Trim(singleResult, " ")
		if len(singleResult) == 0 {
			continue
		}
		dataResultNew = append(dataResultNew, singleResult)
	}
	return dataResultNew
}

func getLastStackTraceLine(e InternalException) string {
	lines := explode(e.GetStackTraceLine(0), "/")
	return lines[len(lines)-1]
}

func TestCatchStack1(t *testing.T) {
	defer InternalCatch(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:62")
	})
	InternalThrow(1, "test1")
}

func TestCatchStack2(t *testing.T) {
	defer InternalCatchCrash(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:69")
	})
	InternalThrow(1, "test2")
}

func TestCatchStack3(t *testing.T) {
	defer InternalCatchCrash(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:76")
	})
	panic("test3")
}

func TestCatchStack4(t *testing.T) {
	defer InternalCatchCrash(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:86")
	})
	defer InternalCatch(func(e InternalException) {
		AssertEqual(t, "should not be here!", false)
	})
	panic("test4")
}

func TestCatchStack5(t *testing.T) {
	defer InternalCatchCrash(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:96")
	})
	defer InternalCatch(func(e InternalException) {
		panic(&e)
	})
	InternalThrow(1, "test5")
}

func TestCatchStack6(t *testing.T) {
	defer InternalCatch(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:106")
	})
	defer InternalCatch(func(e InternalException) {
		panic(&e)
	})
	InternalThrow(1, "test6")
}

func TestCatchStack7(t *testing.T) {
	defer InternalCatchCrash(func(e InternalException) {
		AssertEqual(t, getLastStackTraceLine(e), "exception_test.go:116")
	})
	defer InternalCatchCrash(func(e InternalException) {
		panic(&e)
	})
	panic("test7")
}
