package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonError(t *testing.T) {

	json := jsonError("test")

	assert.Equal(t, "{\"message\":\"test\"}", string(json))
}
