// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MesuringPointsColumns holds the columns for the "mesuring_points" table.
	MesuringPointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "body_mass", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
	}
	// MesuringPointsTable holds the schema information for the "mesuring_points" table.
	MesuringPointsTable = &schema.Table{
		Name:       "mesuring_points",
		Columns:    MesuringPointsColumns,
		PrimaryKey: []*schema.Column{MesuringPointsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "mesuringpoint_user_id",
				Unique:  false,
				Columns: []*schema.Column{MesuringPointsColumns[1]},
			},
			{
				Name:    "mesuringpoint_created_at",
				Unique:  true,
				Columns: []*schema.Column{MesuringPointsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MesuringPointsTable,
		UsersTable,
	}
)

func init() {
}