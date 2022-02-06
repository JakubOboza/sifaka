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
	"github.com/JakubOboza/sifaka/storage"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes given cert from database list",
	Long: `remove cert from the list by id. eg --id=69 will remove cert with id of 69.
	check the 'list' command to get the id for cert you wanna remove.
	No output means classic unix 'ok everything went well'`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := cmd.Flags().GetInt("id")

		if err != nil {
			fmt.Println(err)
			return
		}

		if id < 1 {
			fmt.Println("id can't be less than 1")
			return
		}

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

		certData, err := storageService.FindById(id)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = storageService.Destroy(certData)

		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().IntP("id", "i", 0, "id of certifacte data you want to remove")
}
