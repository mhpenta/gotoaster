package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	goaster "github.com/mhpena/gotoaster"
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add-toast", handleAddToast)

	fmt.Println("Server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Global toaster instance
var toaster = goaster.NewToaster().
	SetPosition(goaster.TopRight).
	SetMaxToasts(5)

func handleHome(w http.ResponseWriter, r *http.Request) {

	body := h.Body(
		h.Div(
			h.H1(g.Text("Goaster Gomponents Demo")),
			h.P(g.Text("A simplified toast notification system using gomponents")),

			h.Div(
				h.H2(g.Text("Actions")),
				h.Button(
					g.Attr("onclick", "fetch('/add-toast?level=success').then(() => window.location.reload())"),
					g.Text("Add Success Toast"),
				),
				g.Text(" "),
				h.Button(
					g.Attr("onclick", "fetch('/add-toast?level=error').then(() => window.location.reload())"),
					g.Text("Add Error Toast"),
				),
				g.Text(" "),
				h.Button(
					g.Attr("onclick", "fetch('/add-toast?level=warning').then(() => window.location.reload())"),
					g.Text("Add Warning Toast"),
				),
				g.Text(" "),
				h.Button(
					g.Attr("onclick", "fetch('/add-toast?level=info').then(() => window.location.reload())"),
					g.Text("Add Info Toast"),
				),
			),

			// Render the toaster component
			goaster.ToasterComponent(toaster),
		),
	)

	page := c.HTML5(c.HTML5Props{
		Title:       "Goaster Gomponents Example",
		Description: "Goaster Gomponents Example",
		Language:    "en",
		Head: []g.Node{
			g.Raw(`<style>
.toaster {
  position: fixed;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  z-index: 9999;
  padding: 1rem;
}

.toaster-top-right {
  top: 0;
  right: 0;
}

.toaster-top-left {
  top: 0;
  left: 0;
}

.toaster-top-center {
  top: 0;
  left: 50%;
  transform: translateX(-50%);
}

.toaster-bottom-right {
  bottom: 0;
  right: 0;
}

.toaster-bottom-left {
  bottom: 0;
  left: 0;
}

.toaster-bottom-center {
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
}

.toast {
  position: relative;
  display: flex;
  flex-direction: column;
  min-width: 300px;
  max-width: 400px;
  background-color: white;
  border-radius: 0.375rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  overflow: hidden;
  animation: toast-in 0.3s ease forwards;
}

.toast-close {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  width: 1.5rem;
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.25rem;
  line-height: 1;
  border: none;
  background: transparent;
  color: rgba(0, 0, 0, 0.5);
  cursor: pointer;
  transition: color 0.2s;
}

.toast-close:hover {
  color: rgba(0, 0, 0, 0.8);
}

.toast-content {
  display: flex;
  padding: 1rem;
  gap: 0.75rem;
  align-items: center;
}

.toast-icon {
  flex-shrink: 0;
  color: currentColor;
}

.toast-message {
  flex-grow: 1;
}

.toast-progress {
  height: 4px;
  background-color: rgba(0, 0, 0, 0.1);
  width: 100%;
  position: absolute;
  bottom: 0;
  left: 0;
  transform-origin: left;
  animation: toast-progress linear forwards;
}

.toast-success {
  border-left: 4px solid #10b981;
  color: #065f46;
}

.toast-error {
  border-left: 4px solid #ef4444;
  color: #991b1b;
}

.toast-warning {
  border-left: 4px solid #f59e0b;
  color: #92400e;
}

.toast-info {
  border-left: 4px solid #3b82f6;
  color: #1e40af;
}

.toast-default {
  border-left: 4px solid #6b7280;
  color: #1f2937;
}

@keyframes toast-in {
  from {
    transform: translateY(100%);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes toast-progress {
  0% {
    transform: scaleX(1);
  }
  100% {
    transform: scaleX(0);
  }
}</style>`),
			goaster.ToasterCSS(),
		},
		Body: []g.Node{body},
	})

	_ = page.Render(w)
}

func handleAddToast(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")

	var toastLevel goaster.Level
	switch level {
	case "success":
		toastLevel = goaster.SuccessLevel
	case "error":
		toastLevel = goaster.ErrorLevel
	case "warning":
		toastLevel = goaster.WarningLevel
	case "info":
		toastLevel = goaster.InfoLevel
	default:
		toastLevel = goaster.DefaultLevel
	}

	// Create a message based on the level
	message := fmt.Sprintf("This is a %s toast created at %s",
		toastLevel,
		time.Now().Format("15:04:05"),
	)

	// Create the toast and add it to the toaster
	toast := goaster.NewToastWithDuration(message, toastLevel, 2*time.Second)
	toaster.AddToast(toast)

	w.WriteHeader(http.StatusOK)
}
