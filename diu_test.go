package diu

import (
	"testing"

	"github.com/compose-generator/diu/cli"
	"github.com/compose-generator/diu/model"
	"github.com/stretchr/testify/assert"
)

func TestGetImageManifest_Success(t *testing.T) {
	// Test successful
	manifest, err := GetImageManifest("hello-world")
	// Test for error
	assert.Nil(t, err)
	// Test layer count
	assert.Equal(t, 1, len(manifest.SchemaV2Manifest.Layers))
}

func TestGetImageManifest_Filure(t *testing.T) {
	// Test failure
	_, err := GetImageManifest("chillibits/missing-image")
	// Test for error
	assert.NotNil(t, err)
}

func TestGetExistingVolumes_Empty(t *testing.T) {
	// Test with empty volumes
	volumes, err := GetExistingVolumes()
	assert.Nil(t, err)
	// Test volume count
	assert.Equal(t, 0, len(volumes))
}

func TestGetExistingVolumes_One(t *testing.T) {
	// Create volume
	cli.ExecuteAndWaitWithOutput("docker", "volume", "create", "test")
	// Test with one volume
	volumes, err := GetExistingVolumes()
	assert.Nil(t, err)
	// Test volume count
	assert.Equal(t, 1, len(volumes))
	v := volumes[0]
	assert.Equal(t, "test", v.Name)
	assert.Equal(t, "local", v.Driver)
	assert.Nil(t, v.Labels)
	assert.Nil(t, v.Options)
	assert.Equal(t, "local", v.Scope)
	// Cleanup
	cli.ExecuteAndWaitWithOutput("docker", "volume", "rm", "test")
}

func TestGetExistingNetworks(t *testing.T) {
	// Create volume
	cli.ExecuteAndWaitWithOutput("docker", "network", "create", "test")
	// Test with one volume
	networks, err := GetExistingNetworks()
	assert.Nil(t, err)
	// Test Network count
	assert.LessOrEqual(t, 1, len(networks))
	var n model.DockerNetwork
	for _, network := range networks {
		if network.Name == "test" {
			n = network
			break
		}
	}
	assert.Equal(t, "test", n.Name)
	assert.Equal(t, "bridge", n.Driver)
	assert.Empty(t, n.Labels)
	assert.Empty(t, n.Options)
	// Cleanup
	cli.ExecuteAndWaitWithOutput("docker", "network", "rm", "test")
}
