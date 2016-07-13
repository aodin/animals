package animals

import (
	"github.com/aodin/fields"
	"github.com/aodin/sol"
	"github.com/aodin/sol/postgres"
	"github.com/aodin/sol/types"
)

var Table = postgres.Table("animals",
	fields.Serial{},
	sol.Column("name", types.Varchar().NotNull()),
	sol.Column(
		"modified_at",
		postgres.Timestamp().NotNull().Default(postgres.NowUTC),
	),
	sol.PrimaryKey("id"),
	sol.Unique("name"),
)
