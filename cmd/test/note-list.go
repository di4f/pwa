package main
import (
	"github.com/di4f/pwa/app"
	"log"
)

var (
	UserNotes NoteListCompo
)

type NoteListPage struct {
	app.Compo
}

func (page *NoteListPage) Render() app.UI {
	form := &NoteListForm{}
	form.List = &UserNotes
	ret := app.Main().Body(
		app.H2().Text("Note list"),
		form,
	)
	return FromRoot(ret)
}

type NoteListCompo struct {
	app.Compo
	Notes []string
}
func (compo *NoteListCompo) Render() app.UI {
	return app.Ul().Body(
		app.Range(compo.Notes).Slice(func(i int) app.UI {
			text := compo.Notes[i]
			return app.Li().Text(text)
		}),
	)
}

type NoteListForm struct {
	app.Compo
	List *NoteListCompo
}

func (compo *NoteListForm) Render() app.UI {
	var input app.UI
	onSubmit := func(c app.Context, e app.Event){
		e.PreventDefault()
		jsInput := input.JSValue()
		val := jsInput.Get("value").String()
		log.Printf("%q", val)
		if val == "" {
			return
		}

		compo.List.Notes = append(
			compo.List.Notes,
			val,
		)
		jsInput.Set("value", "")
		jsInput.Call("focus")
	}
	input = app.Input().
		Type("text").
		AutoFocus(true)
		//OnSubmit(onSubmit)

	return app.Form().OnSubmit(onSubmit).Body(
		input,
		app.Input().Type("button").OnClick(onSubmit).Value("Submit"),
		compo.List.Render(),
	)
}

