package main

import (
	"errors"
	"fmt"
	"testing"
)

type tRow struct {
	sut Failure
}

var tt []tRow = []tRow{
	{sut: Failure{false, false}},
	{sut: Failure{false, true}},
	{sut: Failure{true, false}},
	{sut: Failure{true, true}},
}

func TestErrorFunc(t *testing.T) {
	for _, r := range tt {
		t.Run(fmt.Sprintf("%#v", r.sut),
			func(t *testing.T) {
				errorResults(r.sut, t)
			})
	}
}

func errorResults(sut Failure, t *testing.T) {
	failsDefer := sut.FailsDefer
	err := sut.FailingFunc(failsDefer)

	if err == nil {
		t.Errorf("got success, want FailingFunc(%v) fail", failsDefer)
	}
	if !errors.Is(err, CommonError(failsDefer)) {
		t.Errorf(" got %v, want CommonError(%v)", err, failsDefer)
	}

	if !failsDefer {
		return
	}

	deferWrapsErr := sut.WrapsErr

	if !deferWrapsErr {
		return
	}

	if !errors.Is(err, DeferredError("second defer")) {
		t.Errorf("want %v to contain DeferredError('second defer')", err)
	}
	if !errors.Is(err, DeferredError("first defer")) {
		t.Errorf("want %v to contain DeferredError('first defer')", err)
	}
}
