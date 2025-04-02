package gotoaser

import "time"

const defaultDuration = 5 * time.Second

// Toast represents a single toast notification
type Toast struct {
	ID        string
	Message   string
	Level     Level
	CreatedAt time.Time
	Duration  time.Duration
	ShowIcon  bool
}

// NewToast creates a new toast notification
func NewToast(message string, level Level) *Toast {
	return &Toast{
		ID:        generateID(),
		Message:   message,
		Level:     level,
		CreatedAt: time.Now(),
		Duration:  defaultDuration,
		ShowIcon:  true,
	}
}

// NewToastWithDuration creates a new toast notification with a specified duration
func NewToastWithDuration(message string, level Level, duration time.Duration) *Toast {
	return &Toast{
		ID:        generateID(),
		Message:   message,
		Level:     level,
		CreatedAt: time.Now(),
		Duration:  duration,
		ShowIcon:  true,
	}
}

// generateID creates a simple timestamp-based ID
func generateID() string {
	return time.Now().Format("20060102150405.000")
}
