package geometry

type Ray[F Float] struct {
	Origin    Vec3[F]
	Direction Vec3[F]
}

func NewRay[F Float](origin, direction Vec3[F]) Ray[F] {
	direction.Normalize()

	return Ray[F]{
		Origin:    origin,
		Direction: direction,
	}
}
