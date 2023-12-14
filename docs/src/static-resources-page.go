package main

import (
	"github.com/di4f/pwa/analytics"
	"github.com/di4f/pwa/app"
)

type staticResourcesPage struct {
	app.Compo
}

func newStaticResourcePage() *staticResourcesPage {
	return &staticResourcesPage{}
}

func (p *staticResourcesPage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *staticResourcesPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *staticResourcesPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Images and Static Resources")
	ctx.Page().SetDescription("Documentation about how to deal with images and other static resources.")
	analytics.Page("static-resources", nil)
}

func (p *staticResourcesPage) Render() app.UI {
	return newPage().
		Title("Images and Static Resources").
		Icon(imgFolderSVG).
		Index(
			newIndexLink().Title("Intro"),
			newIndexLink().Title("Access static resources"),
			newIndexLink().Title("    In Handler"),
			newIndexLink().Title("    In components"),
			newIndexLink().Title("Setup Custom Web directory"),
			newIndexLink().Title("    Setup local web directory"),
			newIndexLink().Title("    Setup remote web directory"),

			app.Div().Class("separator"),

			newIndexLink().Title("Next"),
		).
		Content(
			newRemoteMarkdownDoc().Src("/web/documents/static-resources.md"),
		)
}
