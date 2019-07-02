package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	assert.NotEqual(t, GenerateID(), GenerateID())
}