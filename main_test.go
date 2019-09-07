package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BlackWhole struct{}

func (b *BlackWhole) Write(p []byte) (n int, err error) {
	return 0, nil
}

func Test_getInput(t *testing.T) {
	defer os.Setenv("INPUT_FOO", "")

	assert.Panics(t, func() { getInput("foo") })
	assert.Equal(t, "lorem", getInput("foo", "lorem"))

	os.Setenv("INPUT_FOO", "bar")

	assert.Equal(t, "bar", getInput("foo"))
	assert.Equal(t, "bar", getInput("foo", "lorem"))
}

func Test_getInputEnum(t *testing.T) {
	defer os.Setenv("INPUT_FOO", "")

	assert.Panics(t, func() { getInputEnum("foo", []string{"bar"}) })
	assert.Panics(t, func() { getInputEnum("foo", []string{"bar"}, "thing") })
	assert.Equal(t, "thing", getInputEnum("foo", []string{"thing"}, "thing"))

	os.Setenv("INPUT_FOO", "bar")

	assert.Equal(t, "bar", getInputEnum("foo", []string{"bar"}))
	assert.Equal(t, "bar", getInputEnum("foo", []string{"lorem", "bar"}, "lorem"))
}

func Test_isJSON(t *testing.T) {
	assert.True(t, isJSON(`{}`), "empty object")
	assert.True(t, isJSON(`{ "foo": "bar" }`), "object")
	assert.True(t, isJSON(`"bar"`), "string")
	assert.True(t, isJSON(`[ "bar" ]`), "array")
}
