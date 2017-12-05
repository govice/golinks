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

	"github.com/LaughingCabbage/goLinks/types/blockmap"
	"github.com/urfave/cli"
)

func appValidate(c *cli.Context) error {
	//Validate provided path
	path, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError(err, 0)
	}

	//Load blockmap from existing file
	log.Println("checking for existing link file")
	b := &blockmap.BlockMap{}
	if err := b.Load(path); err != nil {
		return cli.NewExitError("link file not found", 0)
	}

	//Validate the existing directory
	log.Println("validating link file with current archive")
	temp := blockmap.New(path)
	if err := temp.Generate(); err != nil {
		return cli.NewExitError(err, 0)
	}

	//Compare file with existing directory
	if !blockmap.Equal(b, temp) {
		return cli.NewExitError("link is invalid", 0)
	}

	log.Println("link is valid")
	return nil
}
