package auth

import (
	"github.com/geeklubcn/doorman/core"
)

type Identifier interface {
	Identify(code string) (core.Identification, bool)
}

type FakeIdentifier struct{}

func (f FakeIdentifier) identify(code string) core.Identification {
	return "fake"
}
