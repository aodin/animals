package animals

import (
	"github.com/aodin/manager"
	"github.com/aodin/sol"
)

type animals struct {
	manager.Manager
}

// Get returns an Animal by ID
func (self animals) Get(id uint64) (animal Animal) {
	self.Query(self.Select().Where(self.C("id").Equals(id)).Limit(1), &animal)
	return
}

// GetByName returns a animal by name
func (self animals) GetByName(name string) (animal Animal) {
	stmt := self.Select().Where(self.C("name").Equals(name)).Limit(1)
	self.Query(stmt, &animal)
	return
}

func (self animals) Use(conn sol.Conn) animals {
	return animals{Manager: self.Manager.Use(conn)}
}

var Manager = animals{Manager: manager.New(Table)}
