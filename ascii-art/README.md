# ASCII Art

## AIM

This program allows the user to input text in the command line and receive back an output displaying that text using ASCII characters from a standard template file.

---

## How It Works

1. The user passes a string as a command-line argument
2. The program reads a pre-built ASCII art template file (`templates/standard.txt`)
3. Each character in the template is 8 lines tall and covers every printable ASCII character from space (` `) to tilde (`~`)
4. The program maps each letter of the user's input to its matching ASCII art block and prints them side by side, row by row

### Newlines

To print on multiple lines, use the literal `\n` in your input:

```
go run . "Hello\nWorld"
```

This will print `Hello` and `World` on separate ASCII art lines.

---

## Usage

```bash
go run . "YourTextHere"
```

**Example:**

```bash
go run . "Hi"
```

---

## Function Breakdown

| Function | Description |
|---|---|
| `getWord()` | Validates and retrieves the user's input from `os.Args` |
| `readCharacters()` | Reads `standard.txt`, splits it into lines, and groups every 8 lines into one ASCII character |
| `printWord()` | Splits input on literal `\n`, handles blank lines, delegates rendering to `print()` |
| `print()` | Iterates row by row across all characters in a word, printing each row side by side |

---

## Features

- Written in **Go**
- Follows Go good practices and conventions
- Supports all printable ASCII characters (space through `~`)
- Handles multi-line input via literal `\n`
- Includes unit test files for core functions

---

## What We Learned

- Reading files with `os.ReadFile`
- Data manipulation — splitting, slicing, and indexing strings and byte slices
- How `os.Args` works for command-line input
- Working with `rune` types for character mapping
- Structuring a Go project across multiple files in the same package

---

## Requirements

- Go 1.26.1
- `templates/standard.txt` must be present in the project root

---

## Authors
    - Daniel Uduma
    - Chukwuebuka Okoh
    - David Oloruntoba