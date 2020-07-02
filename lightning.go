package main

// Material encapsulates the given attributes of the Phong reflection model.
type Material struct {
	color                                                                            *Color
	ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64
	pattern                                                                          *Pattern
}

// DefaultMaterial returns a reference to material with default values.
func DefaultMaterial() *Material {
	return NewMaterial(White, 0.1, .9, .9, 200.0, 0.0, 0.0, 1.0, nil)
}

// NewMaterial creates a new Materials
func NewMaterial(color *Color, ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64, pattern *Pattern) *Material {
	return &Material{color, ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex, pattern}
}
