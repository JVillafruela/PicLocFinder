package main

import (
	"testing"
)

func Test_hasPicExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"w/o ext", args{"file"}, false},
		{"jpeg ext", args{"cover.jpeg"}, true},
		{"JPEG ext", args{"COVER.JPEG"}, true},
		{"JPG ext", args{"file.JPG"}, true},
		{"mp3 ext", args{"FILE.MP3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPicExtension(tt.args.filename); got != tt.want {
				t.Errorf("hasPicExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasGeoJSONExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"w/o ext", args{"file"}, false},
		{"json ext", args{"points.json"}, true},
		{"geojson ext", args{"points.geojson"}, true},
		{"JSON ext", args{"POINTS.JSON"}, true},
		{"GEOJSON ext", args{"POINTS.GEOJSON"}, true},
		{"txt ext", args{"points.txt"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGeoJSONExtension(tt.args.filename); got != tt.want {
				t.Errorf("hasGeoJSONExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
