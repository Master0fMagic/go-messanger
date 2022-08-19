package account

import "go-messanger/service/postgres"

type Provider struct {
	db *postgres.Client
}

func New(client *postgres.Client) *Provider {
	return &Provider{client}
}
