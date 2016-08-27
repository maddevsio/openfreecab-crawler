package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeCanGetConfiguration(t *testing.T) {
	cr := NewConfigurator()
	os.Clearenv()
	os.Setenv("UPDATE_INTERVAL", "160")
	cr.Run()
	conf := cr.Get()
	assert.Equal(t, conf.UpdateInterval, 160)
}
