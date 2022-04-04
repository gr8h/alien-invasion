package internal

import (
	"testing"
)

func TestDestroy_ShouldBeDead_WhenDestroyIsCalled(t *testing.T) {

	var alien Alien = NewAlien(1)

	alien.Destroy()

	if alien.IsAlive() {
		t.Errorf("Alien %d, should be dead!", alien.Id)
	}
}
