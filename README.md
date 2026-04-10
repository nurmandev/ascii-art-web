# ASCII Art Web

## Description
Ascii-art-web is a web Graphical User Interface for the ascii-art tool. It runs a web server built in Go that presents a user-friendly interface to convert text into stylized ASCII art using different banners (Standard, Shadow, and Thinkertoy). The underlying API processes requests and intelligently constructs the text representations.

## Authors
- Advanced Agent AI (Antigravity)

## Usage: How to run
1. Check that you have Go installed on your machine.
2. Clone this repository or navigate to its directory.
3. Make sure that the banner `.txt` files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) are present in the project root directory.
4. Run the following command to start the web server:
   ```bash
   go run .
   ```
5. Open up a web browser and visit `http://localhost:8080`.
6. Enter your text, select a banner style, and click "Generate ASCII Art" to see your result!

## Implementation details: Algorithm
- The program initializes a basic HTTP server using `net/http` package from the Go standard library.
- The `/` route maps to the home handler that renders the `index.html` file using the `html/template` package.
- When the user submits the form, a `POST` request is sent to `/ascii-art`.
- The backend parses the text and standard banner option, checking for incorrect inputs like missing values or out-of-bounds characters. Invalid inputs return `400 Bad Request`.
- If the requested banner does not exist or template is missing, `404 Not Found` is returned.
- Unhandled errors, such as a missing banner `.txt` file during parsing, trigger a `500 Internal Server Error`.
- The textual conversion engine iterates over each character, finds the corresponding visual lines within the desired banner `.txt` file (using its mathematical offset indexing formula), and constructs the completed ASCII art block line by line.

