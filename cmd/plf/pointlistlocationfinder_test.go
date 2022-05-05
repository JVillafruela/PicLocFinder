package main

import (
	"reflect"
	"testing"

	"github.com/paulmach/orb"
)

func Test_pointListFromGeoJSON(t *testing.T) {

	geojson := `
	{
		"name": "Arceaux_EPSG4326",
		"type": "FeatureCollection",
		"features": [{
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [5.72616, 45.18658]
				},
				"properties": {
					"MOB_ARCE_ID": 46,
					"MOB_ARCE_NB": 1
				}
			}, {
				"type": "Feature",
				"geometry": {
					"type": "Point",
					"coordinates": [5.72553, 45.18219]
				},
				"properties": {
					"MOB_ARCE_ID": 970,
					"MOB_ARCE_NB": 2
				}
			}
		]
	} `
	polygon := `
	{
		"type": "FeatureCollection",
		"features": [
		{
			"type": "Feature",
			"properties": {},
			"geometry": {
			"type": "Polygon",
			"coordinates": [
				[
				[5.68769,45.08717],
				[5.68750,45.08679],
				[5.68851,45.08653],
				[5.68877,45.08689],
				[5.68769,45.08717]
				]
			]
			}
		}
		]
	} `

	p1 := orb.Point{5.72616, 45.18658}
	p2 := orb.Point{5.72553, 45.18219}
	points := []orb.Point{p1, p2}

	type args struct {
		dat []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []orb.Point
		wantErr bool
	}{
		{"KO empty", args{[]byte{}}, nil, true},
		{"OK polygon", args{[]byte(polygon)}, []orb.Point{}, false},
		{"OK 2 points", args{[]byte(geojson)}, points, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pointListFromGeoJSON(tt.args.dat)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s : pointListFromGeoJSON() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s : pointListFromGeoJSON() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_pointListBbox(t *testing.T) {
	p1 := orb.Point{5.68769, 45.08717}
	p2 := orb.Point{5.68750, 45.08679}

	type args struct {
		points []orb.Point
	}
	tests := []struct {
		name    string
		args    args
		want    orb.Bound
		wantErr bool
	}{
		{"KO empty", args{[]orb.Point{}}, orb.Bound{}, true},
		{"KO one point", args{[]orb.Point{p1}}, orb.Bound{}, true},
		{"OK two points", args{[]orb.Point{p1, p2}}, orb.MultiPoint{p1, p2}.Bound(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pointListBbox(tt.args.points)
			if (err != nil) != tt.wantErr {
				t.Errorf("pointListBbox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointListBbox() = %v, want %v", got, tt.want)
			}
		})
	}
}
