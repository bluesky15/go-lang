package operation

import (
	"testing"
)

func TestAdd(t *testing.T) {
	//Arrange
	//Act
	got := Add(5, 6)
	//Assert
	if got != 11 {
		t.Errorf("Add(5,6 )= %d; want11", got)
	}
}
