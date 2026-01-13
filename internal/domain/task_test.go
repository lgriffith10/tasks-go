package domain

import "testing"

func TestChangeDoneStatus_ShouldSucceed(t *testing.T) {
	// Arrange
	task := NewTask("Test")

	// Act
	task.ChangeDoneStatus()

	// Assert
	if task.Done == false {
		t.Errorf("Expected task 'Test' to exist")
	}
}

func TestChangeDoneStatus_WithAlreadyDoneTask_ShouldSucceed(t *testing.T) {
	// Arrange
	task := NewTask("Test")
	task.Done = true

	// Act
	task.ChangeDoneStatus()

	// Assert
	if task.Done == true {
		t.Errorf("Expected task 'Test' to not be done")
	}
}
