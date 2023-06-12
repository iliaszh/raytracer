package geometry

import (
	"math"
)

type Vec3[F Float] struct {
	X, Y, Z F
}

func (v *Vec3[F]) Negate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v Vec3[F]) Add(other Vec3[F]) Vec3[F] {
	return Vec3[F]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3[F]) Sub(other Vec3[F]) Vec3[F] {
	other.Negate()
	return v.Add(other)
}

func (v Vec3[F]) Dot(other Vec3[F]) F {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3[F]) Sq() F {
	return v.Dot(v)
}

func (v Vec3[F]) Len() F {
	return F(math.Sqrt(float64(v.Sq())))
}

func (v *Vec3[F]) Normalize() {
	length := v.Len()
	if length == 0.0 {
		panic("normalize: vector length is zero")
	}

	v.X /= length
	v.Y /= length
	v.Z /= length
}

func (v Vec3[F]) Mul(scalar F) Vec3[F] {
	return Vec3[F]{
		X: scalar * v.X,
		Y: scalar * v.Y,
		Z: scalar * v.Z,
	}
}
