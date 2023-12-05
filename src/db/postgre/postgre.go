package postgre

import (
	"database/sql"
	"fmt"
	"jimu/src/config"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var pg *bun.DB

func PGClient() *bun.DB {
	if pg != nil {
		return pg
	}
	// https://bun.uptrace.dev/postgres/#pgdriver

	// dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	/*
		?sslmode=verify-full - enable TLS.
		?sslmode=disable - disables TLS.
		?dial_timeout=5s - timeout for establishing new connections.
		?read_timeout=5s - timeout for socket reads.
		?write_timeout=5s - timeout for socket writes.
		?timeout=5s - sets all three timeouts described above.
		?application_name=myapp - PostgreSQL application name.
	*/
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Postgres.Username, config.Postgres.Password,
		config.Postgres.Endpoint, config.Postgres.Database)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	pg = bun.NewDB(sqldb, pgdialect.New())
	err := pg.Ping()
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("Postgres DB Connected.")

	// Print all queries to stdout.
	pg.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return pg
}
