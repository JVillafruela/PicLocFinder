package main

import (
	"os"

	"github.com/xor-gate/goexif2/exif"
)

// get the coordinates from the exif data
func PicLocation(fname string) (lat, long float64, err error) {

	f, err := os.Open(fname)
	if err != nil {
		return 0, 0, err
	}

	x, err := exif.Decode(f)
	if err != nil {
		return 0, 0, err
	}

	lat, long, err = x.LatLong()
	if err != nil {
		return 0, 0, err
	}

	return lat, long, err
}
