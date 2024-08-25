package database

import (
	"testing"

	"github.com/bdmoriki/full_cycle_api/internal/entity"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

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
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

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
