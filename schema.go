package schema

import (
)

type Any map[string]interface{}
type AnyArray []Any

func NewSchema() Any {
	return Any {
		"@context": "http://schema.org",
	}
}

func NewMovie() Any {
	return NewEmptyType("Movie")
}

func NewPerson() Any {
	return NewEmptyType("Person")
}

func NewOrganization() Any {
	return NewEmptyType("Organization")
}

func NewEmptyType(t string) Any {
	return Any {
		"@type": t,
	}
}
