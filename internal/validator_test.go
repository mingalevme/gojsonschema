package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidateValidDocument(t *testing.T) {
	result, err := Validate("./../test/testdata/schema.json", "./../test/testdata/document-valid.json")
	assert.NoError(t, err)
	assert.True(t, result.Valid())
}

func TestValidateInvalidDocument(t *testing.T) {
	result, err := Validate("./../test/testdata/schema.json", "./../test/testdata/document-invalid.json")
	assert.NoError(t, err)
	assert.False(t, result.Valid())
}

func TestAbsolutizeUrl(t *testing.T) {
	var res string
	var err error
	//
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	//
	res, err = AbsolutizeUrl("file:///schema.json")
	assert.NoError(t, err)
	assert.Equal(t, "file:///schema.json", res)
	//
	res, err = AbsolutizeUrl("document.json")
	assert.NoError(t, err)
	assert.Equal(t, "file://"+cwd+"/document.json", res)
	//
	res, err = AbsolutizeUrl("./document.json")
	assert.NoError(t, err)
	assert.Equal(t, "file://"+cwd+"/document.json", res)
	//
	res, err = AbsolutizeUrl("example.com")
	assert.NoError(t, err)
	assert.Equal(t, "file://"+cwd+"/example.com", res)
	//
	res, err = AbsolutizeUrl("https://example.com")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", res)
	//
	res, err = AbsolutizeUrl("//example.com")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", res)
	//
	res, err = AbsolutizeUrl("//example.com/")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/", res)
	//
	res, err = AbsolutizeUrl("//example.com/path.json")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/path.json", res)
}
