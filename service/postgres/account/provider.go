package account

import (
	"context"
	"github.com/pkg/errors"
	"go-messanger/dto"
	"go-messanger/service/encrypt"
	"go-messanger/service/postgres"
	"reflect"
)

type Provider struct {
	db *postgres.Client
}

func New(client *postgres.Client) *Provider {
	return &Provider{client}
}

func (c *Provider) RegisterNewUser(ctx context.Context, newUser *dto.User) error {
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

func (c *Provider) validateNewUser(ctx context.Context, newUser *dto.User) error {
	if err := c.validateNewUserFieldIsUnique(ctx, newUser, "Username", newUser.Username); err != nil {
		return err
	}
	if err := c.validateNewUserFieldIsUnique(ctx, newUser, "Email", newUser.Email); err != nil {
		return err
	}
	return nil
}

func (c *Provider) validateNewUserFieldIsUnique(ctx context.Context, user *dto.User, fieldName, value string) error {
	t := reflect.TypeOf(user)

	field, found := t.FieldByName(fieldName)
	if !found {
		return nil
	}
	dbField := field.Tag.Get("db")

	isExist, err := c.IsFieldExist(ctx, dbField, value)
	if err != nil {
		return err
	}

	if isExist {
		return errors.Errorf("%s=%s already exist in database", fieldName, value)
	}
	return nil
}

func (c *Provider) IsFieldExist(ctx context.Context, field, value string) (bool, error) {
	var resp bool
	err := c.db.GetOne(ctx, &resp, isFieldUniqueQuery, field, value)
	if err != nil {
		return false, err
	}

	return resp, nil
}
