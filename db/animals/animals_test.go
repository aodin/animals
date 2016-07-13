package animals

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aodin/animals/db/dbtest"
)

func TestAnimals(t *testing.T) {
	conn := dbtest.GetConn(t)
	tx, _ := conn.Must().Begin()
	defer tx.Rollback()

	dbtest.InitSchema(tx, Table)

	// Create a new Animal
	bird := New(" Bird ")
	assert.Nil(t, bird.Save(tx))
	assert.True(t, bird.Exists())
	assert.Equal(t, "Bird", bird.Name)

	// Use the manager
	bird2 := Manager.Use(tx).GetByName("Bird")
	assert.Equal(t, bird.Name, bird2.Name)
}
