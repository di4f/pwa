package ui

import (
	"github.com/di4f/pwa/app"
	//"log"
)

type RadioButtonDesc struct {
	Content app.UI
	Value   string
}

type RadioButtonCompo struct {
	app.Compo
	descs   []*RadioButtonDesc
	value   string
	onClick func(app.Context, app.Event)
}

func (b *RadioButtonCompo) Btn(
	value string,
	content app.UI,
) *RadioButtonCompo {
	b.descs = append(
		b.descs,
		&RadioButtonDesc{
			Content: content,
			Value:   value,
		},
	)

	return b
}

func (b *RadioButtonCompo) GetValue(
	c app.Context,
) string {
	return b.value
}

func (b *RadioButtonCompo) SetValue(
	c app.Context,
	value string,
) error {
	if b.value == value {
		return nil
	}

	b.value = value
	if b.Mounted() {
		b.Update()
	}

	return nil
}

// Like radio button but without the check shit in front
// or after the label and without JS interaction shit.
// Made mostly to be used as navigation buttons.
func (b *RadioButtonCompo) Render() app.UI {
	ret := app.Range(b.descs).Slice(func(i int) app.UI {
		desc := b.descs[i]

		btn := app.Button().
			OnClick(
				func(c app.Context, e app.Event) {
					b.SetValue(c, desc.Value)
					if b.onClick != nil {
						b.onClick(c, e)
					}
				},
			).Body(desc.Content)

		if desc.Value == b.value {
			btn = btn.Class("selected")
		}

		return btn
	})

	return app.Div().Body(
		ret,
	)
}

func RadioButton() *RadioButtonCompo {
	return &RadioButtonCompo{}
}

func (b *RadioButtonCompo) OnClick(
	fn func(app.Context, app.Event),
) *RadioButtonCompo {
	b.onClick = fn
	return b
}
