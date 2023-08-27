package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangyuheng/doorman/core"
)

func TestJwt(t *testing.T) {
	t.Run("should codec", func(t *testing.T) {
		id := core.Identification("fake-id")
		token, err := Jwt.Create(id)
		assert.Nil(t, err)
		res, ok := Jwt.Verify(token)
		assert.True(t, ok)
		assert.Equal(t, id, res)
	})
	t.Run("should fail when token invalid", func(t *testing.T) {
		res, ok := Jwt.Verify("abc")
		assert.False(t, ok)
		assert.Empty(t, res)
	})
}
