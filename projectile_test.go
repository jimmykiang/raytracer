package main

import (
	"fmt"
	"testing"
)

func TestFireProjectile(t *testing.T) {

	gravity := Vector(0, -0.1, 0)
	wind := Vector(-0.01, 0, 0)
	startPosition := Point(0, 1, 0)
	initialVelocity := Vector(1, 1, 0).Normalize()
	speedFactor := 10.0

	environment := Environment{
		gravity: gravity,
		wind:    wind,
	}
	projectileTrayectory := environment.FireProjectile(startPosition, initialVelocity.Multiply(speedFactor))

	fmt.Printf("len(trace) = %+v\n", len(projectileTrayectory))

	for _, p := range projectileTrayectory {
		fmt.Printf("P = %+v\n", *p)
	}
}
