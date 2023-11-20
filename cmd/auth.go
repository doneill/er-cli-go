package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var SITENAME string

// ----------------------------------------------
// auth command
// ----------------------------------------------

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication with EarthRanger",
	Long:  `Authenticate er with EarthRanger`,
	Run: func(cmd *cobra.Command, args []string) {
		auth()
	},
}

// ----------------------------------------------
// funtions
// ----------------------------------------------

func auth() {
	fmt.Println("Enter username:")
	var username string
	_, err := fmt.Scan(&username)
	// _, err = fmt.Scan(&username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Enter password:")
	var password string
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call the authenticate function to get the access token and expires in
	response, err := authenticate(SITENAME, username, password)
	if err != nil {
		fmt.Println("Error authenticating:", err)
		os.Exit(1)
	}

	// Print out the access token and expires in if the request was successful
	if response != nil {
		fmt.Printf("Access Token: %s\n", response.AccessToken)
		fmt.Printf("Expires In: %d\n", response.ExpiresIn)
	}
}

// ----------------------------------------------
// initialize
// ----------------------------------------------

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringVarP(&SITENAME, "sitename", "s", "", "EarthRanger site name")
	authCmd.MarkFlagRequired("sitename")
}
