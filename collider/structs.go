package collider

// Particle that we load into collider to collide it with another particle
type Particle struct {
	Element string  `json:"element"`
	Mass    float64 `json:"mass"`
}
