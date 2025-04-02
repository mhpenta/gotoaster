# GoToaser

A lightweight toast notification system for Go web applications, built with [gomponents](https://github.com/maragudk/gomponents). Inspired by [goaster](https://github.com/indaco/goaster).

## Features

- Simple API for creating toast notifications
- Multiple notification levels (success, error, warning, info)
- Flexible positioning options
- Automatic removal with configurable durations
- Thread-safe for concurrent usage
- Built with gomponents for declarative HTML generation

## Installation

```bash
go get github.com/mhpenta/gotoaster
```

## Quick Start

```go
package main

import (
    "net/http"

    "github.com/mhpenta/gotoaster"
    g "maragu.dev/gomponents"
    h "maragu.dev/gomponents/html"
)

func main() {
    http.HandleFunc("/", handleHome)
    http.ListenAndServe(":8080", nil)
}

// Create a toaster
var toaster = gotoaser.NewToaster().
    SetPosition(gotoaser.TopRight)

func handleHome(w http.ResponseWriter, r *http.Request) {
    // Add a toast
    toast := gotoaser.NewToast("Hello World!", gotoaser.SuccessLevel)
    toaster.AddToast(toast)
    
    // Render page with toaster
    page := h.HTML(
        h.Head(
            h.Title("GoToaser Example"),
            gotoaser.ToasterCSS(),
        ),
        h.Body(
            h.Div(
                h.H1(g.Text("GoToaser Demo")),
                gotoaser.ToasterComponent(toaster),
            ),
        ),
    )
    
    page.Render(w)
}
```

## Configuration

```go
// Create a new toaster
toaster := gotoaser.NewToaster()

// Customize
toaster.SetPosition(gotoaser.BottomCenter)
toaster.SetMaxToasts(3)
toaster.SetAutoRemove(true)
```

## Toast Levels

- `DefaultLevel`
- `SuccessLevel`
- `ErrorLevel`
- `WarningLevel`
- `InfoLevel`

## Positions

- `TopRight` (default)
- `TopLeft`
- `TopCenter`
- `BottomRight`
- `BottomLeft`
- `BottomCenter`

## License

MIT License - Copyright (c) github.com/mhpenta