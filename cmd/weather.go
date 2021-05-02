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
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/lunixbochs/vtclean"
	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getWeather(args)
		errorHandling(err)
		fmt.Println(response)

	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getWeather(args []string) (interface{}, error) {
	where := args[0]

	req, err := http.NewRequest("GET", "http://wttr.in/"+where+"?m", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "curl/7.49.1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// remove escape sequences
	unescaped := vtclean.Clean(string(body), false)

	split := strings.Split(string(unescaped), "\n")

	// Show both celcius and fahernheit
	for i, v := range split {
		if !strings.Contains(v, "°C") {
			continue
		}

		tmpFrom := 0
		tmpTo := 0
		isRange := false
		var TempRangeRegex = regexp.MustCompile("(-?[0-9]{1,3})( ?- ?(-?[0-9]{1,3}))? ?°C")
		submatches := TempRangeRegex.FindStringSubmatch(v)
		if len(submatches) < 2 {
			continue
		}

		tmpFrom, _ = strconv.Atoi(submatches[1])

		if len(submatches) >= 4 && submatches[3] != "" {
			tmpTo, _ = strconv.Atoi(submatches[3])
			isRange = true
		}

		// convert to fahernheit
		tmpFrom = int(float64(tmpFrom)*1.8 + 32)
		tmpTo = int(float64(tmpTo)*1.8 + 32)

		v = strings.TrimRight(v, " ")
		if isRange {
			split[i] = v + fmt.Sprintf(" (%d-%d °F)", tmpFrom, tmpTo)
		} else {
			split[i] = v + fmt.Sprintf(" (%d °F)", tmpFrom)
		}
	}

	out := "```\n"
	for i := 0; i < 7; i++ {
		if i >= len(split) {
			break
		}
		out += strings.TrimRight(split[i], " ") + "\n"
	}
	out += "\n```"

	return out, nil
}
