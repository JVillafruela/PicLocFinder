package main

import (
	"reflect"
	"testing"

	"github.com/paulmach/orb"
)

func Test_bboxBound(t *testing.T) {
	//bbox='5.630665,45.031614,5.634817,45.034214'
	p1 := orb.Point{5.630665, 45.031614}
	p2 := orb.Point{5.634817, 45.034214}
	okBound := orb.MultiPoint{p1, p2}.Bound()
	errBound := orb.Bound{}

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    orb.Bound
		wantErr bool
	}{
		{"OK min,max", args{"5.630665,45.031614,5.634817,45.034214"}, okBound, false},
		{"OK max,min", args{"5.634817,45.034214,5.630665,45.031614"}, okBound, false},
		{"OK with leading spaces", args{"5.630665, 45.031614, 5.634817, 45.034214"}, okBound, false},
		{"KO with trailing spaces", args{"5.630665 ,45.031614,5.634817,45.034214"}, errBound, true},
		{"KO non numeric", args{"5.630665,45.031614,5.634817,fortytwo"}, errBound, true},
		{"KO missing coordinate", args{"5.630665,45.031614,5.634817"}, errBound, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bboxBound(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("bboxBound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bboxBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchLocation(t *testing.T) {
	bound, _ := bboxBound("5.686787,45.085960,5.689791,45.087786")

	type args struct {
		lat   float64
		lon   float64
		bound orb.Bound
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"OK within", args{45.0869, 5.6880, bound}, true},
		{"KO outside", args{45.0412, 5.6959, bound}, false},
		{"KO empty bound", args{45.0412, 5.6959, orb.Bound{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchLocation(tt.args.lat, tt.args.lon, tt.args.bound); got != tt.want {
				t.Errorf("MatchLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
