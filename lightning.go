package main

import "math"

// Material encapsulates the given attributes of the Phong reflection model.
type Material struct {
	color                                                                            *Color
	ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64
	pattern                                                                          *Pattern
}

// PointLight is a light source with no size, existing at a single point in space.
type PointLight struct {
	position  *Tuple
	intensity *Color
}

// DefaultMaterial returns a reference to material with default values.
func DefaultMaterial() *Material {
	return NewMaterial(White, 0.1, .9, .9, 200.0, 0.0, 0.0, 1.0, nil)
}

// NewMaterial creates a new Materials
func NewMaterial(color *Color, ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64, pattern *Pattern) *Material {
	return &Material{color, ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex, pattern}
}

// NewPointLight returns a reference to PointLight.
func NewPointLight(position *Tuple, intensity *Color) *PointLight {
	return &PointLight{position, intensity}
}

// Lighting computes the color resulting from diferent parameters at a specific point of the object.
func Lighting(material *Material, object Shape, light *PointLight, point, eyev, normalv *Tuple, inShadow bool) *Color {

	var color *Color
	if material.pattern != nil {
		color = material.pattern.ColorAtObject(object, point)
	} else {
		color = material.color
	}

	effectiveColor := color.Multiply(light.intensity)

	lightv := light.position.Substract(point).Normalize()

	ambient := effectiveColor.MultiplyByScalar(material.ambient)

	lightDotNormal := lightv.DotProduct(normalv)

	diffuse := Black
	specular := Black

	if lightDotNormal >= 0 {
		diffuse = effectiveColor.MultiplyByScalar(material.diffuse).MultiplyByScalar(lightDotNormal)

		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.DotProduct(eyev)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.shininess)
			specular = light.intensity.MultiplyByScalar(material.specular).MultiplyByScalar(factor)
		}
	}

	if inShadow {
		return ambient
	}
	return ambient.Add(diffuse).Add(specular)

}
