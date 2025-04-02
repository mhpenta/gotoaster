package gotoaser

import (
	"fmt"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// CSS classes for different toast levels
var levelClasses = map[Level]string{
	DefaultLevel: "toast-default",
	SuccessLevel: "toast-success",
	ErrorLevel:   "toast-error",
	WarningLevel: "toast-warning",
	InfoLevel:    "toast-info",
}

// GetPositionClass returns the CSS class for a position
func GetPositionClass(position Position) string {
	return fmt.Sprintf("toaster-%s", position)
}

// ToastComponent renders a single toast notification
func ToastComponent(toast *Toast) g.Node {
	levelClass := levelClasses[toast.Level]

	return Div(
		Class(fmt.Sprintf("toast %s", levelClass)),
		ID(toast.ID),
		CustomData("duration", fmt.Sprintf("%d", int(toast.Duration.Seconds()))),

		// Close button
		Button(
			Type("button"),
			Class("toast-close"),
			OnClick(fmt.Sprintf("removeToast('%s')", toast.ID)),
			g.Text("×"),
		),

		// Content
		Div(
			Class("toast-content"),
			If(toast.ShowIcon,
				Div(
					Class("toast-icon"),
					getIconForLevel(toast.Level),
				),
			),
			Div(
				Class("toast-message"),
				g.Text(toast.Message),
			),
		),

		Div(
			Class("toast-progress"),
		),
	)
}

// ToasterComponent renders the entire toaster container
func ToasterComponent(toaster *Toaster) g.Node {
	toasts := toaster.GetToasts()
	positionClass := GetPositionClass(toaster.GetPosition())

	return Div(
		ID("toaster"),
		Class(fmt.Sprintf("toaster %s", positionClass)),

		// Render each toast
		g.Group(g.Map(toasts, func(t *Toast) g.Node {
			return ToastComponent(t)
		})),

		// Include JS for toast functionality
		Script(g.Raw(toasterJS)),
	)
}

// ToasterCSS generates the CSS for the toaster
func ToasterCSS() g.Node {
	return Style(toasterCSS)
}

// getIconForLevel returns the SVG icon for a toast level
func getIconForLevel(level Level) g.Node {
	var svg string

	switch level {
	case SuccessLevel:
		svg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>`
	case ErrorLevel:
		svg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>`
	case WarningLevel:
		svg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>`
	case InfoLevel:
		svg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>`
	default:
		svg = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>`
	}

	return g.Raw(svg)
}

// OnClick generates an onclick attribute
func OnClick(js string) g.Node {
	return g.Attr("onclick", js)
}

// If conditionally includes a node
func If(condition bool, node g.Node) g.Node {
	if condition {
		return node
	}
	return g.Text("")
}

// CustomData adds a data-* attribute
func CustomData(key, value string) g.Node {
	return g.Attr(fmt.Sprintf("data-%s", key), value)
}
