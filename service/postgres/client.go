package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go-messanger/config"
)

const driverName = "postgres"

type Client struct {
	db *sqlx.DB
}

func New(cfg config.PostgresConfig) (*Client, error) {
	db, err := sqlx.Connect(driverName, cfg.Opts)

	if err != nil {
		return nil, errors.Wrap(err, "Cant connect to DB")
	}
	return &Client{db}, nil
}

func (c *Client) Update(ctx context.Context, query string, params ...any) error {
	_, err := c.db.ExecContext(ctx, query, params)
	return err
}

func (c *Client) GetOne(ctx context.Context, dest any, query string, params ...any) error {
	return c.db.GetContext(ctx, dest, query, params)
}

func (c *Client) Select(ctx context.Context, dest any, query string, params ...any) error {
	return c.db.SelectContext(ctx, dest, query, params)
}

func (c *Client) Close() error {
	return c.db.Close()
}
