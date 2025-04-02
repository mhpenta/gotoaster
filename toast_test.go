package gotoaser

import (
	"testing"
	"time"
)

func TestNewToast(t *testing.T) {
	message := "Test message"
	level := SuccessLevel

	toast := NewToast(message, level)

	if toast.Message != message {
		t.Errorf("Expected message to be '%s', got '%s'", message, toast.Message)
	}

	if toast.Level != level {
		t.Errorf("Expected level to be '%s', got '%s'", level, toast.Level)
	}

	if toast.Duration != 5*time.Second {
		t.Errorf("Expected default duration to be 5s, got %v", toast.Duration)
	}

	if !toast.ShowIcon {
		t.Error("Expected ShowIcon to be true by default")
	}

	if toast.ID == "" {
		t.Error("Expected ID to be generated")
	}
}
