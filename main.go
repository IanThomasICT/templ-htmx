package main

import (
	"context"
	"os"

	t "github.com/ianthomasict/templ-htmx/templates"
)


func main() {
    component := t.Greeting("Ian");
    component.Render(context.Background(), os.Stdout)
}
