package playground_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/marcofeltmann/playground"
)

type tRow struct {
	sut playground.Failure
}

var tt []tRow = []tRow{
	{sut: playground.Failure{false, false}},
	{sut: playground.Failure{false, true}},
	{sut: playground.Failure{true, false}},
	{sut: playground.Failure{true, true}},
}

// TestErrorFunc tests the error behavior of the Failure's FailingFunc.
func TestErrorFunc(t *testing.T) {
	for _, r := range tt {
		// One subtest per tRow means better readability of the failures with `go test -v`
		t.Run(fmt.Sprintf("%#v", r.sut),
			func(t *testing.T) {
				errorResults(r.sut, t)
			})
	}
}

// errorResults runs a complete test over the given "Failure in testing".
func errorResults(sut playground.Failure, t *testing.T) {
	failsDefer := sut.FailsDefer
	err := sut.FailingFunc(failsDefer)
	if err == nil {
		t.Errorf("got success, want FailingFunc(%v) fail", failsDefer)
	}

	if !failsDefer {
		checkErrNoDefers(err, t)
		return
	}

	deferWrapsErr := sut.WrapsErr

	if !deferWrapsErr {
		checkErrNoWraps(err, t)
		return
	}

	checkErrWraps(err, t)
}

func checkErrNoDefers(err error, t *testing.T) {
	if !errors.Is(err, playground.CommonError(false)) {
		t.Errorf(" got %v, want CommonError(false)", err)
	}
}

func checkErrNoWraps(err error, t *testing.T) {
	if !errors.Is(err, playground.DeferredError{Msg: "first defer"}) {
		t.Errorf(" got %v, want DeferredError(first defer)", err)
	}
}

func checkErrWraps(err error, t *testing.T) {
	if !errors.Is(err, playground.CommonError(true)) {
		t.Errorf(" got %v, want CommonError(true)", err)
	}
	if !errors.Is(err, playground.DeferredError{Msg: "second defer"}) {
		t.Errorf("want %v to contain DeferredError('second defer')", err)
	}
	if !errors.Is(err, playground.DeferredError{Msg: "first defer"}) {
		t.Errorf("want %v to contain DeferredError('first defer')", err)
	}
}
