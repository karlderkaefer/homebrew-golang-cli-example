package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type GenderizeResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func printVersion() string {
	version := fmt.Sprintf("genderize %s, commit %s, built at %s", version, commit, date)
	fmt.Println(version)
	return version
}

func printHelp() {
	fmt.Println("Usage: go run main.go [name]")
	fmt.Println("Example: go run main.go scott")
}

func autoUpgrade() {

	upgradeCmd := exec.Command("sh", "-c", `HOMEBREW_GITHUB_API_TOKEN=$(gh auth token) brew upgrade karlderkaefer/homebrew-tap/genderize`)
	upgradeCmd.Stdout = os.Stdout
	upgradeCmd.Stderr = os.Stderr
	if err := upgradeCmd.Run(); err != nil {
		fmt.Println("Error: Unable to upgrade the tool:", err)
		os.Exit(1)
	}

	fmt.Println("Upgrade completed successfully.")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		printVersion()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "upgrade" {
		autoUpgrade()
		return
	}

	if len(os.Args) < 2 {
		printHelp()
		os.Args = append(os.Args, "scott")
	}

	name := os.Args[1]
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	fmt.Println("Checking URL", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: Unable to make the request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d from the API. This may indicate that the request was blocked or there is an issue with the API.\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response Body:", string(body))
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Unable to read the response body:", err)
		os.Exit(1)
	}

	var genderizeResponse GenderizeResponse
	err = json.Unmarshal(body, &genderizeResponse)
	if err != nil {
		fmt.Println("Error: Unable to parse the response body as JSON:", err)
		os.Exit(1)
	}

	fmt.Printf("The predicted gender for the name '%s' is '%s' with a probability of %.2f\n", genderizeResponse.Name, genderizeResponse.Gender, genderizeResponse.Probability)
}
