// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname"
	// FieldAddress holds the string denoting the address vertex property in the database.
	FieldAddress = "address"
	// FieldRenamed holds the string denoting the renamed vertex property in the database.
	FieldRenamed = "renamed"
	// FieldBlob holds the string denoting the blob vertex property in the database.
	FieldBlob = "blob"
	// FieldState holds the string denoting the state vertex property in the database.
	FieldState = "state"

	// Table holds the table name of the user in the database.
	Table = "users"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// SpouseTable is the table the holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "user_spouse_id"
	// CarTable is the table the holds the car relation/edge.
	CarTable = "cars"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "owner_id"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldNickname,
	FieldAddress,
	FieldRenamed,
	FieldBlob,
	FieldState,
}

var (
	fields = schema.User{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[1].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)
)

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(state State) error {
	switch state {
	case StateLoggedIn, StateLoggedOut:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", state)
	}
}
