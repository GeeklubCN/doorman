package token

import (
	"testing"

	"github.com/geeklubcn/doorman/core"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("should codec", func(t *testing.T) {
		id := core.Identification("fake-id")
		token, err := JwtFactory.Create(id)
		assert.Nil(t, err)
		res, ok := JwtFactory.Verify(token)
		assert.True(t, ok)
		assert.Equal(t, id, res)
	})
	t.Run("should fail when token invalid", func(t *testing.T) {
		res, ok := JwtFactory.Verify("abc")
		assert.False(t, ok)
		assert.Empty(t, res)
	})
}
