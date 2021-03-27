package geoutils

import (
	"math"
	"testing"
)

func TestGetGeoHash(t *testing.T) {

	const expected = "gcw9f44x0"

	lat1 := 53.75734328397976
	lon1 := -2.0180395172300716

	hash, err := GeoHash(lat1, lon1, 9)

	if err != nil {
		t.Errorf("Error creating geohash, %s", err.Error())
	}

	if hash != expected {
		t.Errorf("Unexpected geohash got %s expected %s", hash, expected)
	}

}

func TestLocationFromHash(t *testing.T) {
	geohash := "gcw9f44x0"

	expectedLat := 53.757327
	expectedLon := -2.018030

	lat, lon := LocationFromHash(geohash)

	if floatEquals(lat, expectedLat, 0.0001) && floatEquals(lon, expectedLon, 0.00001) {
		t.Errorf("Geohash does not decode to expected lat/lon have [%f, %f] expected [%f, %f]", lat, lon, expectedLat, expectedLon)
	}

}

func assertStringEquals(expected string, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %v, was %v", expected, actual)
	}
}

func assertFloatEquals(expected float64, actual float64, epsilon float64, t *testing.T) {
	if math.Abs(actual-expected) > epsilon {
		t.Errorf("Expected %f, was %f", expected, actual)
	}
}

func getString(str string, err error) string {
	return str
}

func TestHashValues(t *testing.T) {
	assertStringEquals("7zzzzzzzzz", getString(GeoHash(0, 0, DEFAULT_PRECISION)), t)
	assertStringEquals("2pbpbpbpbp", getString(GeoHash(0, -180, DEFAULT_PRECISION)), t)
	assertStringEquals("rzzzzzzzzz", getString(GeoHash(0, 180, DEFAULT_PRECISION)), t)
	assertStringEquals("5bpbpbpbpb", getString(GeoHash(-90, 0, DEFAULT_PRECISION)), t)
	assertStringEquals("0000000000", getString(GeoHash(-90, -180, DEFAULT_PRECISION)), t)
	assertStringEquals("pbpbpbpbpb", getString(GeoHash(-90, 180, DEFAULT_PRECISION)), t)
	assertStringEquals("gzzzzzzzzz", getString(GeoHash(90, 0, DEFAULT_PRECISION)), t)
	assertStringEquals("bpbpbpbpbp", getString(GeoHash(90, -180, DEFAULT_PRECISION)), t)
	assertStringEquals("zzzzzzzzzz", getString(GeoHash(90, 180, DEFAULT_PRECISION)), t)

	assertStringEquals("9q8yywe56g", getString(GeoHash(37.7853074, -122.4054274, DEFAULT_PRECISION)), t)
	assertStringEquals("dqcjf17sy6", getString(GeoHash(38.98719, -77.250783, DEFAULT_PRECISION)), t)
	assertStringEquals("tj4p5gerfz", getString(GeoHash(29.3760648, 47.9818853, DEFAULT_PRECISION)), t)
	assertStringEquals("umghcygjj7", getString(GeoHash(78.216667, 15.55, DEFAULT_PRECISION)), t)
	assertStringEquals("4qpzmren1k", getString(GeoHash(-54.933333, -67.616667, DEFAULT_PRECISION)), t)
	assertStringEquals("4w2kg3s54y", getString(GeoHash(-54, -67, DEFAULT_PRECISION)), t)
}

func assertHashRoundtrip(lat float64, lng float64, t *testing.T) {
	hashString, _ := GeoHash(lat, lng, DEFAULT_PRECISION)
	retLat, retLng := LocationFromHash(hashString)

	assertFloatEquals(lat, retLat, 0.01, t)
	assertFloatEquals(lng, retLng, 0.01, t)
}

func TestLocationFromHashRoundtrip(t *testing.T) {
	assertHashRoundtrip(37.7853074, -122.4054274, t)
	assertHashRoundtrip(38.98719, -77.250783, t)
	assertHashRoundtrip(29.3760648, 47.9818853, t)
	assertHashRoundtrip(78.216667, 15.55, t)
	assertHashRoundtrip(-54.933333, -67.616667, t)
	assertHashRoundtrip(-54, -67, t)

	assertHashRoundtrip(0, 0, t)
	assertHashRoundtrip(0, -180, t)
	assertHashRoundtrip(0, 180, t)
	assertHashRoundtrip(-90, 0, t)
	assertHashRoundtrip(-90, -180, t)
	assertHashRoundtrip(-90, 180, t)
	assertHashRoundtrip(90, 0, t)
	assertHashRoundtrip(90, -180, t)
	assertHashRoundtrip(90, 180, t)
}

func TestCustomPrecision(t *testing.T) {
	assertStringEquals("000000", getString(GeoHash(-90, -180, 6)), t)
	assertStringEquals("zzzzzzzzzzzzzzzzzzzz", getString(GeoHash(90, 180, 20)), t)
	assertStringEquals("p", getString(GeoHash(-90, 180, 1)), t)
	assertStringEquals("bpbpb", getString(GeoHash(90, -180, 5)), t)
	assertStringEquals("9q8yywe5", getString(GeoHash(37.7853074, -122.4054274, 8)), t)
	assertStringEquals("dqcjf17sy6cppp8vfn", getString(GeoHash(38.98719, -77.250783, 18)), t)
	assertStringEquals("tj4p5gerfzqu", getString(GeoHash(29.3760648, 47.9818853, 12)), t)
	assertStringEquals("u", getString(GeoHash(78.216667, 15.55, 1)), t)
	assertStringEquals("4qpzmre", getString(GeoHash(-54.933333, -67.616667, 7)), t)
	assertStringEquals("4w2kg3s54", getString(GeoHash(-54, -67, 9)), t)
}
