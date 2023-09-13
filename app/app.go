package app

import "context"

type Application interface {
	Init(ctx context.Context) error
	Run(ctx context.Context) error
	Reload(ctx context.Context) error
	Stop(ctx context.Context) error
}

type NamedApplication interface {
	Application
	Name() string
}

type TestApplication struct {
	IsEnable bool

	stop chan struct{}
}

func (t *TestApplication) Init(ctx context.Context) error {
	return nil
}

func (t *TestApplication) Run(ctx context.Context) error {
	<-t.stop

	return nil
}

func (t *TestApplication) Reload(ctx context.Context) error {
	if err := t.Stop(ctx); err != nil {
		return err
	}

	if err := t.Run(ctx); err != nil {
		return err
	}

	return nil
}

func (t *TestApplication) Stop(ctx context.Context) error {
	t.stop <- struct{}{}
	return nil
}
