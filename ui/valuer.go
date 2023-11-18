package ui

import (
	"github.com/omnipunk/pwa/v9/app"
)

type Valuer[T any] interface {
	SetValue(app.Context, T) error
	GetValue(app.Context) T
}

