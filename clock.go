package main

// Clock will return a *Canvas with the pixels drawn in a circle ()representing the 12 hours of a clock.
func Clock() *Canvas {
	colors := [3]*Color{NewColor(1, 0, 0), NewColor(0, 1, 0), NewColor(0, 0, 1)}

	width, height := 400, 400
	canvas := NewCanvas(width, height)

	canvas.SetOrigin(canvas.width/2, canvas.height/2)

	origin := Point(0, 0, 0)

	canvas.WriteTuple(origin, NewColor(1, 1, 1))

	current := origin.Transform(Translation(0, float64(width*3/8), 0))

	for i := 0; i < 12; i++ {
		current = current.Transform(RotationZ(PI / 6))
		canvas.WriteTuple(current, colors[i%3])
	}

	return canvas
}
