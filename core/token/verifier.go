package token

import (
	"github.com/geeklubcn/doorman/core"
)

type Verifier interface {
	verify(token string) (core.Identification, bool)
}
