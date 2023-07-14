# Logger


## Install

```shell
go get -u github.com/lazychanger/go-contrib/logger
```

## How to use

```go
package main

import (
	log "github.com/lazychanger/go-contrib/logger"
	// zerolog driver
	_ "github.com/lazychanger/go-contrib/logger/zerolog"
)

func main() {
	log.SetLevel(log.InfoLevel)
	// default is json format 
	log.SetFormat(log.Color)

	log.Trace("Hello Trace")
	log.Debug("Hello Debug")
	log.Info("Hello Info")
	log.Warn("Hello Warn")
	log.Error("Hello Error")
	log.Fatal("Hello Fatal")
	log.Panic("Hello Panic")
}

```


## Todo

** More driver **
- [x] [zerolog](https://github.com/rs/zerolog)
- [x] [cloudwego/hlog](github.com/cloudwego/hertz/pkg/common/hlog)
- [ ] zap
- [ ] logrus