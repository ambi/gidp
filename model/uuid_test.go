package model

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	uuidPattern := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)

	uuid1 := NewUUID()
	uuid2 := NewUUID()

	assert.NotEqual(t, uuid1, uuid2)
	assert.Regexp(t, uuidPattern, uuid1)
	assert.Regexp(t, uuidPattern, uuid2)
}
