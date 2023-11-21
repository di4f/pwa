package main

import (
	"github.com/omnipunk/pwa/v9/app"
	"github.com/omnipunk/pwa/v9/ui"
	"net/http"
	"log"
)

type Root struct {
	app.Compo
}

func (r *Root) OnPreRender(c app.Context) {
	log.Println("prerendering root")
}

func (r *Root) Render() app.UI {
	radio := ui.RadioButton()
	nav := ui.Nav(nil)

	radio.Btn("test1", app.Text("Test1")).
		Btn("test2", app.Text("Test2")).
		Btn("test3", app.Text("Test3")).
		Btn("sub", app.Text("Subs")).
		OnClick(func(c app.Context, e app.Event){
			nav.SetValue(c, radio.GetValue(c))
		})

	nav.Valuer(radio).Default(
			"test1",
		).Define(
			"test1",
			app.P().Text("Test 1"),
		).Define(
			"test2",
			app.Div().Text("Test 2"),
		).Define(
			"test3",
			app.Strong().Text("Test 3"),
		).Define(
			"sub",
			&RootSub{
				Parent: nav,
			},
		)

	return app.Div().ID("root").Body(
		app.Header().Text("Hello, World!"),
		app.Nav().Body(
			radio,
		),
		nav,
	)
}

type RootSub struct {
	app.Compo
	Parent *ui.NavCompo
}

func (rs *RootSub) Render() app.UI {
	nav := ui.Nav(rs.Parent)
	radio := ui.RadioButton()


	nav.Valuer(radio).Default(
			"sub1",
		).Define(
			"sub1",
			app.Text("Sub1"),
		).Define(
			"sub2",
			app.Text("Sub2"),
		).Define(
			"sub3",
			app.Text("Sub3"),
		)

	radio.Btn("sub1", app.Text("Sub1")).
		Btn("sub2", app.Text("Sub2")).
		Btn("sub3", app.Text("Sub3")).
		OnClick(func(c app.Context, e app.Event){
			nav.SetValue(c, radio.GetValue(c))
		})

	return app.Div().Body(
		app.Nav().Body(
			radio,
		),
		nav,
	)
}

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	ret := app.Div().Body(
		app.H1().Text("Hello, World!"),
		app.Button().Text("To main").OnClick(func(c app.Context, e app.Event){
			c.Navigate("/")
		}),
	)
	return ret
}

func main() {
	app.Route("/hello/", &hello{})
	app.RouteWithRegexp(".*", &Root{})
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

