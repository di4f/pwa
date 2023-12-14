package main

import (
	"github.com/omnipunk/pwa/v9/app"
	"github.com/omnipunk/pwa/v9/ui"
	"net/http"
	//"log"
)

type RootCompo struct {
	app.Compo
	children []app.UI
}

func Root(children ...app.UI) *RootCompo {
	ret := &RootCompo{
		children: children,
	}
	return ret
}

func (r *RootCompo) Render() app.UI {
	return app.Div().ID("root").Body(
		app.Header().Text("Hello, World!"),
		app.Nav().Body(
			app.Button().Text("Subs").OnClick(func(c app.Context, e app.Event){
				c.Navigate("/test1/")
			}),
			app.Button().Text("Root/Hello").OnClick(func(c app.Context, e app.Event){
				c.Navigate("/")
			}),
		),
		app.Range(r.children).Slice(func(i int) app.UI {
			return r.children[i]
		}),
	)
}

type RootSub struct {
	app.Compo
	Parent *ui.NavCompo
}

func (rs *RootSub) Render() app.UI {
	return Root(
		app.Div().Text("The sub text"),
	)
}

type Text struct {
	app.Compo
	Text string
}

func (t *Text) Render() app.UI {
	return nil
}

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	ret := app.Div().Body(
		app.H1().Text("Hello, World!"),
		app.Button().Text("To main").OnClick(func(c app.Context, e app.Event){
			c.Navigate("/hello/")
		}),
	)
	return Root(ret)
}

func main() {
	app.Route("/", &hello{})
	app.Route("/hello/", &hello{})
	app.Route("/test1/", &RootSub{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name: "Hello",
		Description: "The 'Hello, World' example",
		Styles: []string{
			"/web/hello.css",
		},
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

