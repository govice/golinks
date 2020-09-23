package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/govice/golinks/block"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var (
	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "push staged link file",
		Run: func(cmd *cobra.Command, args []string) {
			stagePath := viper.Get(cStagingPath).(string)
			pushRoute := viper.Get(cRemote).(string)
			userEmail := viper.Get(cEmail).(string)
			userToken := viper.Get(cToken).(string)
			stagedFiles, err := ioutil.ReadDir(stagePath)
			if err != nil {
				cli.NewExitError(err, 1)
			}

			for _, info := range stagedFiles {
				verb("pushing staged file: " + info.Name())
				filePath := filepath.Join(stagePath, info.Name())
				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Fatal(err)
				}

				data := &pushData{
					Data: fileData,
				}

				dataJSON, err := json.Marshal(data)
				if err != nil {
					log.Fatal(err)
				}

				var buffer bytes.Buffer
				if _, err := buffer.Write(dataJSON); err != nil {
					log.Fatal(err)
				}

				req, err := http.NewRequest("POST", pushRoute+"/api/chain", &buffer)
				if err != nil {
					log.Fatal(err)
				}

				req.Header.Add("Accept", "application/json")

				q := req.URL.Query()
				q.Add("email", userEmail)
				q.Add("token", userToken)

				req.URL.RawQuery = q.Encode()

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()

				if res.StatusCode != http.StatusOK {
					log.Fatal(errors.New("push failed"))
				}

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}

				block := &block.Block{}
				if err := json.Unmarshal(body, block); err != nil {
					log.Fatal(err)
				}

				fmt.Println("Index: ", block.Index)
				fmt.Println("Hash: ", base64.StdEncoding.EncodeToString(block.BlockHash))
				fmt.Println("Parent: ", base64.StdEncoding.EncodeToString(block.ParentHash))
			}
		},
	}
)

type pushData struct {
	Data []byte `json:"data"`
}
