package gotoaser

// Level represents the severity level of a toast notification.
type Level string

// Position represents the placement for the toast container.
type Position string

// Available toast levels
const (
	DefaultLevel Level = "default"
	SuccessLevel Level = "success"
	ErrorLevel   Level = "error"
	WarningLevel Level = "warning"
	InfoLevel    Level = "info"
)

// Available toast positions
const (
	TopRight     Position = "top-right"
	TopLeft      Position = "top-left"
	TopCenter    Position = "top-center"
	BottomRight  Position = "bottom-right"
	BottomLeft   Position = "bottom-left"
	BottomCenter Position = "bottom-center"
)
