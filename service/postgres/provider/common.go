package provider

const (
	registerNewAccountQuery = `insert into accounts(username, email, "password", registration_date)
values($1, $2, $3,  extract (epoch from now()) :: bigint);`
	isFieldUniqueQuery = `select exists (
	select $1 from accounts a where $1 = $2)`
)
