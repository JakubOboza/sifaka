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
	"path/filepath"

	"github.com/JakubOboza/sifaka/certutils"
	"github.com/spf13/cobra"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/JakubOboza/sifaka/config"
	"github.com/JakubOboza/sifaka/storage"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add check and cert/url to database",
	Long:  `checks and adds cert or url to database of sifaka`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := sqlx.Connect("sqlite3", config.DatabasePath())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		repo := storage.NewSqliteRepository(conn)
		storageService := storage.NewService(repo)

		if err != nil {
			fmt.Println(err)
			return
		}

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

			createCertData := &storage.CreateCertificateData{
				Name:     rawUrl,
				Subject:  cert.Subject,
				Type:     cert.Type,
				NotAfter: cert.NotAfter,
			}

			certData, err := storageService.Store(createCertData)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Cert added with ID:", certData.ID)

		}

		certFilePath, err := cmd.Flags().GetString("file")

		if err != nil {
			fmt.Println(err)
		} else if certFilePath != "" {

			rawData, err := ioutil.ReadFile(certFilePath)

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

			filename := filepath.Base(certFilePath)

			createCertData := &storage.CreateCertificateData{
				Name:     filename,
				Subject:  cert.Subject,
				Type:     cert.Type,
				NotAfter: cert.NotAfter,
			}

			certData, err := storageService.Store(createCertData)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Cert added with ID:", certData.ID)

		}

		if certFilePath == "" && rawUrl == "" {
			cmd.Help()
			fmt.Println("Please specify --url= or --file==")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("url", "u", "", "url to https page eg: --url=https://lambdacu.be")
	addCmd.Flags().StringP("file", "f", "", "file to x509 cert file eg: --file=MyImportantCert.cer")
}
