// Code generated by ent, DO NOT EDIT.

package mesuringpoint

import (
	"time"
)

const (
	// Label holds the string label denoting the mesuringpoint type in the database.
	Label = "mesuring_point"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldBodyMass holds the string denoting the body_mass field in the database.
	FieldBodyMass = "body_mass"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the mesuringpoint in the database.
	Table = "mesuring_points"
)

// Columns holds all SQL columns for mesuringpoint fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldBodyMass,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(int64) error
	// BodyMassValidator is a validator for the "body_mass" field. It is called by the builders before save.
	BodyMassValidator func(float64) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)