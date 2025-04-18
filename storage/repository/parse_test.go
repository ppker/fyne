package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileURI(t *testing.T) {
	assert.Equal(t, "file:///tmp/foo.txt", NewFileURI("/tmp/foo.txt").String())
	assert.Equal(t, "file://C:/tmp/foo.txt", NewFileURI("C:/tmp/foo.txt").String())
}

func TestParseURI(t *testing.T) {
	uri, err := ParseURI("file:///tmp/foo.txt")
	assert.NoError(t, err)
	assert.Equal(t, "file:///tmp/foo.txt", uri.String())

	uri, err = ParseURI("file:/tmp/foo.txt")
	assert.NoError(t, err)
	assert.Equal(t, "file:///tmp/foo.txt", uri.String())

	uri, err = ParseURI("file://C:/tmp/foo.txt")
	assert.NoError(t, err)
	assert.Equal(t, "file://C:/tmp/foo.txt", uri.String())
}

func TestParseInvalidURI(t *testing.T) {
	uri, err := ParseURI("/tmp/foo.txt")
	assert.Error(t, err)
	assert.Nil(t, uri)

	uri, err = ParseURI("file")
	assert.Error(t, err)
	assert.Nil(t, uri)

	uri, err = ParseURI("file:")
	assert.Error(t, err)
	assert.Nil(t, uri)

	uri, err = ParseURI("file://")
	assert.Error(t, err)
	assert.Nil(t, uri)

	uri, err = ParseURI(":foo")
	assert.Error(t, err)
	assert.Nil(t, uri)

	uri, err = ParseURI("scheme://0[]/invalid")
	assert.Error(t, err)
	assert.Nil(t, uri)
}
