package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	gotoaster "github.com/mhpena/gotoaster"
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
var toaster = gotoaster.NewToaster().
	SetPosition(gotoaster.TopRight).
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
			gotoaster.ToasterComponent(toaster),
		),
	)

	page := c.HTML5(c.HTML5Props{
		Title:       "Goaster Gomponents Example",
		Description: "Goaster Gomponents Example",
		Language:    "en",
		Head: []g.Node{
			gotoaster.ToasterCSS(),
			h.Link(h.Rel("stylesheet"), h.Href("/styles.css")),
		},
		Body: []g.Node{body},
	})

	_ = page.Render(w)
}

func handleAddToast(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")

	var toastLevel gotoaster.Level
	switch level {
	case "success":
		toastLevel = gotoaster.SuccessLevel
	case "error":
		toastLevel = gotoaster.ErrorLevel
	case "warning":
		toastLevel = gotoaster.WarningLevel
	case "info":
		toastLevel = gotoaster.InfoLevel
	default:
		toastLevel = gotoaster.DefaultLevel
	}

	// Create a message based on the level
	message := fmt.Sprintf("This is a %s toast created at %s",
		toastLevel,
		time.Now().Format("15:04:05"),
	)

	// Create the toast and add it to the toaster
	toast := gotoaster.NewToastWithDuration(message, toastLevel, 2*time.Second)
	toaster.AddToast(toast)

	w.WriteHeader(http.StatusOK)
}
