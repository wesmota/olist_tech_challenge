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

	"github.com/spf13/cobra"
	"github.com/wesmota/olist_tech_challenge/rest-api/csv-reader/reader"
)

var csvFilePath string
var sqliteBDPath string

// csvreaderCmd represents the csvreader command
var csvreaderCmd = &cobra.Command{
	Use:   "csvreader",
	Short: "Command to read a csv file and load into sqlite database.",
	Long: `csvReader is a command to read a csv file containg authors names and 
	import it to a sqlite3 database.
	Usage:
	 - csvreader -f authors.csv -s olist.db`,
	Run: func(cmd *cobra.Command, args []string) {
		csvLoader := reader.NewCSVLoader(csvFilePath, sqliteBDPath)
		csvLoader.Load()
		fmt.Println("csvreader called")
	},
}

func init() {
	rootCmd.AddCommand(csvreaderCmd)
	csvreaderCmd.Flags().StringVarP(&csvFilePath, "csvFilePath", "f", "", "csv file fullpath to be imported")
	csvreaderCmd.Flags().StringVarP(&sqliteBDPath, "sqliteBDPath", "s", "", "sqlite db fullpath to be imported")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csvreaderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvreaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
