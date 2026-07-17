# ascii-art

A command-line program written in Go that takes a string as input and outputs it as large ASCII art characters using pre-made banner files.

---

## How It Works

Each character in the input is looked up in a banner file where every printable ASCII character is pre-drawn across exactly 8 lines. The program prints all characters side by side, row by row, to produce the final ASCII art output.

**Example:**
```bash
go run main.go "Hello"
```
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

---

## Usage
```bash
go run main.go "<your text>"
```

### Examples
```bash
go run main.go "Hello"
go run main.go "Hello There"
go run main.go "Hello\nThere"
go run main.go "Hello\n\nThere"
go run main.go "123"
go run main.go "{Hello There}"
go run main.go ""
```

---

## Supported Input

| Input type         | Example            | Behaviour                             |
|--------------------|--------------------|---------------------------------------|
| Uppercase letters  | `"HELLO"`          | Drawn in large ASCII art              |
| Lowercase letters  | `"hello"`          | Drawn in large ASCII art              |
| Numbers            | `"123"`            | Drawn in large ASCII art              |
| Spaces             | `"Hello There"`    | Space character has an 8-line drawing |
| Special characters | `"{!@#}"`          | All printable ASCII characters work   |
| `\n`               | `"Hello\nThere"`   | Starts a new block of ASCII art       |
| `\n\n`             | `"Hello\n\nThere"` | Empty line between two blocks         |
| Empty string       | `""`               | Prints nothing                        |

---

## Banner Files

The project includes three banner styles. The program uses **standard** by default.

| File              | Style description               |
|-------------------|---------------------------------|
| `standard.txt`    | Clean block letters             |
| `shadow.txt`      | Characters with a shadow style  |
| `thinkertoy.txt`  | Rounded playful style           |

Each banner file contains every printable ASCII character from space to `~`, drawn across exactly 8 lines and separated by blank lines. All three files have 855 lines total.

---

## Project Structure
```
ascii-art/
├── go.mod
├── main.go          ← core program logic
├── main_test.go     ← unit tests
├── standard.txt     ← standard banner file
├── shadow.txt       ← shadow banner file
└── thinkertoy.txt   ← thinkertoy banner file
```

---

## How the Character Lookup Works

Every character in the banner file has a fixed position calculated by this formula:
```
startIndex = (asciiValue - 32) * 9 + 1
```

- `asciiValue - 32` — space is ASCII 32 and is the first character in the file, so we offset from it
- `* 9` — each character occupies 9 lines (8 drawing lines + 1 blank separator)
- `+ 1` — Go slices are zero-indexed so we adjust by 1

The program reads exactly 8 lines from that position to get the full character drawing.

---

## Running Tests
```bash
go test ./...
```

Expected output when all tests pass:
```
ok      ascii-art       0.002s
```

---

## Requirements

- Go 1.16 or higher
- Standard Go packages only — no external dependencies

---

## Packages Used

- `fmt` — printing output
- `os` — reading files and command-line arguments
- `strings` — splitting strings and building output

---

## Authors

-Divine Ameh (diameh)
-Victor Ameh (vameh)
-Oluwatofunmi Osifowora (oosifowo)
