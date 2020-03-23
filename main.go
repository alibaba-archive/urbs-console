package main

import (
	"context"
	"flag"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/api"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/util"
)

var help = flag.Bool("help", false, "show help info")
var version = flag.Bool("version", false, "show version info")

func main() {
	flag.Parse()
	if *help || *version {
		util.PrintVersion()
	}

	app := api.NewApp()
	ctx := gear.ContextWithSignal(context.Background())
	host := "http://" + conf.Config.SrvAddr
	if conf.Config.CertFile != "" && conf.Config.KeyFile != "" {
		host = "https://" + conf.Config.SrvAddr
	}

	logger.Default.Infof("Urbs-Console start %s", host)
	logger.Default.Err("Urbs-Console closed %v", app.ListenWithContext(
		ctx, conf.Config.SrvAddr, conf.Config.CertFile, conf.Config.KeyFile))
}
