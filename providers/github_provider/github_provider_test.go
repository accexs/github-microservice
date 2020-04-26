package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}