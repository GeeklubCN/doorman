package token

import (
	"github.com/wangyuheng/doorman/core"
)

type Verifier interface {
	Verify(token string) (core.Identification, bool)
}
