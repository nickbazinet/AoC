package main

import "testing"

func TestFirstRow(t *testing.T)  {
	
	got, _ := CalculateCalibrationValue("1abc2")
	want := 12

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSecondRow(t *testing.T)  {
	
	got, _ := CalculateCalibrationValue("pqr3stu8vwx")
	want := 38

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestThirdRow(t *testing.T)  {
	
	got, _ := CalculateCalibrationValue("a1b2c3d4e5f")
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFourthRow(t *testing.T)  {
	
	got, _ := CalculateCalibrationValue("treb7uchet")
	want := 77

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

