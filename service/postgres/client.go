package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go-messanger/config"
)

const driverName = "postgres"

type Client struct {
	db *sqlx.DB
}

func New(cfg *config.PostgresConfig) (*Client, error) {
	db, err := sqlx.Connect(driverName, cfg.Options)

	if err != nil {
		return nil, errors.Wrap(err, "Cant connect to DB")
	}
	return &Client{db}, nil
}

func (c *Client) Update(query string, params ...any) error {
	_, err := c.db.Exec(query, params)
	return err
}

func (c *Client) Get(dest any, query string, params ...any) error {
	return c.db.Get(dest, query, params)
}

func (c *Client) Select(dest any, query string, params ...any) error {
	return c.db.Select(dest, query, params)
}
