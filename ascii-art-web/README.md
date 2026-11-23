# ASCII Art Web
A web-based Graphical User Interface (GUI) for the ASCII Art generator, built entirely in Go using only standard packages.

## Description
This project takes a string of text and a banner style as inputs via a web interface, processes it on a Go server, and displays the graphical ASCII art representation back to the user.

## Implementation Details
- **Routing:** Handled via the standard `net/http` package.

- **Templating:** Uses `html/template` to serve the main GUI and dynamically render ASCII output.

- **Algorithm:** Splitting the input text by newlines and mapping each character's ASCII value to its corresponding 8-line block within the chosen banner file (`standard`, `shadow`, or `thinkertoy`).

- **Error Handling:** Returns explicit HTTP status codes (200 OK, 400 Bad Request, 404 Not Found, 500 Internal Server Error) based on request validity and server-side file availability.

## Usage
1. Ensure you have Go installed on your system.

2. Place your banner files in the expected assets directory.

3. Run the server from the root directory:
    ```bash
    go run main.go
    ```

4. Open your web browser and navigate to `http://localhost:8080`.

5. Enter your text, select a banner style, and click submit to view the generated ASCII art.

## Authors
* Grace Kalu
* Daniel Uduma
