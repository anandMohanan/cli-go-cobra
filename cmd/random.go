/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJoke(url)
	joke := joke{}
	err := json.Unmarshal(responseBytes, &joke)
	errorHandling(err)

	fmt.Println(joke.Joke)

}

func getJoke(baseAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	errorHandling(err)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke cli")

	response, ero := http.DefaultClient.Do(request)
	errorHandling(ero)
	responseBytes, erro := ioutil.ReadAll(response.Body)
	errorHandling(erro)

	return responseBytes
}

// func getRandomJoke() {
// 	url := "https://icanhazdadjoke.com/"
// 	response, err := http.Get(url)
// 	errorHandling(err)
// 	body := response.Body
// 	joke := joke{
// 		ID:     body.id,
// 		Joke:   body.joke,
// 		Status: body.status,
// 	}
// 	fmt.Println(joke.joke)

// }
