package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// PageData represents the template rendering data context
type PageData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

func main() {
	// Hybrid mode: Run in CLI mode if arguments are provided (keeps tests passing)
	if len(os.Args) == 2 || len(os.Args) == 3 {
		runCLI()
		return
	}

	if len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	startServer()
}

func runCLI() {
	input := os.Args[1]
	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" {
		fmt.Println("Error: Invalid banner style. Choose standard, shadow, or thinkertoy")
		return
	}

	bannerBytes, err := os.ReadFile(bannerName + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file:", err)
		return
	}

	result, err := GenerateASCIIArt(input, string(bannerBytes))
	if err != nil {
		fmt.Println("Error generating ASCII art:", err)
		return
	}

	fmt.Print(result)
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleASCIIArt)
	http.HandleFunc("/ascii-art-web", handleASCIIArt)

	fmt.Printf("🚀 ASCII Command Station running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

// render helper parses and renders templates cleanly, reporting any missing files
func render(w http.ResponseWriter, status int, data PageData) {
	tmplPath := filepath.Join("templates", "index.html")
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		http.Error(w, "404 Not Found: Template file missing", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "500 Internal Server Error: Failed to parse template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error: Failed to render template", http.StatusInternalServerError)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found: Path not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request: GET method required", http.StatusBadRequest)
		return
	}

	render(w, http.StatusOK, PageData{Banner: "standard"})
}

func handleASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request: POST method required", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "400 Bad Request: Invalid banner specified", http.StatusBadRequest)
		return
	}

	// Validate ASCII printable characters
	for _, r := range text {
		if (r < 32 || r > 126) && r != '\n' && r != '\r' {
			http.Error(w, "400 Bad Request: Non-ASCII characters are not allowed", http.StatusBadRequest)
			return
		}
	}

	bannerBytes, err := os.ReadFile(banner + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404 Not Found: Banner file missing", http.StatusNotFound)
			return
		}
		http.Error(w, "500 Internal Server Error: Failed to read banner", http.StatusInternalServerError)
		return
	}

	result, err := GenerateASCIIArt(text, string(bannerBytes))
	if err != nil {
		http.Error(w, "500 Internal Server Error: Art generation failed", http.StatusInternalServerError)
		return
	}

	render(w, http.StatusOK, PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	})
}

// GenerateASCIIArt converts input string to graphic representation using banner contents
func GenerateASCIIArt(input string, bannerContent string) (string, error) {
	cleaned := strings.ReplaceAll(bannerContent, "\r\n", "\n")
	splitted := strings.Split(cleaned, "\n")

	if len(splitted) < 855 {
		return "", fmt.Errorf("invalid banner file structure")
	}

	asciiMap := make(map[rune][]string)
	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9
		if start+9 > len(splitted) {
			return "", fmt.Errorf("banner file is too short")
		}
		asciiMap[rune(i)] = splitted[start+1 : start+9]
	}

	// Normalize input newlines
	normalizedInput := strings.ReplaceAll(input, "\r\n", "\n")
	normalizedInput = strings.ReplaceAll(normalizedInput, "\\n", "\n")

	if normalizedInput == "" {
		return "\n", nil
	}

	lines := strings.Split(normalizedInput, "\n")

	// If the entire input is just newlines, repeat newlines
	onlyNewlines := true
	for _, line := range lines {
		if line != "" {
			onlyNewlines = false
			break
		}
	}

	if onlyNewlines {
		return strings.Repeat("\n", len(lines)-1), nil
	}

	var sb strings.Builder
	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				sb.WriteString(asciiMap[ch][row])
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}