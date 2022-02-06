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

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/JakubOboza/sifaka/config"
	"github.com/JakubOboza/sifaka/display"
	"github.com/JakubOboza/sifaka/storage"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all certs in database by when they expire",
	Long: `lists all certs in database by when they expire.
	uses csv as a format of output`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := sqlx.Connect("sqlite3", config.DatabasePath())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		repo := storage.NewSqliteRepository(conn)
		storageService := storage.NewService(repo)
		certsDatums, err := storageService.AllByExpiration()

		if err != nil {
			fmt.Println(err)
			return
		}

		display.PrintCertificateDataList(&certsDatums)

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
