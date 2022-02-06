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
	"github.com/JakubOboza/sifaka/notifications"
	"github.com/JakubOboza/sifaka/server"
	"github.com/JakubOboza/sifaka/storage"
	"github.com/JakubOboza/sifaka/watcher"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")

		if err != nil {
			fmt.Println(err)
			return
		}

		conn, err := sqlx.Connect("sqlite3", config.DatabasePath())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		repo := storage.NewSqliteRepository(conn)
		err = repo.Migrate(config.DatabasePath())

		if err != nil {
			fmt.Println("database:", err)
		}

		storageService := storage.NewService(repo)
		notificationService := notifications.New()

		go watcher.New(storageService, notificationService).Start()

		app := server.New(port, storageService)
		app.Setup()
		app.Start()

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntP("port", "p", 6123, "port on which sifaka server will run. Default port is 6123")
}
