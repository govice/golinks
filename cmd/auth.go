package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cEmail string = "email"
	cToken string = "token"
)

var (
	setAuthEmail     string
	setAuthToken     string
	skipVerification bool
	authCmd          = &cobra.Command{
		Use:   "auth",
		Short: "setup authentication credentials",
		Run: func(cmd *cobra.Command, args []string) {
			if setAuthEmail != "" {
				viper.Set(cEmail, setAuthEmail)
			}

			if setAuthToken != "" {
				viper.Set(cToken, setAuthToken)
			}

			write := false
			if skipVerification {
				write = true
			} else {
				if ok, err := authorized(); err != nil {
					log.Fatal(err)
				} else if !ok {
					log.Fatal("Failed to authenticate user")
				}

				write = true
			}
			if write {
				if err := viper.WriteConfig(); err != nil {
					log.Fatal(err)
				}
			}

			cmd.Println("OK")
		},
	}
)

func authorized() (bool, error) {
	email := viper.Get(cEmail).(string)
	token := viper.Get(cToken).(string)
	authURL := viper.Get(cAuth).(string)

	userAuth := &UserAuth{
		Email: email,
		Token: token,
	}

	authJSON, err := json.Marshal(userAuth)
	if err != nil {
		return false, err
	}

	var buffer bytes.Buffer
	if _, err := buffer.Write(authJSON); err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", authURL, &buffer)
	if err != nil {
		return false, err
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

type UserAuth struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
