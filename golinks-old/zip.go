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
	"github.com/laughingcabbage/golinks/types/fs"
	"log"

	"github.com/urfave/cli"
)

func appZip(c *cli.Context) error {
	log.Println("verifying directories")
	dirPath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError("invalid source path", 0)
	}
	targetPath, err := verifyPath(c.Args().Get(1))
	if err != nil {
		return cli.NewExitError("invalid target path", 0)
	}

	log.Println("compressing file(s)")
	if err := fs.Compress(dirPath, targetPath); err != nil {
		return cli.NewExitError(err, 0)
	}

	log.Println("compressed archive to " + targetPath)
	return nil
}

func appUnzip(c *cli.Context) error {
	log.Println("verifying directories")
	archivePath, err := verifyPath(c.Args().First())
	if err != nil {
		return cli.NewExitError("invalid source path", 0)
	}
	targetPath, err := verifyPath(c.Args().Get(1))
	if err != nil {
		return cli.NewExitError("invalid target path", 0)
	}

	log.Println("decompressing file(s)")
	if err := fs.Decompress(archivePath, targetPath); err != nil {
		return cli.NewExitError("failed to decompress files", 0)
	}
	log.Println("decompressed archive to " + targetPath)

	return nil
}
