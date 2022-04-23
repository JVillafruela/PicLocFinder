package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// get a Bound object from bounding box coordinates
// bbox format : lon1,lat1,lon2,lat2 eg. 5.630665,45.031614,5.634817,45.034214
func bboxBound(str string) (orb.Bound, error) {

	var minLon, minLat, maxLat, maxLon float64

	n, err := fmt.Sscanf(str, "%f,%f,%f,%f", &minLon, &minLat, &maxLon, &maxLat)
	if n != 4 {
		return orb.Bound{}, err
	}
	p1 := orb.Point{minLon, minLat}
	p2 := orb.Point{maxLon, maxLat}

	return orb.MultiPoint{p1, p2}.Bound(), nil
}

// Return true if the point is inside the Bound
func MatchLocation(lat, lon float64, bound orb.Bound) bool {
	point := orb.Point{lon, lat}
	return bound.Contains(point)
}

// for v0.2
func ExampleGeoJson() {

	dat, err := os.ReadFile("E:/temp/sample.geojson")
	check(err)
	fc := geojson.NewFeatureCollection()
	err = json.Unmarshal(dat, &fc)
	check(err)

	for i, v := range fc.Features {
		switch v.Geometry.(type) {
		case orb.Point:
			point := v.Geometry.(orb.Point)
			println(i, "Point", point.Lat(), point.Lon())
		case orb.LineString:
			println(i, "LineString")
		case orb.Polygon:
			println(i, "Polygon")
		}
	}

	/*
		rawJSON := []byte(`
		{ "type": "FeatureCollection",
		  "features": [
			{ "type": "Feature",
			  "geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
			  "properties": {"prop0": "value0"}
			}
		  ]
		}`)

		//fc, _ := geojson.UnmarshalFeatureCollection(rawJSON)

		// or

		fc := geojson.NewFeatureCollection()
		err := json.Unmarshal(rawJSON, &fc)
		fmt.Println(fc, err)

		// Geometry will be unmarshalled into the correct geo.Geometry type.
		point := fc.Features[0].Geometry.(orb.Point)
		fmt.Println("geo", fc.Features[0].Geometry)
		fmt.Println("point", point)

	*/
}
