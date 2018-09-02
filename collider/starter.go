package collider

import "fmt"

const massThreshold = 17.0

// Collide accelerates each particle to 99.99% of speed of light and collides them
func Collide(materials []Particle) (string, error) {
	if len(materials) == 0 {
		return "Nothing happened", nil
	}
	if err := checkMass(materials); err != nil {
		return "", fmt.Errorf("experiment failed: %v", err)
	}
	return fmt.Sprintf("That was a good experiment. Keep going"), nil
}

func checkMass(materials []Particle) error {
	total := 0.0
	for _, particle := range materials {
		total += particle.Mass
	}
	if total <= 0 {
		return fmt.Errorf("State law prohibits breaking laws of physics")
	} else if total > massThreshold {
		return fmt.Errorf("Chance to create black hole is too high")
	}
	return nil
}
