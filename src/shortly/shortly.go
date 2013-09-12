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
        // Bail
    }

	if len(flag.Args()) == 2 && flag.Arg(0) == "make-config" {
		err := WriteSampleConfig(flag.Arg(1))
		if err != nil {
			panic(err)
		}
        return
	} else if flag.Arg(0) == "serve" {
        if len(flag.Args()) == 2 {
            err := loadConfig("sample.cfg")
            if err != nil {
                panic(err)
            }
        } else {
            fmt.Println("Please specify the name of the config file")
            usage()
        }
    } else {
        usage()
    }

    runtime.GOMAXPROCS(runtime.NumCPU())
    // Run the server!!

	println(configuration.Server.Host)
}
