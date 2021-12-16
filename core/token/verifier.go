package token

import (
	"github.com/geeklubcn/doorman/core"
)

type Verifier interface {
	Verify(token string) (core.Identification, bool)
}
