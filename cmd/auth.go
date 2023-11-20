package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication with EarthRanger",
	Long:  `Authenticate er with EarthRanger`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(authCmd)

	fmt.Println("Enter sitename:")
	var sitename string
	_, err := fmt.Scan(&sitename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Enter username:")
	var username string
	_, err = fmt.Scan(&username)
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
	response, err := authenticate(sitename, username, password)
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
