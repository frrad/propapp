package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/BurntSushi/toml"
	"github.com/frrad/propapp/lib/counties"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	inPath := c.String("inpath")
	outPath := inPath

	dat, err := ioutil.ReadFile(inPath)
	if err != nil {
		return err
	}

	var countyData map[string]counties.State
	if _, err := toml.Decode(string(dat), &countyData); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}

	defer f.Close()
	w := bufio.NewWriter(f)
	enc := toml.NewEncoder(w)

	sortCounties(countyData)

	if err := enc.Encode(countyData); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

// sortCounties sorts counties in place for each state.
func sortCounties(states map[string]counties.State) {
	for _, stateData := range states {
		cs := stateData.Counties
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Name < cs[j].Name
		})
	}
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "inpath",
			Value: "counties.toml",
			Usage: "specify path to input",
		},
	}

	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
