package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		t.Error("Expected not nil")
		return
	}

	currentCount := counter.AddOne()

	if currentCount != 1 {
		t.Errorf("Expected 1 but got %d", currentCount)
		return
	}

	counter2 := GetInstance()
	if counter2 != counter {
		t.Errorf("Initialized a new counter instance, not singleton!")
		return
	}

	newCurrentCount := counter2.AddOne()
	if newCurrentCount != 2 {
		t.Errorf("Expected to increment from the instantiated singleton starting count which is 1, +1=>2")
		return
	}
}
