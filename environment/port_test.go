package environment_test

import (
	"os"
	"testing"

	"github.com/masl/answertag/environment"
	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	t.Run("Returns port from PORT environment variable if it is set", func(t *testing.T) {
		os.Setenv("PORT", "4000")

		port, err := environment.Port(3000)

		assert.NoError(t, err)
		assert.Equal(t, 4000, port)
	})

	t.Run("Returns default port if PORT environment variable is not set", func(t *testing.T) {
		os.Unsetenv("PORT")

		port, err := environment.Port(3000)

		assert.NoError(t, err)
		assert.Equal(t, 3000, port)
	})

	t.Run("Returns error if PORT environment variable is set to an invalid value", func(t *testing.T) {
		os.Setenv("PORT", "moin")

		_, err := environment.Port(8080)

		assert.Error(t, err)
	})
}
