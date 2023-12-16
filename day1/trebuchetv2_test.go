package main

import "testing"

func TestFirstRowV2(t *testing.T)  {
	
	got, _ := CalculateCalibrationValueV2("two1nine")
	want := 29

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSecondRowV2(t *testing.T)  {
	
	got, _ := CalculateCalibrationValueV2("eightwothree")
	want := 83

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestThirdRowV2(t *testing.T)  {
	
	got, _ := CalculateCalibrationValueV2("abcone2threexyz")
	want := 13

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFourthRowV2(t *testing.T)  {
	
	got, _ := CalculateCalibrationValueV2("xtwoone3four")
	want := 24

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

