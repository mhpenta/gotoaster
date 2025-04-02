package gotoaser

import (
	"testing"
)

func TestNewToaster(t *testing.T) {
	toaster := NewToaster()

	if toaster.position != TopRight {
		t.Errorf("Expected default position to be TopRight, got %s", toaster.position)
	}

	if toaster.maxToasts != 5 {
		t.Errorf("Expected default maxToasts to be 5, got %d", toaster.maxToasts)
	}

	if !toaster.autoRemove {
		t.Error("Expected autoRemove to be true by default")
	}
}

func TestAddRemoveToast(t *testing.T) {
	toaster := NewToaster()
	toast := NewToast("Test message", InfoLevel)

	toaster.AddToast(toast)

	if len(toaster.GetToasts()) != 1 {
		t.Errorf("Expected toasts length to be 1, got %d", len(toaster.GetToasts()))
	}

	toaster.RemoveToast(toast.ID)

	if len(toaster.GetToasts()) != 0 {
		t.Errorf("Expected toasts length to be 0 after removal, got %d", len(toaster.GetToasts()))
	}
}
