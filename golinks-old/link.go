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
	"log"

	"github.com/laughingcabbage/golinks/types/blockmap"
	"github.com/urfave/cli"
)

func appLink(c *cli.Context) error {
	log.Println("running linker")
	absPath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	//todo this does not do what I thought it did.
	//todo
	blkmap := blockmap.New(absPath)
	log.Println("generating link in " + absPath)
	if err := blkmap.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}
	blkmap.PrintBlockMap()

	log.Println("saving blockmap to .link file")
	// if err := blockmap.Save(absPath); err != nil {
	// 	return cli.NewExitError(err, 0)
	// }

	//todo fs.saveGob
	return nil
}
