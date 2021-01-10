package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Operative holds the schema definition for the Operative entity.
type Operative struct {
	ent.Schema
}

// Fields of the Operative.
func (Operative) Fields() []ent.Field {
	return []ent.Field{

		field.String("operativeType").NotEmpty(),
		field.String("operativeName").NotEmpty().Unique(),
	}
}

// Edges of the Operative.
func (Operative) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fromoperative", Operativerecord.Type).StorageKey(edge.Column("operative_type")),
	}
}
