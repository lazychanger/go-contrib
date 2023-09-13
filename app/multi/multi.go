package multi

import (
	"context"
	"github.com/lazychanger/go-contrib/app"
	"github.com/lazychanger/go-contrib/zaplog"
	"golang.org/x/sync/errgroup"
	"reflect"
)

type multi struct {
	apps []app.Application
}

func (m *multi) Init(basectx context.Context) error {
	group := &errgroup.Group{}

	for _, _app := range m.apps {
		cur := _app
		group.Go(func() error {
			subctx, cancel := context.WithCancel(basectx)
			defer cancel()

			return cur.Init(subctx)
		})
	}
	return group.Wait()
}

func (m *multi) Run(basectx context.Context) error {
	group := &errgroup.Group{}

	for _, _app := range m.apps {
		cur := _app

		group.Go(func() error {
			if _, ok := cur.(*multi); !ok {
				name := reflect.TypeOf(cur).String()
				if namedApp, ok := cur.(app.NamedApplication); ok {
					name = namedApp.Name()
				}

				zaplog.Debugf("%s is running", name)

				defer func() {
					zaplog.Debugf("%s has been done", name)
				}()
			}

			subctx, cancel := context.WithCancel(basectx)
			defer cancel()

			return cur.Run(subctx)
		})
	}

	return group.Wait()
}

func (m *multi) Reload(basectx context.Context) error {
	group := &errgroup.Group{}

	for _, _app := range m.apps {
		cur := _app

		group.Go(func() error {
			subctx, cancel := context.WithCancel(basectx)
			defer cancel()

			return cur.Reload(subctx)
		})
	}
	return group.Wait()
}

func (m *multi) Stop(basectx context.Context) error {
	group := &errgroup.Group{}

	for _, _app := range m.apps {
		cur := _app
		group.Go(func() error {
			subctx, cancel := context.WithCancel(basectx)
			defer cancel()

			return cur.Stop(subctx)
		})
	}
	return group.Wait()
}

func (m *multi) Enabled() bool {
	return true
}

func New(apps ...app.Application) app.Application {

	sortApps := &sortApplications{apps: make([]app.Application, 0)}

	for _, application := range apps {
		if enabledApp, ok := application.(EnabledApplication); !ok || !enabledApp.Enabled() {
			continue
		}
		sortApps.apps = append(sortApps.apps, application)
	}

	return &multi{
		apps: sortApps.Applications(),
	}
}
