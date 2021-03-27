package geoutils

import (
	"errors"
	"fmt"
	"strings"
)

const MAX_PRECISION = 22
const DEFAULT_PRECISION = 10

func GeoHash(latitude float64, longitude float64, precision uint) (string, error) {
	if precision < 1 {
		return "", errors.New("Precision of GeoHash must be larger than zero!")
	}
	if precision > MAX_PRECISION {
		return "", errors.New(fmt.Sprintf("Precision of a GeoHash must be less than %d!", (MAX_PRECISION + 1)))
	}
	if !CoordinatesValid(latitude, longitude) {
		return "", errors.New(fmt.Sprintf("Not valid location coordinates: [%f, %f]", latitude, longitude))
	}
	longitudeRange := []float64{-180, 180}
	latitudeRange := []float64{-90, 90}

	var buffer strings.Builder

	for i := 0; uint(i) < precision; i++ {
		var hashValue = 0
		var val float64
		var vrange []float64

		for j := 0; j < BITS_PER_CHAR; j++ {
			even := (((i * BITS_PER_CHAR) + j) % 2) == 0
			if even {
				val = longitude
				vrange = longitudeRange
			} else {
				val = latitude
				vrange = latitudeRange
			}
			mid := (vrange[0] + vrange[1]) / 2
			if val > mid {
				hashValue = (hashValue << 1) + 1
				vrange[0] = mid
			} else {
				hashValue = hashValue << 1
				vrange[1] = mid
			}
		}
		buffer.WriteByte(ToBase32Char(hashValue))
	}
	return buffer.String(), nil
}

func LocationFromHash(hash string) (lattitude float64, longitude float64) {
	var decoded = 0
	var numBits = len(hash) * BITS_PER_CHAR

	for i := 0; i < len(hash); i++ {
		charVal := ToBase32Value(hash[i])
		decoded = decoded << BITS_PER_CHAR
		decoded = decoded + charVal
	}

	var minLng = float64(-180)
	var maxLng = float64(180)

	var minLat = float64(-90)
	var maxLat = float64(90)

	for i := 0; i < numBits; i++ {
		// Get the high bit
		bit := (decoded >> (numBits - i - 1)) & 1

		// Even bits are longitude, odd bits are latitude
		if i%2 == 0 {
			if bit == 1 {
				minLng = (minLng + maxLng) / 2
			} else {
				maxLng = (minLng + maxLng) / 2
			}
		} else {
			if bit == 1 {
				minLat = (minLat + maxLat) / 2
			} else {
				maxLat = (minLat + maxLat) / 2
			}
		}
	}

	lat := (minLat + maxLat) / 2
	lng := (minLng + maxLng) / 2

	return lat, lng
}
