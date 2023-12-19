package main

import "testing"


func TestFirstRow_highestBlue(t *testing.T) {
	got := CreateGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	gotBlue := got.HighestColorSets["blue"]
	gotRed := got.HighestColorSets["red"]
	wantBlue := 6
	wantRed := 4

	if gotBlue != wantBlue {
		t.Errorf("got %d want %d", gotBlue, wantBlue)
	}

	if gotRed != wantRed {
		t.Errorf("got %d want %d", gotRed, wantRed)
	}
}

func TestFirstRow_isConfigurationValid_isValid(t *testing.T) {
	game := CreateGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	maxAvailableColors := map[string]int{
		"blue": 14,
		"red": 12,
		"green": 13,
	}  
	got := game.IsConfigurationValid(maxAvailableColors)
	want := true

	if got != want {
		t.Errorf("Got %t want %t", got, want)
	}
}

func TestFirstRow_isConfigurationValid_isNotValid(t *testing.T) {
	game := CreateGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	maxAvailableColors := map[string]int{
		"blue": 1,
		"red": 12,
		"green": 13,
	}  
	got := game.IsConfigurationValid(maxAvailableColors)
	want := false

	if got != want {
		t.Errorf("Got %t want %t", got, want)
	}
}

func TestFirstRow_PowerOfSet(t *testing.T) {
	game := CreateGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	got := game.GetPowerOfSet()
	want := 48

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

func TestSecondRow_PowerOfSet(t *testing.T) {
	game := CreateGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	got := game.GetPowerOfSet()
	want := 12

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

func TestThirdRow_PowerOfSet(t *testing.T) {
	game := CreateGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	got := game.GetPowerOfSet()
	want := 1560

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

func TestFourthRow_PowerOfSet(t *testing.T) {
	game := CreateGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	got := game.GetPowerOfSet()
	want := 630

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}


