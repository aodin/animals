package animals

import (
	"strings"
	"time"

	"github.com/aodin/errors"
	"github.com/aodin/fields"
	"github.com/aodin/sol"
)

// Animals is a animal
type Animal struct {
	fields.Serial
	Name       string    `json:"name"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (animal Animal) Error(conn sol.Conn) *errors.Error {
	return nil
}

// Save persists the Animal to the database
func (animal *Animal) Save(conn sol.Conn) error {
	return conn.Query(Table.Insert().Values(animal).Returning(), animal)
}

// New creates a new Animal. It does not persist to the database.
func New(name string) (animal Animal) {
	animal.Name = strings.TrimSpace(name)
	animal.ModifiedAt = time.Now().UTC()
	return
}
