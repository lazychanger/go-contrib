//go:build windows

package grace

import (
	"context"
	"github.com/lazychanger/go-contrib/app"
	"time"
)

func GracefulRun(ctx context.Context, application app.Application) error {
	stop := make(chan error)
	done := make(chan error)
	go func() {
		done <- application.Run(ctx)
	}()

	<-stop

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
	defer cancel()
	application.Stop(ctx)

	return <-done
}
