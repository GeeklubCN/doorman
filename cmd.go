package main

import "flag"

type Cmd struct {
	configPath string
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.configPath, "c", "./conf/config.yaml", "Path to the configuration filename")
	flag.Parse()
	return cmd
}
