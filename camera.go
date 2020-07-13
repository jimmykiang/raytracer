package main

import "math"

// Camera defines different parameters of it contained in a struct.
type Camera struct {
	hsize, vsize                                  int
	fieldOfView, halfWidth, halfHeight, pixelSize float64
	tranform                                      Matrix
}

// NewCamera returns a pointer to a default camera.
func NewCamera(hsize, vsize int, fieldOfView float64) *Camera {
	c := &Camera{hsize, vsize, fieldOfView, 0, 0, 0, NewIdentityMatrix()}
	c.SetPixelSize()
	return c
}

// SetPixelSize updates the pixel size of the camera based on the parameters defined in the Camera struct.
func (cam *Camera) SetPixelSize() {
	halfView := math.Tan(cam.fieldOfView / 2)
	aspect := float64(cam.hsize) / float64(cam.vsize)
	if aspect >= 1 {
		cam.halfWidth = halfView
		cam.halfHeight = halfView / aspect
	} else {
		cam.halfWidth = halfView * aspect
		cam.halfHeight = halfView
	}
	cam.pixelSize = (cam.halfWidth * 2) / float64(cam.hsize)
}
