package dates

import (
	"reflect"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNowString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNowString(); got != tt.want {
				t.Errorf("GetNowString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNowDBString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNowDBString(); got != tt.want {
				t.Errorf("GetNowDBString() = %v, want %v", got, tt.want)
			}
		})
	}
}
