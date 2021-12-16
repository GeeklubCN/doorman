package token

import (
	"github.com/geeklubcn/doorman/core"
)

type Factory interface {
	Create(id core.Identification) (string, error)
}
