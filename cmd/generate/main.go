package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/frrad/propapp/lib/counties"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	inPath := c.String("inpath")
	outPath := c.String("outpath")

	dat, err := ioutil.ReadFile(inPath)
	if err != nil {
		return err
	}

	var countyData counties.Counties
	if _, err := toml.Decode(string(dat), &countyData); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	err = writeHTML(w, countyData)
	if err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func writeHTML(w io.Writer, countyData counties.Counties) error {
	x, err := countyData.AsHTML()
	if err != nil {
		return err
	}
	fmt.Fprint(w, x)
	return nil
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "inpath",
			Value: "counties.toml",
			Usage: "specify path to input",
		},
		cli.StringFlag{
			Name:  "outpath",
			Value: "index.html",
			Usage: "specify path to output",
		},
	}

	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
