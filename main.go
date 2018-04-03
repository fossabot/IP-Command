package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var check_arg string
	var check_bool bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "global ip, g",
			Value:       "global ip",
			Usage:       "global ip",
			Destination: &check_arg,
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "local ip, l",
			Usage:       "local ip",
			Destination: &check_bool,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			fmt.Println("argc")
		}

		if check_bool == true {
			getglobalip()
			return nil
		}

		if check_arg == "a" {
			fmt.Println("wait")
		} else {
			fmt.Println("your local ip")
			getmyip()
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getmyip() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}

func getglobalip() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
