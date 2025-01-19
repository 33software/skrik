package tests

import (
	"errors"
	"skrik/internal/entities"
	"skrik/internal/repository"
	"testing"

	"gorm.io/gorm"
)

func TestFindUserById(t *testing.T) {
	user := &entities.User{Username: "testuser", Password: "123"}
	TestDb.Create(&user)
	repo := repository.NewUserRepository(TestDb)

	foundUser, err := repo.FindUserById(user.ID)
	if err != nil {
		t.Errorf("expected no err. got err: %v", err)
	}
	if foundUser.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %s", foundUser.Username)
	}
	_, err = repo.FindUserById(99999)
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected record not found error; got %v", err)
	}

}
