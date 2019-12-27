// Package object implements the objects system for the language.
package object

import "fmt"

// T represents the type of the object.
type T string

// Object types.
const (
	BooleanT     = "BOOLEAN"
	IntegerT     = "INTEGER"
	NullT        = "NULL"
	ReturnValueT = "RETURN_VALUE"
)

// Object represents an object or value type within the language.
type Object interface {
	Type() T
	Inspect() string
}

// Integer represents an object of that type.
type Integer struct {
	Value int64
}

// Inspect returns a representation of the object value.
func (i *Integer) Inspect() string { return fmt.Sprintf("%v", i.Value) }

// Type returns the type of the object.
func (i *Integer) Type() T { return IntegerT }

// Boolean represents an object of that type.
type Boolean struct {
	Value bool
}

// Inspect returns a representation of the object value.
func (b *Boolean) Inspect() string { return fmt.Sprintf("%v", b.Value) }

// Type returns the type of the object.
func (b *Boolean) Type() T { return BooleanT }

// Null represents an object of that type.
type Null struct{}

// Inspect returns a representation of the object value.
func (n *Null) Inspect() string { return "null" }

// Type returns the type of the object.
func (n *Null) Type() T { return NullT }

// ReturnValue represents an object of that type.
type ReturnValue struct {
	Value Object
}

// Inspect returns a representation of the object value.
func (r *ReturnValue) Inspect() string { return r.Value.Inspect() }

// Type returns the type of the object.
func (r *ReturnValue) Type() T { return ReturnValueT }
