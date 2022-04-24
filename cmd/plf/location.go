package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/geojson"
)

type LocationFinder interface {
	Match(lat, lon float64) bool
}

type BboxLocationFinder struct {
	bbox orb.Bound
}

type CircleLocationFinder struct {
	center orb.Point
	// radius in meters
	radius int64
}

// TODO PointListLocationFinder

// BboxLocationFinder constructor
func NewBboxLocationFinder(bbox string) (BboxLocationFinder, error) {
	bound, err := BboxBound(bbox)
	if err != nil {
		return BboxLocationFinder{}, err
	}
	lf := BboxLocationFinder{bbox: bound}
	return lf, nil
}

// Check if the point is inside the bounding box
func (lf BboxLocationFinder) Match(lat, lon float64) bool {
	point := orb.Point{lon, lat}
	return lf.bbox.Contains(point)
}

// CircleLocationFinder constructor
func NewCircleLocationFinder(lat, lon float64, radius int64) (CircleLocationFinder, error) {
	err := validateLatitude(lat)
	if err != nil {
		return CircleLocationFinder{}, err
	}
	err = validateLongitude(lon)
	if err != nil {
		return CircleLocationFinder{}, err
	}
	if radius <= 0 {
		return CircleLocationFinder{}, errors.New("invalid radius")
	}
	point := orb.Point{lon, lat}
	return CircleLocationFinder{point, radius}, nil
}

func (lf CircleLocationFinder) Match(lat, lon float64) bool {
	point := orb.Point{lon, lat}
	return geo.Distance(lf.center, point) <= float64(lf.radius)
}

// get a Bound object from bounding box coordinates
// bbox format : lon1,lat1,lon2,lat2 eg. 5.63066,45.03161,5.63481,45.03421
func BboxBound(str string) (orb.Bound, error) {

	var minLon, minLat, maxLat, maxLon float64

	n, err := fmt.Sscanf(str, "%f,%f,%f,%f", &minLon, &minLat, &maxLon, &maxLat)
	if n != 4 {
		return orb.Bound{}, err
	}
	err = validateLatitude(minLat)
	if err != nil {
		return orb.Bound{}, err
	}
	err = validateLatitude(maxLat)
	if err != nil {
		return orb.Bound{}, err
	}
	err = validateLongitude(minLon)
	if err != nil {
		return orb.Bound{}, err
	}
	err = validateLongitude(maxLon)
	if err != nil {
		return orb.Bound{}, err
	}

	p1 := orb.Point{minLon, minLat}
	p2 := orb.Point{maxLon, maxLat}

	return orb.MultiPoint{p1, p2}.Bound(), nil
}

// validate a latitude in WGS84 system
func validateLatitude(lat float64) error {
	if lat < -90 || lat > 90 {
		return errors.New("invalid latitude (WGS84 [-90,+90])")
	}
	return nil
}

// validate a longitude in WGS84 system
func validateLongitude(lon float64) error {
	if lon < -180 || lon > 180 {
		return errors.New("invalid longitude (WGS84 [-180,+180])")
	}
	return nil
}

// for v0.3
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

// to be deleted
func check(e error) {
	if e != nil {
		panic(e)
	}
}
