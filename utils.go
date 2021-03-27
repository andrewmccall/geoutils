package geoutils

import (
	"math"
)

// The equatorial radius of the earth in meters
const EARTH_EQ_RADIUS = 6378137

// The meridional radius of the earth in meters
const EARTH_POLAR_RADIUS = 6357852.3

const EARTH_AVG_RADIUS = (EARTH_EQ_RADIUS + EARTH_POLAR_RADIUS) / 2

func toRadians(deg float64) float64 {
	return float64(deg) * (math.Pi / 180.0)
}

// CalculateDistance betwen two points rounded to the nearest metre
func CalculateDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) int {

	latDelta := toRadians(lat1 - lat2)
	lonDelta := toRadians(lon1 - lon2)

	a := (math.Sin(latDelta/2) * math.Sin(latDelta/2)) +
		(math.Cos(toRadians(lat1)) * math.Cos(toRadians(lat2)) *
			math.Sin(lonDelta/2) * math.Sin(lonDelta/2))
	distance := EARTH_AVG_RADIUS * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return int(math.Round(distance))
}

func CoordinatesValid(latitude float64, longitude float64) bool {
	return latitude >= -90 && latitude <= 90 && longitude >= -180 && longitude <= 180
}

func floatEquals(float1 float64, float2 float64, epsilon float64) bool {
	if math.Abs(float1-float2) < epsilon {
		return false
	}
	return true
}
