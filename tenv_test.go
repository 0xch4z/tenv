package tenv

import (
	"math/rand"
	"os"
	"testing"
)

func randStr() string {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}

func TestSet(t *testing.T) {
	someKey := randStr()
	originalValue := randStr()

	os.Setenv(someKey, originalValue)
	c := New()

	newValue := "newValue"
	c.Set(someKey, newValue)

	maskedEnvVal := os.Getenv(someKey)
	if maskedEnvVal != "newValue" {
		t.Errorf(`expected Set to set temporary env variable to "%s"; got "%s"`, newValue, maskedEnvVal)
	}

	c.Restore()

	restoredEnvVal := os.Getenv(someKey)
	if restoredEnvVal != originalValue {
		t.Errorf(`expected variable to be restored to "%s"; got "%s"`, originalValue, restoredEnvVal)
	}
}

func TestUnset(t *testing.T) {
	someKey := randStr()
	originalValue := randStr()

	os.Setenv(someKey, originalValue)
	c := New()

	c.Unset(someKey)

	maskedEnvVal := os.Getenv(someKey)
	if maskedEnvVal != "" {
		t.Errorf(`expected Set to set temporary env variable to be unset; got "%s"`, maskedEnvVal)
	}

	c.Restore()

	restoredEnvVal := os.Getenv(someKey)
	if restoredEnvVal != originalValue {
		t.Errorf(`expected variable to be restored to "%s"; got "%s"`, originalValue, restoredEnvVal)
	}
}
