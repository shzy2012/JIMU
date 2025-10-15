package clickhouse

import (
	"context"
	"crypto/tls"
	"fmt"
	"jimu/src/config"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT name, toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name, uuid string
		if err := rows.Scan(&name, &uuid); err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, uuid: %s", name, uuid)
	}

}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: config.Clickhouse.Endpoints,
			Auth: clickhouse.Auth{
				Database: config.Clickhouse.Database,
				Username: config.Clickhouse.Username,
				Password: config.Clickhouse.Password,
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},
			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
