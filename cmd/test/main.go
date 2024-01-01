package main

import (
	"github.com/di4f/pwa/app"
	//"github.com/di4f/pwa/ui"
	"net/http"
	//"log"
)

type RootCompo struct {
	app.Compo
	children []app.UI
}

// Create the new root component that is on top of the rest.
func FromRoot(children ...app.UI) *RootCompo {
	ret := &RootCompo{
		children: children,
	}
	return ret
}

func (r *RootCompo) Render() app.UI {
	return app.Div().ID("root").Body(
		app.Header().Text("The example PWA application"),
		app.Nav().Body(
			app.Button().Text("Calculator").OnClick(func(c app.Context, e app.Event) {
				c.Navigate("/calc/")
			}),
			app.Button().Text("Note list").OnClick(func(c app.Context, e app.Event) {
				c.Navigate("/note-list/")
			}),
		),
		app.Range(r.children).Slice(func(i int) app.UI {
			return r.children[i]
		}),
	)
}

type RootPage struct {
	app.Compo
}

func (page *RootPage) Render() app.UI {
	ret := app.Main().Body(
		app.H2().Text("This is the root page, click something to move between them"),
	)
	return FromRoot(ret)
}

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	ret := app.Div().Body(
		app.H1().Text("Hello, World!"),
		app.Button().Text("To main").OnClick(func(c app.Context, e app.Event) {
			c.Navigate("/hello/")
		}),
	)
	return FromRoot(ret)
}

func main() {
	app.Route("/", &RootPage{})
	app.Route("/calc/", &CalcPage{})
	app.Route("/note-list/", &NoteListPage{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "The testing PWA application",
		Description: "The testing PWA application example",
		Styles: []string{
			"/web/hello.css",
		},
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
