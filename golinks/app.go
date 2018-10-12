/*
 *Copyright 2017-2018 Kevin Gentile
 *
 *Licensed under the Apache License, Version 2.0 (the "License");
 *you may not use this file except in compliance with the License.
 *You may obtain a copy of the License at
 *
 *http://www.apache.org/licenses/LICENSE-2.0
 *
 *Unless required by applicable law or agreed to in writing, software
 *distributed under the License is distributed on an "AS IS" BASIS,
 *WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *See the License for the specific language governing permissions and
 *limitations under the License.
 */

package main

import (
	"github.com/laughingcabbage/golinks/types/walker"
	"os"

	"github.com/urfave/cli"
)

const (
	//walkerWorkers int    = 1
	defaultRoot string = "r"
	//configName  string = "config"
)

var outputDir = defaultRoot

var w = walker.New(defaultRoot)

func main() {
	app := cli.NewApp()
	app.Name = "golinks"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "Kevin Gentile",
			Email: "kevin@kevingentile.com",
		},
	}
	app.Copyright = "(c) 2017-2018 Kevin Gentile"
	app.HelpName = "golinks"
	app.Usage = "a blockchain for your filesystem"

	app.Commands = []cli.Command{
		{
			Name:        "walk",
			Usage:       "walk a given archive",
			Description: "walk a given archive. print by default",
			Action:      appWalk,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "output, o",
					Usage:       "set output 'DIR'",
					Destination: &outputDir,
				},
				cli.BoolFlag{
					Name:  "print, p",
					Usage: "print walked archive",
				},
			},
		},
		{
			//todo save failing on link command
			Name:        "link",
			Usage:       "link [directory]",
			Description: "generate link file for the provided directory",
			Action:      appLink,
		},
		{
			Name:        "validate",
			Usage:       "validate [directory]",
			Description: "validate an existing archive link",
			Action:      appValidate,
		},
		{
			Name:        "zip",
			Usage:       "zip [directory] [target]",
			Description: "compress a directory into a zip file",
			Action:      appZip,
		},
		{
			Name:        "unzip",
			Usage:       "unzip [archive] [target]",
			Description: "decompress an archive to target folder",
			Action:      appUnzip,
		},
		{
			Name:        "maketest",
			Usage:       "build test directory",
			Description: "builds a test directory in the provided location",
			Action:      appBuildTestDir,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "small, s",
					Usage: "small directory",
				},
				cli.BoolFlag{
					Name:  "medium, m",
					Usage: "medium directory",
				},
				//todo usage "-large"
				cli.BoolFlag{
					Name:  "large, l",
					Usage: "large directory",
				},
			},
		},
	}

	//legooo
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
