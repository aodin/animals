package dbtest

import (
	"os"
	"sync"
	"testing"

	"github.com/aodin/sol"
	_ "github.com/aodin/sol/postgres"
)

var testconn *sol.DB
var once sync.Once

// GetConn returns a connection pool. It should only be used for testing.
func GetConn(t *testing.T) *sol.DB {
	credentials := os.Getenv("ANIMALS_TEST")
	if credentials == "" {
		t.Fatalf("No testing database credentials set")
	}

	once.Do(func() {
		var err error
		if testconn, err = sol.Open("postgres", credentials); err != nil {
			t.Fatalf("Failed to open connection: %s", err)
		}
		testconn.SetMaxOpenConns(25)
	})
	return testconn
}

func InitSchema(conn sol.Conn, tables ...sol.Tabular) {
	// Create the given schemas
	for _, table := range tables {
		if table == nil || table.Table() == nil {
			continue
		}
		conn.Query(table.Table().Create().IfNotExists().Temporary())
	}
}
