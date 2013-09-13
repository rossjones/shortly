package main

import (
	"flag"
    "fmt"
    "runtime"
)


func init() {
	flag.Parse()
}

func usage() {
    fmt.Println("\n  Usage:")
}

func main() {

	if len(flag.Args()) == 0 {
        usage()
    }

    var command = flag.Arg(0)
    var arg = ''

    if len(flag.Args()) == 2 {
        arg = flag.Arg(1)
    }

	if command == "make-config" {
		err := WriteSampleConfig(arg)
		if err != nil {
			panic(err)
		}
        return
	} else command == "serve" {
        err := loadConfig(arg)
        if err != nil {
            panic(err)
        }
    } else {
        usage()
    }

    runtime.GOMAXPROCS(runtime.NumCPU())
    // Run the server!!

	println(configuration.Server.Host)
}
