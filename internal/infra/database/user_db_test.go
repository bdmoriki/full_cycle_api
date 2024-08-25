package database

import (
	"testing"

	"github.com/bdmoriki/full_cycle_api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := GetDBUser()
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("Bruno Moriki", "a@a.com", "123456")
	userDB := NewUser(db)
	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Email, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := GetDBUser()
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("Maria Silva", "maria@b.com", "123456789")
	userDB := NewUser(db)
	err = userDB.Create(user)

	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Email, userFound.Password)
}
