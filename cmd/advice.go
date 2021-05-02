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
	"net/http"

	"github.com/spf13/cobra"
)

// adviceCmd represents the advice command
var adviceCmd = &cobra.Command{
	Use:   "advice",
	Short: "Random Advice",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomAdvice()
	},
}

func init() {
	rootCmd.AddCommand(adviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// type advice struct {
// 	ID     int    `json:"slip_id"`
// 	Advice string `json:"advice"`
// }

// func getRandomAdvice() {
// 	url := "https://api.adviceslip.com/advice"
// 	responseBytes := getAdvice(url)
// 	advice := advice{}
// 	err := json.Unmarshal(responseBytes, &advice)
// 	errorHandling(err)
// 	fmt.Println(advice.Advice)
// }

// func getAdvice(baseAPI string) []byte {
// 	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
// 	errorHandling(err)
// 	response, erro := http.DefaultClient.Do(request)
// 	errorHandling(erro)
// 	//fmt.Println(response)
// 	responseBytes, ero := ioutil.ReadAll(response.Body)
// 	errorHandling(ero)
// 	//fmt.Println(responseBytes)

// 	return responseBytes
// }

type AdviceSlip struct {
	Advice string `json:"advice"`
	ID     string `json:"slip_id"`
}

type RandomAdviceResp struct {
	Slip *AdviceSlip `json:"slip"`
}

func getRandomAdvice() {
	url := "http://api.adviceslip.com/advice"
	response, err := getAdvice(url)
	errorHandling(err)
	fmt.Println(response)
}

func getAdvice(url string) (interface{}, error) {

	resp, err := http.Get(url)
	errorHandling(err)
	random := true
	var decoded interface{}
	if random {
		decoded = &RandomAdviceResp{}
	}
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	errorHandling(err)
	advice := "No advice found :'("
	if random {
		slip := decoded.(*RandomAdviceResp).Slip
		if slip != nil {
			advice = slip.Advice
		}

	}
	return advice, nil
}
