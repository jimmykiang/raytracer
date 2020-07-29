package main

// Group will implement all the methods defined in the interface Shape becoming a Shape itself.
type Group struct {
	transform Matrix
	children  []Shape
}

func NewGroup() *Group {

	return &Group{
		transform: IdentityMatrix,
		children:  make([]Shape, 0),
	}
}

func (g *Group) AddChild(shapes ...Shape) {

	for i := 0; i < len(shapes); i++ {
		g.children = append(g.children, shapes[i])
		shapes[i].SetParent(g)
	}
}

func (g *Group) SetMaterial(material *Material)      {}
func (g *Group) SetTransform(Matrix)                 {}
func (g *Group) Transform() Matrix                   { return nil }
func (g *Group) Material() *Material                 { return nil }
func (g *Group) Intersect(*Ray) []*Intersection      { return nil }
func (g *Group) localIntersect(*Ray) []*Intersection { return nil }
func (g *Group) NormalAt(*Tuple) *Tuple              { return nil }
func (g *Group) localNormalAt(*Tuple) *Tuple         { return nil }
func (g *Group) GetParent() Shape                    { return nil }
func (g *Group) SetParent(shape Shape)               {}
