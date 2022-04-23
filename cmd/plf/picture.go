package main

import (
	"fmt"
	"log"
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

func ExampleDecode() {
	fname := "E:/OSM/gps/2022/2022-04-16 Varces/IMG_20220416_150409.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel.StringVal())

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v \n", numer, denom)

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}
