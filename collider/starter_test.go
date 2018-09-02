package collider

import (
	"strings"
	"testing"
)

func TestCollide(t *testing.T) {
	m := []Particle{
		Particle{"hydrogen", 1.00},
		Particle{"calcium", 0.22},
	}
	result, err := Collide(m)
	if err != nil {
		t.Fatalf("Got error while running experiment: %v", err)
	}

	if !strings.Contains(result, "good experiment") {
		t.Errorf("Expected experiment to be good, but it was %q", result)
	}
}

func TestNothingCollide(t *testing.T) {
	m := []Particle{}
	result, err := Collide(m)
	if err != nil {
		t.Fatalf("Got error while running experiment: %v", err)
	}

	if result != "Nothing happened" {
		t.Errorf("Expected nothing to happen, but something happened: %q", result)
	}
}

func TestNegativeMassCollide(t *testing.T) {
	m := []Particle{
		Particle{"hydrogen", -1.00},
		Particle{"calcium", 0.22},
	}
	_, err := Collide(m)
	if err == nil {
		t.Errorf("Expected to get an error, but there wasn't any")
	}

	if !strings.Contains(err.Error(), "laws of physics") {
		t.Errorf("Expected to get error about laws of physics but got %q", err.Error())
	}
}

func TestHugeMassCollide(t *testing.T) {
	m := []Particle{
		Particle{"hydrogen", 2834.00},
		Particle{"calcium", 0.22},
	}
	_, err := Collide(m)
	if err == nil {
		t.Errorf("Expected to get an error, but there wasn't any")
	}

	if !strings.Contains(err.Error(), "black hole") {
		t.Errorf("Expected to get error about black hole but got %q", err.Error())
	}
}
