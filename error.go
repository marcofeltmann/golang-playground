package main

import "fmt"

// DeferredError is an error type indicating that something failed in `defer`
type DeferredError string

// Error() implementes the error interface
func (e DeferredError) Error() string {
	return fmt.Sprintf("deferred error: %q", string(e))
}

// CommonError is an error type indicating some common failure.
type CommonError bool

// Error() implementes the error interface
func (e CommonError) Error() string {
	return fmt.Sprintf("common error: isDeferEnabled: %v", bool(e))
}

// Failure is a type that contains failing behavior.
// It's failing behavior can be configured via the FailsDefer and WrapsErr fields.
type Failure struct {
	FailsDefer bool
	WrapsErr   bool
}

// FailingFunc always returns an error.
// According to the configuration of the Failure type, those errors may vary.
func (f *Failure) FailingFunc(failsDefer bool) (err error) {
	if f.FailsDefer {
		defer func() { err = f.addError("first defer", err) }()
		defer func() { err = f.addError("second defer", err) }()
	}
	return CommonError(f.FailsDefer)
}

// addError returns an error.
// Depending on the configuration of Failure this error may wrap the oldErr.
func (f *Failure) addError(value string, oldErr error) error {
	err := DeferredError(value)
	if !f.WrapsErr {
		return err
	}
	// I'd love to wrap `oldErr` to the given `DeferredError` type…
	return fmt.Errorf("%q: %w", value, oldErr)
}
