//go:build !windows

package grace

import (
	"context"
	"github.com/lazychanger/go-contrib/app"
	"github.com/lazychanger/go-contrib/zaplog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func GracefulRun(ctx context.Context, application app.Application) error {
	defer func() {
		if err := recover(); err != nil {
			zaplog.Warnln(err)
			debug.PrintStack()
		}
	}()

	if err := application.Init(ctx); err != nil {
		return err
	}

	sigs := make(chan os.Signal)
	done := make(chan error)
	defer func() { close(sigs); close(done) }()

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go func() {
		done <- application.Run(ctx)
	}()

	for {
		select {
		case sig := <-sigs:
			zaplog.Infof("got signal: %s", sig.String())

			switch sig {
			case syscall.SIGHUP:
				// reload failed
				if err := application.Reload(ctx); err != nil {
					return err
				}
			default:
				subCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))

				if err := application.Stop(subCtx); err != nil {
					zaplog.Warnf("application run stop failed: %s", err)
				}

				cancel()
				return <-done
			}

		case err := <-done:
			return err
		}
	}
}
