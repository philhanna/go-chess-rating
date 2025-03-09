package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Main defines the interface that all rating platform implementations must satisfy.
type Main interface {
	GetURL() string
	ParseContent(content string) []string
	Run()
}

// BaseMain provides common functionality for fetching and processing rating information.
type BaseMain struct {
	Player string
}

// NewBaseMain creates a new instance of BaseMain.
func NewBaseMain(player string) *BaseMain {
	return &BaseMain{Player: player}
}

// Run executes the rating retrieval process.
func (b *BaseMain) Run(m Main) {
	url := m.GetURL()
	content, err := b.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching data: %v\n", err)
		return
	}

	output := m.ParseContent(content)
	if len(output) == 0 {
		fmt.Printf("No ratings found for \"%s\"\n", b.Player)
		return
	}

	for _, line := range output {
		fmt.Println(line)
	}
}

// Get sends an HTTP GET request to the specified URL and returns the response body as a string.
func (b *BaseMain) Get(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Example subclass for a chess rating platform.
type ExampleChess struct {
	*BaseMain
}

// NewExampleChess creates a new instance of ExampleChess.
func NewExampleChess(player string) *ExampleChess {
	return &ExampleChess{BaseMain: NewBaseMain(player)}
}

// GetURL returns the URL for fetching the rating data.
func (e *ExampleChess) GetURL() string {
	return fmt.Sprintf("https://example.com/chess/rating/%s", e.Player)
}

// ParseContent processes the retrieved content and extracts ratings.
func (e *ExampleChess) ParseContent(content string) []string {
	// Dummy parsing logic
	if content == "" {
		return nil
	}
	return []string{"Example rating: 2000"}
}
