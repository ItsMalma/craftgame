package phys

const epsilon = 0.0

type AABB struct {
	X0, Y0, Z0, X1, Y1, Z1 float32
}

func NewAABB(x0, y0, z0, x1, y1, z1 float32) AABB {
	return AABB{
		X0: x0,
		Y0: y0,
		Z0: z0,
		X1: x1,
		Y1: y1,
		Z1: z1,
	}
}

func (aabb AABB) Expand(xa, ya, za float32) AABB {
	if xa < 0 {
		aabb.X0 += xa
	}
	if xa > 0 {
		aabb.X1 += xa
	}

	if ya < 0 {
		aabb.Y0 += ya
	}
	if ya > 0 {
		aabb.Y1 += ya
	}

	if za < 0 {
		aabb.Z0 += za
	}
	if za > 0 {
		aabb.Z1 += za
	}

	return aabb
}

func (aabb AABB) Grow(xa, ya, za float32) AABB {
	return NewAABB(
		aabb.X0-xa,
		aabb.Y0-ya,
		aabb.Z0-za,
		aabb.X1+xa,
		aabb.Y1+ya,
		aabb.Z1+za,
	)
}

func (aabb AABB) ClipXCollide(c AABB, xa float32) float32 {
	if !(c.Y1 <= aabb.Y0) && !(c.Y0 >= aabb.Y1) {
		if !(c.Z1 <= aabb.Z0) && !(c.Z0 >= aabb.Z1) {
			max := float32(0.0)
			if xa > 0.0 && c.X1 <= aabb.X0 {
				max = aabb.X0 - c.X1 - epsilon
				if max < xa {
					xa = max
				}
			}

			if xa < 0.0 && c.X0 >= aabb.X1 {
				max = aabb.X1 - c.X0 + epsilon
				if max > xa {
					xa = max
				}
			}

			return xa
		} else {
			return xa
		}
	} else {
		return xa
	}

}

func (aabb AABB) ClipYCollide(c AABB, ya float32) float32 {
	if c.X1 > aabb.X0 && c.X0 < aabb.X1 {
		if c.Z1 > aabb.Z0 && c.Z0 < aabb.Z1 {
			var max float32
			if ya > 0.0 && c.Y1 <= aabb.Y0 {
				max = aabb.Y0 - c.Y1 - epsilon
				if max < ya {
					ya = max
				}
			}

			if ya < 0.0 && c.Y0 >= aabb.Y1 {
				max = aabb.Y1 - c.Y0 + epsilon
				if max > ya {
					ya = max
				}
			}

			return ya
		} else {
			return ya
		}
	} else {
		return ya
	}
}

func (aabb AABB) ClipZCollide(c AABB, za float32) float32 {
	if c.X1 > aabb.X0 && c.X0 < aabb.X1 {
		if c.Y1 > aabb.Y0 && c.Y0 < aabb.Y1 {
			var max float32
			if za > 0.0 && c.Z1 <= aabb.Z0 {
				max = aabb.Z0 - c.Z1 - epsilon
				if max < za {
					za = max
				}
			}

			if za < 0.0 && c.Z0 >= aabb.Z1 {
				max = aabb.Z1 - c.Z0 + epsilon
				if max > za {
					za = max
				}
			}

			return za
		} else {
			return za
		}
	} else {
		return za
	}
}

func (aabb AABB) Intersects(c AABB) bool {
	if c.X1 > aabb.X0 && c.X0 < aabb.X1 {
		if c.Y1 > aabb.Y0 && c.Y0 < aabb.Y1 {
			return c.Z1 > aabb.Z0 && c.Z0 < aabb.Z1
		} else {
			return false
		}
	} else {
		return false
	}
}

func (aabb AABB) Move(xa, ya, za float32) AABB {
	return NewAABB(
		aabb.X0+xa,
		aabb.Y0+ya,
		aabb.Z0+za,
		aabb.X1+xa,
		aabb.Y1+ya,
		aabb.Z1+za,
	)
}
