package db

import (
	"jimu/src/db/postgre"
	"time"

	"github.com/uptrace/bun"
)

func init() {
	NewPgExample().CreateTable()
}

// https://bun.uptrace.dev/guide/models.html#table-names
type PgExample struct {
	postgre.DB[PgExample] `json:"-" bun:"-"`
	bun.BaseModel         `bun:"table:users,alias:u"`

	ID        int64  `bun:"id,pk,autoincrement"`
	Name      string `bun:"name,notnull"`
	email     string // unexported fields are ignored
	CreatedAt string
	UpdatedAt time.Time
}

func NewPgExample() *PgExample {
	return &PgExample{}
}
