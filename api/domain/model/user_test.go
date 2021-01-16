package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	var id int = 1
	var name string = "anonymous"
	var createdAt time.Time = time.Now()
	var updatedAt time.Time = time.Now()
	user := NewUser(id, name, createdAt, updatedAt)

	if user == nil {
		t.Errorf("failed NewUser()")
	}

	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, createdAt, user.CreatedAt)
	assert.Equal(t, updatedAt, user.UpdatedAt)
}
