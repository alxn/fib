package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibs(t *testing.T) {
	for x := int64(0); x < 15; x++ {
		f := fib(x)
		assert.Equal(t, f, fibMemo(x))
		assert.Equal(t, f, fibDynamic(x))
	}
}
