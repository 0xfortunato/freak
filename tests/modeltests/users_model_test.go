package modeltests

import (
	"log"
	"testing"

	"github.com/0xfortunato/freak/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		ID:       1,
		Email:    "testing@gmail.com",
		Nickname: "testing",
		Password: "testingpasswd",
	}

	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("this is the error saving the new user: %v\n", err)
		return
	}

	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Password, savedUser.Password)
}

func TestGetUserByID(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf("error user refreshing table: %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v\n", err)
	}

	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed one user: %v\n", err)
	}

	userUpdate := models.User{
		ID:       1,
		Nickname: "modiUpdate",
		Email:    "modiupdate@gmail.com",
		Password: "modipasswd",
	}

	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating a user: %v\n", err)
		return
	}

	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)
}

func TestDeleteAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error deleting the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
