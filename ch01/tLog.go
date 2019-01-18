package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	Sugar.Info("nihao")
	Sugar.Warn("nihao")
	Sugar.Error("nihao")
	log.Fatal("llllll")
	Sugar.Fatal("nihao")
	//Sugar.F
	fmt.Print("nihaffffo")
	Sugar.Panic("nihao")
}

// Sugar is a zap sugared logger
var Sugar *zap.SugaredLogger

func init() {
	var err error
	logger, err := zap.NewProduction(
		zap.AddStacktrace(zapcore.PanicLevel),
	)
	if err != nil {
		log.Fatalln("failed to initialize zap logger")
	}
	Sugar = logger.Sugar()
}
