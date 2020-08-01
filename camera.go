package main

import (
	"math"
	"sync"
)

// Camera defines different parameters of it contained in a struct.
type Camera struct {
	hsize, vsize                                  int
	fieldOfView, halfWidth, halfHeight, pixelSize float64
	transform                                     Matrix
}

// NewCamera returns a pointer to a default camera.
func NewCamera(hsize, vsize int, fieldOfView float64) *Camera {
	c := &Camera{hsize: hsize,
		vsize:       vsize,
		fieldOfView: fieldOfView,
		halfWidth:   0,
		halfHeight:  0,
		pixelSize:   0,
		transform:   NewIdentityMatrix(),
	}
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

// RayForPixel computes the world coordinates at the center of the given pixel,
// and then construct a ray that passes through that point.
func (cam *Camera) RayForPixel(x, y int) *Ray {
	px := float64(x)
	py := float64(y)
	xoffset := (px + 0.5) * cam.pixelSize
	yoffset := (py + 0.5) * cam.pixelSize

	worldx := cam.halfWidth - xoffset
	worldy := cam.halfHeight - yoffset

	pixel := cam.transform.Inverse().MultiplyMatrixByTuple(Point(worldx, worldy, -1))

	origin := cam.transform.Inverse().MultiplyMatrixByTuple(Point(0, 0, 0))

	direction := pixel.Substract(origin).Normalize()

	return NewRay(origin, direction)
}

// SetTransform sets the camera’s transformation describing how the world is moved relative to the camera.
func (cam *Camera) SetTransform(transform Matrix) {
	cam.transform = transform
}

// Render calculates the render of a given world on a canvas from the view of the camera.
func (cam *Camera) Render(world *World, recursionDepth int) *Canvas {
	image := NewCanvas(cam.hsize, cam.vsize)

	var wg sync.WaitGroup

	for y := 0; y < cam.vsize; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < cam.hsize; x++ {
				ray := cam.RayForPixel(x, y)
				color := world.ColorAt(ray, recursionDepth)
				image.WritePixel(x, y, color)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
	return image
}

type renderResult struct {
	x, y  int
	color *Color
}

// renderWorker function meant to be instantiated as a independent thread.
func (cam *Camera) renderWorker(world *World, recursionDepth int, jobs <-chan int, results chan<- *renderResult, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	for y := range jobs {
		for x := 0; x < cam.hsize; x++ {
			ray := cam.RayForPixel(x, y)

			// fmt.Println("worker:", workerId, "got job y:", y, "x:", x, "color", world.ColorAt(ray, recursionDepth))

			results <- &renderResult{
				color: world.ColorAt(ray, recursionDepth),
				x:     x,
				y:     y,
			}

		}
	}
}

// RenderWithThreadPool calculates the render of a given world on a canvas from the view of the camera.
// this will use a limited amount of threads as workers for rendering.
func (cam *Camera) RenderWithThreadPool(world *World, recursionDepth int) *Canvas {
	image := NewCanvas(cam.hsize, cam.vsize)

	var wg sync.WaitGroup

	// Amount of threads
	threadSize := 7
	numJobs := cam.vsize
	resultSize := cam.vsize * cam.hsize
	jobs := make(chan int, numJobs)
	results := make(chan *renderResult, resultSize)

	for worker := 1; worker <= threadSize; worker++ {
		wg.Add(1)
		go cam.renderWorker(world, recursionDepth, jobs, results, &wg, worker)
	}

	for y := 0; y < numJobs; y++ {

		jobs <- y
	}
	close(jobs)

	for a := 0; a < resultSize; a++ {

		resultStruct := <-results
		image.WritePixel((*resultStruct).x, (*resultStruct).y, (*resultStruct).color)
	}

	wg.Wait()

	return image
}
