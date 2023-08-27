package token

import (
	"github.com/wangyuheng/doorman/core"
)

type Factory interface {
	Create(id core.Identification) (string, error)
}
