package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MesuringPoint struct {
	ent.Schema
}

func (MesuringPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Positive(),
		field.Float("body_mass").Positive(),
		field.Time("created_at").Default(time.Now),
	}

}

func (MesuringPoint) Edges() []ent.Edge {
	return nil
}

func (MesuringPoint) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("created_at").Unique(),
	}
}
