package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Bruno Moriki", "a@a.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Bruno Moriki", user.Name)
	assert.Equal(t, "a@a.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Bruno Moriki", "a@a.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
