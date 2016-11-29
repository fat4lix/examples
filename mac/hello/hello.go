package main

import "github.com/murlokswarm/app"

// Hello implements app.Componer interface.
type Hello struct {
	Greeting string
}

// Render returns the HTML markup that describes the appearance of the
// component.
// It supports standard HTML and extends it slightly to handle other component
// declaration or Golang callbacks.
// Can be templated following rules from https://golang.org/pkg/text/template.
func (h *Hello) Render() string {
	return `
<div class="WindowLayout">    
    <div class="HelloBox">
        <h1>
            Hello,
          <span>
                {{if .Greeting}}
                    {{html .Greeting}}
                {{else}}
                    World
                {{end}}
            </span>
        </h1>
        <input type="text"
               value="{{html .Greeting}}"
               placeholder="What is your name?"
               autofocus="true"
               _onchange="OnInputChange" />
    </div>
</div>
    `
}

// OnInputChange is the handler called when an onchange event occurs.
// In the HTML markup, a Go component method is target by prefixing the event with "_".
// eg _onchange.
func (h *Hello) OnInputChange(arg app.ChangeArg) {
	h.Greeting = arg.Value // Changing the greeting.
	app.Render(h)          // Tells the app to update the rendering of the component.
}

func init() {
	// Registers the Hello component.
	// Allows the app to create a Hello component when it finds its declaration
	// into a HTML markup.
	app.RegisterComponent(&Hello{})
}
