#!/bin/sh

if /opt/homebrew/bin/go run flags.go -foo baz | grep -q "foo: baz"; then
	echo "'-foo baz' translates to 'foo: baz'"
else
	echo "'-foo baz' seems broken!"
fi

if /opt/homebrew/bin/go run flags.go --foo=qux | grep -q "foo: qux"; then
	echo "'--foo=qux' translates to 'foo: qux'"
else 
	echo "'--foo=qux' seems broken!"
fi

