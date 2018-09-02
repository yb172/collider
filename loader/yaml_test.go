package loader

import "testing"

func TestLoadFromYaml(t *testing.T) {
	m, err := LoadFromYaml("../warehouse/starter-kit.yaml")
	if err != nil {
		t.Fatalf("Error while loading particles: %v", err)
	}

	if len(m) != 2 {
		t.Fatalf("Expected to have 2 particles but got: %v", m)
	}

	if m[0].Element != "hydrogen" {
		t.Errorf("Expected to have hydrogen as first particle but got %v", m[0].Element)
	}
	if m[0].Mass != 1.64 {
		t.Errorf("Expected to have 1.64 of hydrogen as first particle but got %v", m[0].Mass)
	}

	if m[1].Element != "oxygen" {
		t.Errorf("Expected to have oxygen as second particle but got %v", m[1].Element)
	}
	if m[1].Mass != 8.4 {
		t.Errorf("Expected to have 8.4 of oxygen as second particle but got %v", m[1].Mass)
	}
}

func TestLoadFromNonExistingYaml(t *testing.T) {
	_, err := LoadFromYaml("../warehouse/advanced-kit.yaml")
	if err == nil {
		t.Errorf("Expected to get an error, but there wasn't any")
	}
}

func TestLoadFromBrokenYaml(t *testing.T) {
	_, err := LoadFromYaml("../warehouse/broken.yaml")
	if err == nil {
		t.Errorf("Expected to get an error, but there wasn't any")
	}
}
