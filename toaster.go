package gotoaser

import (
	"sync"
)

// Toaster manages a collection of toast notifications
type Toaster struct {
	mu         sync.Mutex
	toasts     []*Toast
	position   Position
	maxToasts  int
	autoRemove bool
}

// NewToaster creates a new toast manager
func NewToaster() *Toaster {
	return &Toaster{
		position:   TopRight,
		maxToasts:  5,
		autoRemove: true,
	}
}

// SetPosition sets the position of the toast container
func (t *Toaster) SetPosition(position Position) *Toaster {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.position = position
	return t
}

// SetMaxToasts sets the maximum number of toasts to display at once
func (t *Toaster) SetMaxToasts(max int) *Toaster {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.maxToasts = max
	return t
}

// SetAutoRemove sets whether toasts should be automatically removed
func (t *Toaster) SetAutoRemove(auto bool) *Toaster {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.autoRemove = auto
	return t
}

// AddToast adds a toast to the toaster
func (t *Toaster) AddToast(toast *Toast) *Toaster {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Add toast to the beginning so newest are first
	t.toasts = append([]*Toast{toast}, t.toasts...)

	// Trim to max toasts
	if len(t.toasts) > t.maxToasts {
		t.toasts = t.toasts[:t.maxToasts]
	}

	return t
}

// RemoveToast removes a toast by ID
func (t *Toaster) RemoveToast(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i, toast := range t.toasts {
		if toast.ID == id {
			t.toasts = append(t.toasts[:i], t.toasts[i+1:]...)
			return
		}
	}
}

// ClearToasts removes all toasts
func (t *Toaster) ClearToasts() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.toasts = []*Toast{}
}

// GetToasts returns a copy of the current toasts
func (t *Toaster) GetToasts() []*Toast {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Return a copy to prevent race conditions
	toastsCopy := make([]*Toast, len(t.toasts))
	copy(toastsCopy, t.toasts)
	return toastsCopy
}

// GetPosition returns the current position
func (t *Toaster) GetPosition() Position {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.position
}
