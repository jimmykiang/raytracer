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

	env := Environment{
		gravity: gravity,
		wind:    wind,
	}
	trace := env.FireProjectile(startPosition, initialVelocity.Multiply(speedFactor))

	fmt.Printf("len(trace) = %+v\n", len(trace))

	for _, p := range trace {
		fmt.Printf("p = %+v\n", p)
	}
}
