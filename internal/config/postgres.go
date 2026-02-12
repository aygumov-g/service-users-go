package config

import "fmt"

type postgres struct {
	dBHost     string
	dBUser     string
	dBPassword string
	dBName     string
}

func (p postgres) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		p.dBUser,
		p.dBPassword,
		p.dBHost,
		p.dBName)
}
