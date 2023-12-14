package main

import (
	"github.com/di4f/pwa/analytics"
	"github.com/di4f/pwa/app"
)

type referencePage struct {
	app.Compo
}

func newReferencePage() *referencePage {
	return &referencePage{}
}

func (p *referencePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *referencePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *referencePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Reference for building PWA with Go and WASM")
	ctx.Page().SetDescription("Go-app API reference for building Progressive Web Apps (PWA) with Go (Golang) and WebAssembly (WASM).")
	analytics.Page("reference", nil)
}

func (p *referencePage) Render() app.UI {
	return newPage().
		Title("Reference").
		Icon(golangSVG).
		Index(
			app.A().
				Class("index-link").
				Class(fragmentFocus("pkg-overview")).
				Href("#pkg-overview").
				Text("Overview"),
			newReferenceContent().
				Class("reference-index").
				Index(true),
			app.Div().Class("separator"),
		).
		Content(
			newReferenceContent().Class("reference"),
		)
}
