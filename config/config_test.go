package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	config, _ := LoadConfiguration("config_test.json")
	assert.EqualValues(t, "4", config.Version)
	assert.EqualValues(t, "7075", config.Node.PeeringPort)
}
