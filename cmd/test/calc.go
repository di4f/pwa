package main

import (
	"github.com/di4f/pwa/app"
)

type CalcPage struct {
	app.Compo
}

func (page *CalcPage) Render() app.UI {
	button := func(text string) app.UI {
		return app.Button().Text(text)
	}

	row := func(ui ...app.UI) app.UI {
		return app.Div().Class("row").Body(
			ui...,
		)
	}

	ret := app.Main().Body(
		app.H2().Text("The calculator"),
		app.Form().Body(
			row(
				button("1"),
				button("2"),
				button("3"),
			),
			row(
				button("4"),
				button("5"),
				button("6"),
			),
			row(
				button("7"),
				button("8"),
				button("9"),
			),
		),
	)
	return FromRoot(ret)
}
