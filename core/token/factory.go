package token

import (
	"github.com/geeklubcn/doorman/core"
)

type Factory interface {
	create(id core.Identification) (string, error)
}
