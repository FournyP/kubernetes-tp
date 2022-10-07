// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/FournyP/kubernetes-tp/apis/text/ent/schema"
	"github.com/FournyP/kubernetes-tp/apis/text/ent/text"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	textFields := schema.Text{}.Fields()
	_ = textFields
	// textDescName is the schema descriptor for name field.
	textDescName := textFields[0].Descriptor()
	// text.NameValidator is a validator for the "name" field. It is called by the builders before save.
	text.NameValidator = textDescName.Validators[0].(func(string) error)
}