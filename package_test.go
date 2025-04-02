package gotoaser_test

import (
	"testing"

	toast "github.com/mhpena/gotoaster"
)

func TestPackageImport(t *testing.T) {
	// Create a new toaster - this tests that the package can be imported properly
	toaster := toast.NewToaster()
	if toaster == nil {
		t.Error("Failed to create toaster")
	}

	// Add a toast
	toaster.AddToast(toast.NewToast("Test message", toast.InfoLevel))

	// Verify toast was added
	toasts := toaster.GetToasts()
	if len(toasts) != 1 {
		t.Errorf("Expected 1 toast, got %d", len(toasts))
	}
}
