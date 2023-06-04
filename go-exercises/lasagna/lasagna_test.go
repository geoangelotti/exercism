package lasagna

import "testing"

type lasagnaTests struct {
	name     string
	layers   int
	time     int
	expected int
}

func TestOvenTime(t *testing.T) {
	tt := lasagnaTests{
		name:     "Calculates how many minutes the lasagna should be in the oven",
		layers:   0,
		time:     40,
		expected: 40,
	}
	if got := OvenTime; got != tt.expected {
		t.Errorf("OvenTime(%d)  = %d; want %d", tt.expected, got, tt.expected)
	}
}

func TestRemainingOvenTime(t *testing.T) {
	tt := lasagnaTests{
		name:     "Remaining minutes in oven",
		layers:   0,
		time:     15,
		expected: 25,
	}
	if got := RemainingOvenTime(tt.time); got != tt.expected {
		t.Errorf("RemainingOvenTime(%d) = %d; want %d", tt.time, got, tt.expected)
	}
}

func TestPreparationTimeSingle(t *testing.T) {
	tt := lasagnaTests{
		name:     "Preparation time in minutes for one layer",
		layers:   1,
		time:     0,
		expected: 2,
	}
	if got := PreparationTime(tt.layers); got != tt.expected {
		t.Errorf("PreparationTime(%d) = %d; want %d", tt.layers, got, tt.expected)
	}
}

func TestPreparationTimeMultiple(t *testing.T) {
	tt := lasagnaTests{
		name:     "Preparation time in minutes for multiple layers",
		layers:   4,
		time:     0,
		expected: 8,
	}
	if got := PreparationTime(tt.layers); got != tt.expected {
		t.Errorf("PreparationTime(%d) = %d; want %d", tt.layers, got, tt.expected)
	}
}

func TestElapsedTime(t *testing.T) {
	tt := lasagnaTests{
		name:     "Total time in minutes for one layer",
		layers:   1,
		time:     30,
		expected: 32,
	}
	if got := ElapsedTime(tt.layers, tt.time); got != tt.expected {
		t.Errorf("ElapsedTime(%d, %d) = %d; want %d", tt.layers, tt.time, got, tt.expected)
	}
}
