package ui

import (
	"github.com/di4f/pwa/app"
)

type Valuer[T any] interface {
	SetValue(app.Context, T) error
	GetValue(app.Context) T
}
