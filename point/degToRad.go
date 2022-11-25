package point

import "math"

func degreesToRadians(deg float32) float32 {
	return deg * (math.Pi / 180.0)
}
