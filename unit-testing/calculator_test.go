package unit_testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Subtract(2, 3)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Multiply(4, 2)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestAddTestfiy(t *testing.T) {
	result := Calculator{}.Add(2, 3)
	expected := 5.0
	assert.Equal(t, expected, result)

}
