package commands

import (
	"fmt"
	"os"
	"time"

	AI "go-browser/AI"
	IO "go-browser/IO"
	searchbrowser "go-browser/search-browser"
	"go-browser/site"
	utils "go-browser/utils"
)

// HandleHelp displays the help page based on the user's input.
func HandleHelp() {
	page := utils.UserWriteNum("Enter page number for help (e.g., 1, 2, 3):")
	utils.DisplayHelp(page)
}

// HandleExit prints a message and exits the program.
func HandleExit() {
	fmt.Println("Exiting program.")
	os.Exit(1)
}

// HandleCreate prompts the user for a file name and content, then creates the file.
func HandleCreate() {
	name := utils.UserWriteString("Enter file name:")
	text := utils.UserWriteString("Enter file content:")
	path := utils.UserWriteString("Enter path: (press enter for main folder)")
	IO.CreateFile(name, text, path)
}

// HandleRead prompts the user for a file name and reads the content of the file.
func HandleRead() {
	name := utils.UserWriteString("Enter file name for reading:")
	IO.ReadFile(name)
}

// HandleDelete prompts the user for a file name and deletes the file.
func HandleDelete() {
	name := utils.UserWriteString("Enter file name for deletion:")
	IO.DeleteFile(name)
}

// HandleUpdate prompts the user for a file name and new content, then updates the file.
func HandleUpdate() {
	name := utils.UserWriteString("Enter file name for update:")
	text := utils.UserWriteString("Enter new content:")
	IO.UpdateFile(name, text)
}

// HandleRename prompts the user for a file name and a new name, then renames the file.
func HandleRename() {
	name := utils.UserWriteString("Enter file name to rename:")
	newName := utils.UserWriteString("Enter new file name:")
	IO.RenameFile(name, newName)
}

// HandleAbout displays information about the program version.
func HandleAbout() {
	fmt.Println("Version 0.0.0 - Go Browser Tool")
}

// HandleList lists all files in the current directory.
func HandleList() {
	path := utils.UserWriteString("Enter Path: (. for main folder)")
	IO.ListFile(path)
}

// HandleAIChat prompts the user for text and sends it to the AI model for a response.
func HandleAIChat() {
	text := utils.UserWriteString("Enter text:")
	response := AI.ChatGPT(text)
	fmt.Println(response)
}

// HandleGoogle prompts the user for search text and performs a Google search.
func HandleGoogle() {
	search := utils.UserWriteString("Enter text for search:")

	result, err := searchbrowser.SearchGoogle(search)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the search results
	for _, res := range result {
		fmt.Println(res)
	}
}

// HandleSitePerformance prompts the user for a URL and performs a performance test on the site.
func HandleSitePerformance() {
	url := utils.UserWriteString("Enter site URL to test performance:")
	live := utils.UserWriteBool("Enable live monitoring? (true/false):")
	timeout := 10 * time.Second
	err := site.MeasureSitePerformance(url, timeout, live)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// HandleSiteContent prompts the user for input to fetch and display content from a specified site.
func HandleSiteContent() {
	url := utils.UserWriteString("Enter site URL:")
	element := utils.UserWriteString("Specify the target element (or leave empty for all):")
	includeAttributes := utils.UserWriteBool("(true/false) Include attributes in HTML elements?")
	filter := utils.UserWriteBool("(true/false) Filter unnecessary elements like scripts?")
	langBool := utils.UserWriteBool("(true/false) Default language en-US")
	lang := "en-US"
	if !langBool {
		lang = utils.UserWriteString("Language: (en-US, es-ES, zh-CN..)")
	}

	// Create an instance of SiteOptions with user input
	options := site.SiteOptions{
		URL:               url,
		Element:           element,
		Language:          lang,
		IncludeAttributes: includeAttributes,
		Filter:            filter,
	}

	// Call SiteContent with the SiteOptions instance
	content, err := site.SiteContent(options)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the extracted content
	fmt.Println("Extracted Content:")
	fmt.Println(content)

	// Ask the user if they want to save the content
	save := utils.UserWriteBool("(true/false) Save content?")
	if save {
		name := utils.UserWriteString("Enter file name to save content: (.text, .html...)")
		path := utils.UserWriteString("Enter path: (press enter for main folder)")
		if err := IO.CreateFile(name, content, path); err != nil {
			fmt.Printf("Failed to save content to file: %v\n", err)
		} else {
			fmt.Println("Content saved successfully.")
		}
	}
}

func HandleHttpRequest() {
	// Get the URL input from the user
	url := utils.UserWriteString("Enter URL for Fetch:")

	// Example: https://jsonplaceholder.typicode.com/posts
	// POST
	// {"id":101,"title":"foo","body":"bar","userId":1}
	method := utils.UserWriteString("Enter method:")

	// Get JSON body input if needed (for POST method)
	bodyContent := utils.UserWriteJson("Enter Body:")

	// If the method is POST, format the body as JSON
	var body []byte
	if method == "POST" && bodyContent != "" {
		body = []byte(bodyContent)
	} else if method == "POST" && bodyContent == "" {
		// Body cannot be empty for POST method
		fmt.Println("For POST method, the body cannot be empty.")
		return
	} else {
		// For GET method, the body is empty
		body = nil
	}

	// Call the HTTP request function with the given parameters
	response, statusCode, err := site.HttpRequest(url, method, body)
	if err != nil {
		// Handle errors if the HTTP request fails
		fmt.Println("Error:", err)
	} else {
		// Print the response and status code
		fmt.Printf("Response: %s\nStatus Code: %d\n", response, statusCode)
	}
}
