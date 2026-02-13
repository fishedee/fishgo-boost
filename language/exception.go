package language

import (
	"github.com/fishedee/fishgo-boost/assert"
)

type Exception = assert.InternalException

var NewException = assert.NewInternalException

var Throw = assert.InternalThrow

var CatchCrash = assert.InternalCatchCrash

var Catch = assert.InternalCatch
