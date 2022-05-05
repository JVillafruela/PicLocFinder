package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/quadtree"
)

type PointListLocationFinder struct {
	qt *quadtree.Quadtree
	// radius in meters
	radius int64
}

func NewPointListLocationFinder(filename string, radius int64) (PointListLocationFinder, error) {
	var pl PointListLocationFinder
	dat, err := os.ReadFile(filename)
	if err != nil {
		return pl, errors.New("error reading geoJSON file")
	}
	points, err := pointListFromGeoJSON(dat)
	if err != nil {
		return pl, err
	}
	qt, err := pointListQuadtree(points)
	if err != nil {
		return pl, err
	}

	pl.qt = qt
	pl.radius = radius
	return pl, nil

}

func (lf PointListLocationFinder) Match(lat, lon float64) bool {
	point := orb.Point{lon, lat}
	nearest := lf.qt.Find(point)
	return geo.Distance(nearest.Point(), point) <= float64(lf.radius)
}

// get points from geoJSON data
func pointListFromGeoJSON(dat []byte) ([]orb.Point, error) {

	points := make([]orb.Point, 0)

	fc := geojson.NewFeatureCollection()
	err := json.Unmarshal(dat, &fc)
	if err != nil {
		return nil, err
	}

	for _, v := range fc.Features {
		switch v.Geometry.(type) {
		case orb.Point:
			p := v.Geometry.(orb.Point)
			points = append(points, p)
			//println(i, "Point", p.Lat(), p.Lon())
		case orb.LineString:
			//println(i, "LineString")
		case orb.Polygon:
			//println(i, "Polygon")
		}
	}
	return points, nil
}

// get bounding box from a list of points
func pointListBbox(points []orb.Point) (orb.Bound, error) {
	if len(points) < 2 {
		return orb.Bound{}, errors.New("Point list empty or counting only one point")
	}

	minLat := points[0].Lat()
	minLon := points[0].Lon()
	maxLat := minLat
	maxLon := minLon

	for _, p := range points {
		lat := p.Lat()
		lon := p.Lon()

		if lat < minLat {
			minLat = lat
		}
		if lat > maxLat {
			maxLat = lat
		}
		if lon < minLon {
			minLon = lon
		}
		if lon > maxLon {
			maxLon = lon
		}
	}
	p1 := orb.Point{minLon, minLat}
	p2 := orb.Point{maxLon, maxLat}

	return orb.MultiPoint{p1, p2}.Bound(), nil
}

// make quadtree from a list of points
func pointListQuadtree(points []orb.Point) (*quadtree.Quadtree, error) {
	bbox, err := pointListBbox(points)
	if err != nil {
		return nil, err
	}

	qt := quadtree.New(bbox)
	for _, p := range points {
		qt.Add(p)
	}

	return qt, nil
}
