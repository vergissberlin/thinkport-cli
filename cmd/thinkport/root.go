/*
Copyright © 2023 André Lademann <alademann@thinkport.digital>
*/
package thinkport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/savioxavier/termlink"
	"github.com/spf13/cobra"
)

const apiURL = "https://api.thinkport.andrelademann.de"

var version = "0.0.53"

// Connection pool
var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
	},
}

// API response cache
var cache = map[string]interface{}{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "thinkport",
	Short:   "The thinkport command line interface",
	Version: version,
	Long: ` _______ _     _       _                     _   
|__   __| |   (_)     | |                   | |  
   | |  | |__  _ _ __ | | ___ __   ___  _ __| |_ 
   | |  | '_ \| | '_ \| |/ / '_ \ / _ \| '__| __|
   | |  | | | | | | | |   <| |_) | (_) | |  | |_ 
   |_|  |_| |_|_|_| |_|_|\_\ .__/ \___/|_|   \__|
                           | |                   
                           |_|                   

Get latest informations about the company.
You can get also more information about the
company on our website ` + termlink.ColorLink("https://thinkport.digital", "https://thinkport.digital", "blue") + `.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
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
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetJSON gets a JSON from a REST API and unmarshals it into a struct
func GetJSON(url string, target interface{}) error {

	// Check cache
	if val, ok := cache[url]; ok {
		target = val
		return nil
	}

	// Get from API
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	// Read and unmarshal
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	// Add to cache and return
	cache[url] = target
	return nil
}

// Get version number of latest release from GitHub API
func GetLatestVersion() string {

	// Check cache
	if val, ok := cache["version"]; ok {
		return val.(string)
	}

	// Get from API
	res, err := client.Get("https://api.github.com/repos/vergissberlin/thinkport/releases/latest")
	if err != nil {
		return version
	}

	// Read and unmarshal
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return version
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return version
	}

	// Check if version is set
	if data["tag_name"] == nil {
		return version
	}

	// Add to cache and return
	cache["version"] = data["tag_name"]
	return data["tag_name"].(string)
}
