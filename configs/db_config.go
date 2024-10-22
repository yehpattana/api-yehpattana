package configs

import (
	"fmt"
)

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (d *db) Url() string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s",
		d.username,
		d.password,
		d.host,
		d.port,
		d.database,
	)
}

func (d *db) MaxOpenConnections() int { return d.maxConnections }
