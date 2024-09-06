/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"mime/multipart"
	"bytes"
	"io"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to a server",
	Long: `Upload a file to a server using a POST request.
This command handles multipart form data and can be used to upload files to a server.

Example:
	upload --url http://localhost:8080/upload --path /path/to/file.txt`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}

// upload function that takes a url and a file path and uploads the file to the server using a POST request
func upload(url string, file string) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer f.Close()

	// Create a new buffer to store the file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form file
	part, err := writer.CreateFormFile("file", filepath.Base(file))
	if err != nil {
		fmt.Println("Error creating form file: ", err)
		return
	}

	// Copy the file to the form file
	_, err = io.Copy(part, f)
	if err != nil {
		fmt.Println("Error copying file: ", err)
		return
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer: ", err)
		return
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	// Set the content type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: ", resp.Status)
		return
	}

	fmt.Println("File uploaded successfully")
}



func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	uploadCmd.Flags().StringP("url", "u", "", "The URL to upload the file to")
	uploadCmd.Flags().StringP("path", "p", "", "The file to upload")
	uploadCmd.MarkFlagRequired("url")
	uploadCmd.MarkFlagRequired("path")

	uploadCmd.Run = func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		file, _ := cmd.Flags().GetString("path")
		upload(url, file)
	}
}
