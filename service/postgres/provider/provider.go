package provider

import (
	"context"
	"github.com/pkg/errors"
	"go-messanger/dto"
	"go-messanger/service/encrypt"
	"go-messanger/service/postgres"
	"reflect"
)

var (
	ErrUsernameNotUnique = errors.New("username already exists")
	ErrEmailNotUnique    = errors.New("email already exists")
)

type AccountProvider struct {
	db *postgres.Client
}

func NewAccountProvider(client *postgres.Client) *AccountProvider {
	return &AccountProvider{client}
}

func (c *AccountProvider) RegisterNewUser(ctx context.Context, newUser dto.User) error {
	if err := c.validateNewUser(ctx, newUser); err != nil {
		return err
	}

	hashedPassword, err := encrypt.HashPassword(newUser.Password)
	if err != nil {
		return errors.Wrap(err, "Error hashing password")
	}
	newUser.Password = hashedPassword

	return c.db.Update(ctx, registerNewAccountQuery, newUser.Username, newUser.Email, newUser)
}

func (c *AccountProvider) validateNewUser(ctx context.Context, newUser dto.User) error {
	if err := c.validateUsernameIsUnique(ctx, newUser); err != nil {
		return err
	}
	if err := c.validateEmailIsUnique(ctx, newUser); err != nil {
		return err
	}
	return nil
}

func (c *AccountProvider) validateUsernameIsUnique(ctx context.Context, user dto.User) error {
	exist, err := c.validateNewUserFieldIsUnique(ctx, user, "Username", user.Username)
	if err != nil {
		return err
	}
	if exist {
		return ErrUsernameNotUnique
	}

	return nil
}

func (c *AccountProvider) validateEmailIsUnique(ctx context.Context, user dto.User) error {
	exist, err := c.validateNewUserFieldIsUnique(ctx, user, "Email", user.Email)
	if err != nil {
		return err
	}
	if exist {
		return ErrEmailNotUnique
	}

	return nil
}

func (c *AccountProvider) validateNewUserFieldIsUnique(ctx context.Context, user dto.User, fieldName, value string) (bool, error) {
	t := reflect.TypeOf(user)

	field, found := t.FieldByName(fieldName)
	if !found {
		return false, nil
	}
	dbField := field.Tag.Get("db")

	isExist, err := c.isFieldExist(ctx, dbField, value)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (c *AccountProvider) isFieldExist(ctx context.Context, field, value string) (bool, error) {
	var resp bool
	err := c.db.GetOne(ctx, &resp, isFieldUniqueQuery, field, value)
	if err != nil {
		return false, err
	}

	return resp, nil
}
