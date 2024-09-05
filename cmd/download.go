/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file from a server",
	Long: `Download a file from a server using a GET request.

Example:
	download --url http://localhost:8080/file.txt --path /path/to/save/file.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
	},
}

// function to download a file from a server using a GET request and save it to the local filesystem
func download(url string, file string) {
	// Create a new file to save the downloaded file
	out, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer out.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error copying file: ", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add a flag for the file to download
	downloadCmd.Flags().StringP("path", "f", "", "The path to the file to download")
	downloadCmd.Flags().StringP("url", "u", "", "The URL to download the file from")
	downloadCmd.MarkFlagRequired("path")
	downloadCmd.MarkFlagRequired("url")

	// Set the run function for the download command
	downloadCmd.Run = func(cmd *cobra.Command, args []string) {
		// Get the flags
		url, _ := cmd.Flags().GetString("url")
		file, _ := cmd.Flags().GetString("path")

		// Download the file
		download(url, file)
	}
}
