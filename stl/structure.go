package stl

import "github.com/ungerik/go3d/float64/vec3"

// Triangle is the definition of a plan in a STL file
type Triangle struct {
	Vector    [3]float32
	Vertex1   [3]float32
	Vertex2   [3]float32
	Vertex3   [3]float32
	AttrCount uint16
}

// Mesh is the complete structure of the file
type Mesh struct {
	Header        [80]uint8
	TriangleCount uint32
	Triangles     []Triangle
}

// MakeTriangle Create a triangle with 3 points
func MakeTriangle(a, b, c vec3.T) (triangle Triangle) {
	triangle.Vertex1 = [3]float32{
		float32(a[0]),
		float32(a[1]),
		float32(a[2]),
	}
	triangle.Vertex2 = [3]float32{
		float32(b[0]),
		float32(b[1]),
		float32(b[2]),
	}
	triangle.Vertex3 = [3]float32{
		float32(c[0]),
		float32(c[1]),
		float32(c[2]),
	}
	vect0 := &vec3.T{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
	vect1 := &vec3.T{
		a[0] - c[0],
		a[1] - c[1],
		a[2] - c[2],
	}
	ortho := vec3.Cross(vect0, vect1)
	normal := ortho.Normalize()
	triangle.Vector = [3]float32{
		float32(normal[0]),
		float32(normal[1]),
		float32(normal[2]),
	}
	return
}
