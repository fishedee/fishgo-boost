package assert

import (
	"testing"
)

func TestAssertException(t *testing.T) {
	//below test case should fail!
	AssertException(t, 1, "", func() {
		InternalThrow(2, "")
	})
	AssertException(t, 1, "123", func() {
		InternalThrow(1, "456")
	})
	AssertException(t, 1, "", func() {
	})
}

func TestAssertPanic(t *testing.T) {
	//below test case should fail!
	AssertException(t, 1, "123", func() {
		panic("456")
	})

	AssertException(t, 1, "", func() {
	})
}
