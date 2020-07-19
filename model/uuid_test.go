package model

import (
	"regexp"
	"testing"
)

var uuidPattern = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)

func TestNewUUID(t *testing.T) {
	uuid1 := NewUUID()
	uuid2 := NewUUID()

	if uuid1 == uuid2 {
		t.Errorf("NewUUID() should generate a new ID, but got the same ID (%s)", uuid1)
	}

	if !uuidPattern.MatchString(uuid1) {
		t.Errorf("NewUUID() format should be UUIDv4, but got %s", uuid1)
	}

	if !uuidPattern.MatchString(uuid2) {
		t.Errorf("NewUUID() format should be UUIDv4, but got %s", uuid2)
	}
}
