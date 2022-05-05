package main

import (
	"github.com/paulmach/orb"
)

type BboxLocationFinder struct {
	bbox orb.Bound
}

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
