package multi

import "github.com/lazychanger/go-contrib/app"

type EnabledApplication interface {
	app.Application

	Enabled() bool
}
