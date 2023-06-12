package geometry

import (
	"math"
)

type Sphere[F Float] struct {
	Origin Vec3[F]
	Radius F
}

func (s Sphere[F]) Intersect(ray Ray[F]) (dist F, intersects bool) {
	sphereR := s.Origin.Sub(ray.Origin)

	proj := sphereR.Dot(ray.Direction)

	distSq := sphereR.Sq() - proj*proj

	radiusSq := s.Radius * s.Radius
	if distSq >= radiusSq {
		return 0.0, false
	}

	distToIntersection := F(math.Sqrt(float64(radiusSq - distSq)))

	dist = proj - distToIntersection
	if dist >= 0.0 {
		return dist, true
	}

	dist = proj + distToIntersection
	if dist < 0.0 {
		return 0.0, false
	}

	return dist, true
}
