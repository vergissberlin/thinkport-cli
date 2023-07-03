/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const apiURL = "https://api.thinkport.andrelademann.de"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thinkport",
	Short: "The thinkport command line interface",
	Long: `
Get latest informations about the company.
You can get also more information about the company on the website https://thinkport.digital.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.thinkport.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJSON gets a JSON from a REST API and unmarshals it into a struct
func GetJSON(url string, target interface{}) error {
	// Get JSON from REST API
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	// Read JSON
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Unmarshal JSON into struct
	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}
	return nil
}
