#!/bin/sh

# My location on an Apple Mac M1 with `go` installed via `homebrew`
# Your setting may differ so feel free to change this

GO_EXEC="/opt/homebrew/bin/go"

# Very simple test case structure, I know.
if $GO_EXEC run flags.go -foo baz | grep -q "foo: baz"; then
	echo "'-foo baz' translates to 'foo: baz'"
else
	echo "'-foo baz' seems broken!"
fi

if $GO_EXEC run flags.go --foo=qux | grep -q "foo: qux"; then
	echo "'--foo=qux' translates to 'foo: qux'"
else 
	echo "'--foo=qux' seems broken!"
fi

