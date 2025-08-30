
package cli 
import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

type GithubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}


func getGithubActivityOfUser(username string){
	if len(username) == 0 {
		fmt.Println("Invalid username, Kindly provide a valid username")
		return
	}

	url := "https://api.github.com/users/"+username+"/events"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error fetching data from GitHub API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch data: %s\n", resp.Status)
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var events []GithubEvent
	if err := json.Unmarshal(body, &events); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No recent public activity for this user.")
		return
	}

	for _, e := range events {
		switch e.Type {
		case "PushEvent":
			fmt.Printf("Pushed to %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "PullRequestEvent":
			fmt.Printf("Pull request in %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "IssuesEvent":
			fmt.Printf("Opened Issue in %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "CreateEvent":
			fmt.Printf("Created in %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "PullRequestReviewEvent":
			fmt.Printf("Reviewed pull request in %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "ForkEvent":
			fmt.Printf("Forked %s at %s\n", e.Repo.Name, e.CreatedAt)
		case "WatchEvent":
			fmt.Printf("Starred %s at %s\n", e.Repo.Name, e.CreatedAt)
		default:
			fmt.Printf("Event: %s Repo: %s Date: %s\n", e.Type, e.Repo.Name, e.CreatedAt)
		}
	}
}