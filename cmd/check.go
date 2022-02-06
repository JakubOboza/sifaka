/*
Copyright © 2022 Jakub Oboza <jakub.oboza@gmail.com>

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
	"net/url"

	"github.com/JakubOboza/sifaka/certutils"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check when certain cert expires by url or file",
	Long: `Checks performs simple check for x509 certifcates
	or website or from file to show when it will expire`,
	Run: func(cmd *cobra.Command, args []string) {

		rawUrl, err := cmd.Flags().GetString("url")

		if err != nil {
			fmt.Println(err)
		} else if rawUrl != "" {

			parsedUrl, err := url.Parse(rawUrl)

			if err != nil {
				fmt.Println(err)
				return
			}

			cert, err := certutils.GetCertInfoForUrl(parsedUrl)

			if err != nil {
				fmt.Println(err)
				return
			}

			cert.DisplayInfo()

		}

		filePath, err := cmd.Flags().GetString("file")

		if err != nil {
			fmt.Println(err)
		} else if filePath != "" {

			rawData, err := ioutil.ReadFile(filePath)

			if err != nil {
				fmt.Println(err)
				return
			}

			cert, err := certutils.GetCertInfoForFile(rawData)

			if err != nil {
				fmt.Println(err)
				return
			}

			cert.DisplayInfo()

		}

		if filePath == "" && rawUrl == "" {
			fmt.Println("Please specify --url= or --file==")
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringP("url", "u", "", "url to https page eg: --url=https://lambdacu.be")
	checkCmd.Flags().StringP("file", "f", "", "file to x509 cert file eg: --file=MyImportantCert.cer")
}
