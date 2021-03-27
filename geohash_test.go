package geoutils

import "testing"

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
