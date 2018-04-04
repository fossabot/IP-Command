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
	var check_g bool
	var check_l bool
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "local ip, l",
			Usage:       "local ip",
			Destination: &check_l,
		},
		cli.BoolFlag{
			Name:        "global ip, g",
			Usage:       "global ip",
			Destination: &check_g,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			fmt.Println("argc")
		}

		if check_g == true {
			getglobalip()
			return nil
		}

		if check_l == true {
			getmyip()
			return nil
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
