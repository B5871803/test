package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Operativerecord holds the schema definition for the Operativerecord entity.
type Operativerecord struct {
	ent.Schema
}

// Fields of the Operativerecord.
func (Operativerecord) Fields() []ent.Field {
	return []ent.Field{

		field.Time("OperativeTime").Default(time.Now),
	}

}

// Edges of the Operativerecord.
func (Operativerecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("examinationroom", Examinationroom.Type).Ref("fromexaminationroom").Unique(),
		edge.From("nurse", Nurse.Type).Ref("nurseop").Unique(),
		edge.From("operative", Operative.Type).Ref("fromoperative").Unique(),
		edge.From("tool", Tool.Type).Ref("fromtool").Unique(),
	}
}
