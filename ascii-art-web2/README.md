ASCII-art-Web
 ASCII-Art Web is a project written in Go that generates ASCII-art images from plain text received as input on the web.
It uses assets for ASCII fonts and Go templates to provide a clean, interactive web interface, allowing users to enter text and receive a graphical ASCII representation in real time.

## The file structure is shown below:-
  ```plaintext
ascii-art-web/
│
├── cmd/                  # Application entry point
│    └── main.go          # Starts the server
├── pkg/                  # Core packages
├── internal/             # Internal modules
│    ├── config.go        # Configuration handling
│    ├── handler.go       # HTTP request handlers
│    └── processor.go     # ASCII processing logic
├── web/                  # Web interface files
│    ├── index.html       # Input form
│    └── response.html    # ASCII-art output page
├── assets/               # ASCII font files
│    ├── standard.txt
│    ├── shadow.txt
│    └── thinkertoy.txt
├── file/                 # Optional storage or uploaded files
└── README.md             # Project documentation
  ```

### Features
1. Real time ASCII-art generation from the user input.
2. Uses POST requests ti send text from the web form to the server.
3. Handles HTTP requests and generates ASCII-art responses dynamically.
4. Supports multiple fonts stored in assats.
5. Clean web intrface(index.html for input, response.html for output)
6. Error handling for empty or unsupported input.

#### Usage
- Open the index page in the browser.
- Enter the text you want to convert.
- Select a banner from the available options
- Click Generate ASCII Art(send a POST request to the server).
- View the generated ASCII-art on the response page.
- click Home to go back to home page

##### Technology and version 
- go version go1.22.2 linux/amd64
- net/http package for POST and HTTP request handling.
- html/template package
- HTML5 and CSS

##### LINK
`http://localhost:8081`

##### Error Handling
- Empty input - field cannot be empty message pops up.
- Unsupported Characters - skipped or replaced
- Server errors - return HTTP 500 with explanation.


###### Author
1. Divine
2. Ola
3. John 