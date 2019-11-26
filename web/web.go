// Package web provides the user interface over http.
package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	stan "github.com/nats-io/stan.go"
)

const (
	headerStr = `<!DOCTYPE html>
<html>
<body>
<head>
  <title>{{.}}</title>
</head>
`
	bodyStr = `<body>
{{.}}
</body>
`
	footerStr = `</html>`
)

var (
	header, body, footer *template.Template
)

func init() {
	header = template.Must(template.New("Header").Parse(headerStr))
	body = template.Must(template.New("Body").Parse(bodyStr))
	footer = template.Must(template.New("Footer").Parse(footerStr))
}

// 10 OMIT

// RootHandler handles the root path.
type RootHandler struct {
	Msg string
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", h.Msg)
}

// 20 OMIT

// OrderHandler handles the order path.
type OrderHandler struct {
	stan.Conn // accessed using Conn
	Subject   string
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Publish(h.Subject, []byte("Hello from OrderHandler\n"))
	if err != nil {
		log.Printf("OrderHandler could not publish: %v", err)
	}
	fmt.Fprintf(w, "Published to %q in OrderHandler\n", h.Subject)
}

// 30 OMIT
