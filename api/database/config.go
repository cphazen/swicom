package database

import "fmt"

var (
	dbUsername         = "postgres"
	dbPassword         = "postgres"
	dbName             = "postgres"
	dbHost             = "db"
	dbPort             = "5432"
	pgConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbPort, dbName)
)
