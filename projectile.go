package main

// Projectile contains the representation of the position as (point)
// and the velocity as (vector).
type Projectile struct {
	position *Tuple
	velocity *Tuple
}

// Environment contains the representation of the gravity
// and wind, both as vectors.
type Environment struct {
	gravity *Tuple
	wind    *Tuple
}

// Tick updates the projectile position and velocity,
// the calculations are affected by the environment (gravity + wind) vector settings.
func Tick(env *Environment, p *Projectile) *Projectile {
	position := p.position.Add(p.velocity)
	velocity := p.velocity.Add(env.gravity).Add(env.wind)

	return &Projectile{position, velocity}
}

// FireProjectile simulates and outputs the trayectory ([]point) of a projectile
// based on a initial position (point) and initial velocity (vector)
// it stops when the projectile hits the ground (Y == 0).
func (e *Environment) FireProjectile(projectilePoint, initialVelocity *Tuple) []Tuple {
	projectileTrayectory := []Tuple{}

	// stand alone variables to decouple and avoid overwriting the same memory access problems.
	projectilePointTmp := *projectilePoint
	initialVelocityTmp := *initialVelocity

	currentProjectile := Projectile{
		position: &projectilePointTmp,
		velocity: &initialVelocityTmp,
	}

	for currentProjectile.position.y >= 0 {
		projectileTrayectory = append(projectileTrayectory, *currentProjectile.position)
		currentProjectile = *(Tick(e, &currentProjectile))
	}

	return projectileTrayectory
}

// WriteToCanvas writes the projectile position as a color pixel on the canvas.
func (p *Projectile) WriteToCanvas(canvas *Canvas, color *Color) {
	canvas.WritePixel(int(p.position.x), canvas.height-int(p.position.y), color)
}
