package utils

type Vector3 struct {
	X, Y, Z float32
}

func NewVector(x float32, y float32, z float32) Vector3 {
	return Vector3{x, y, z}
}
