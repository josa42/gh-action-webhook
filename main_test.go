package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isJSON(t *testing.T) {
	assert.True(t, isJSON(`{}`), "empty object")
	assert.True(t, isJSON(`{ "foo": "bar" }`), "object")
	assert.True(t, isJSON(`"bar"`), "string")
	assert.True(t, isJSON(`[ "bar" ]`), "array")
}
