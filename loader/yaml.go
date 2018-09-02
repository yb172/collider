package loader

import (
	"fmt"
	"io/ioutil"

	"github.com/yb172/collider/collider"
	"gopkg.in/yaml.v2"
)

// LoadFromYaml loads collider with particles specified in provided yaml file
func LoadFromYaml(path string) ([]collider.Particle, error) {
	yamlBox, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("forklift haven't found yaml box: %v", err)
	}
	var particles []collider.Particle
	if err := yaml.Unmarshal(yamlBox, &particles); err != nil {
		return nil, fmt.Errorf("there was a problem while unpacking: %v", err)
	}
	return particles, nil
}
