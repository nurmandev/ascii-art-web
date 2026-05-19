# ASCII Art Web

A simple and interactive Go application that converts ordinary text into cool ASCII art. It features a modern web interface and a command-line interface (CLI), supporting three different font styles.

---

## Features

*   **Web Interface**: A clean, modern page where you can type your text, select a style, and see the generated ASCII art instantly.
*   **Command-Line (CLI)**: A quick way to generate ASCII art directly in your terminal.
*   **3 Font Styles**:
    *   `standard` - A classic and bold design.
    *   `shadow` - A cool 3D shadow block design.
    *   `thinkertoy` - A retro, balloon-like connected design.

---

## How to Run

The app is built entirely with Go's standard library, so there are no external dependencies to install.

### 1. Web Server

To launch the web interface, run this command in the project folder:

```bash
go run .
```

By default, the server starts on port **8080**. Open your browser and go to:
👉 [http://localhost:8080](http://localhost:8080)

#### Custom Port (Optional)
To use a different port, set the `PORT` environment variable:
*   **Windows (PowerShell)**: `$env:PORT="9000"; go run .`
*   **Linux / macOS**: `PORT=9000 go run .`

---

### 2. Command Line (CLI)

To output ASCII art directly to the terminal:

```bash
go run . [TEXT] [FONT]
```

**Examples:**
```bash
go run . "hello" standard
go run . "hello" shadow
go run . "hello" thinkertoy
```

*Note: If you don't specify a font, the app defaults to `standard`.*

---

### 3. Running Tests

To run the automated tests:

```bash
go test -v
```

---

## How It Works

The program converts text into graphic ASCII letters (each 8 lines tall) using the following simple steps:

1.  **Reads the Font File**: It opens the selected font template (`standard.txt`, `shadow.txt`, or `thinkertoy.txt`).
2.  **Locates Characters**: Each character in the template is 8 lines tall. The program calculates the position of each character based on its ASCII number value.
3.  **Draws the Art**: It loops line-by-line (0 to 7) for the input text, joining the corresponding lines of each character side-by-side.
4.  **Delivers the Output**: Prints the reconstructed lines directly to your terminal or displays them on the webpage.
