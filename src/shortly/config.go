package main

import (
	"bufio"
	"code.google.com/p/gcfg"
	"fmt"
	"os"
)

type Config struct {
	App struct {
		Debug bool
	}
	Server struct {
		Host string
		Bind string
		Port int
	}
	Database struct {
		Url string
	}
	Templates struct {
		Path string
	}
}

var configuration Config

func loadConfig(filename string) error {
	return gcfg.ReadFileInto(&configuration, filename)
}

func WriteSampleConfig(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, SAMPLE_CONFIG)
	return w.Flush()
}

var SAMPLE_CONFIG = `
# This is a sample config file for shortly.  You will need
# to change nearly all of the values to get a workable system.

[app]
debug = false

[server]
host = localhost  # The hostname that we are serving at
port = 2112       # The port to serve content on

[database]
url =    # The URL to the postgres database

[templates]
path = /opt/shortly/shortly-templates
`
