/*
 *Copyright 2017 Kevin Gentile
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
	"github.com/urfave/cli"
)

func appWalk(c *cli.Context) error {
	//Verify provided walk path is valid
	path, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	//Construct new walker and begin execution
	w = walker.New(path)

	if err := w.Walk(); err != nil {
		return cli.NewExitError(err.Error(), 0)
	}

	//Process post walk commands
	for _, flagName := range c.FlagNames() {
		if !c.IsSet(flagName) {
			continue
		}
		switch flagName {
		case "print":
			w.PrintArchive()
		default:
			w.PrintArchive()
		}
	}
	return nil
}