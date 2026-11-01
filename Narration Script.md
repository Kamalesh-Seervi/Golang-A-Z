# Go Fundamentals: Recording Script

Four lessons: Hello World, variables, user input, time. Target 16-18 minutes, epilogue 3-4 minutes (4:00 max).

## Rules while recording

- Keep Safari or VS Code large, with the menu bar, clock and dock visible.
- Move the cursor, pause, point, then act. Pause about 1 second after every click.
- Announce every keyboard shortcut (for example, "Command-S to save").
- Keep each typed chunk under 45 seconds. Between chunks, point at the code or run it.
- Say What, Where and Why for each step. Avoid "this" and "that".
- If something fails, say what happened, make one fix, and confirm the result.

## Before recording

Restore the starting files so there is something to edit:

```sh
cd /Users/kamalesh/Documents/Developer/Golang-A-Z
for d in 01_hello 02_variables 03_user_input 04_time; do rm -rf "$d"; cp -R "Deliverables/input 2/$d" "$d"; done
```

Open `go.dev/doc` in Safari and the Google Docs Goal document in a second tab. Open the workspace in VS Code with the terminal panel visible. Do not show the metadata or any PII.

## 1. Intro (Safari, about 1 min)

**Show:** the Google Docs Goal document for a few seconds, then the Go docs at go.dev/doc.

> "This is the Goal document for today's lesson. I'm now moving to Safari, to the official Go documentation at go.dev/doc. I'll use it as my beginner-friendly reference. Go is a compiled, statically typed language from Google, known for simple syntax and fast builds. Today I'll refine four small programs in Visual Studio Code: Hello World, variables, user input and time. I'll fix some misleading comments, save each change, and run every lesson."

## 2. Lesson 01, Hello World (about 2 min)

**Show:** `01_hello/hello.go`.

> "I'm switching to Visual Studio Code and opening `01_hello/hello.go` in the Explorer on the left. The program prints Hello, World. At the bottom, a comment says typing `fmt.Println` makes Go import `fmt` automatically. That's wrong. The Go compiler never adds imports. The Go extension in VS Code does it through its tooling when the file is saved."

**Do:** replace the last comment with:

```go
// The import line above can be added for you by editor tooling (the Go extension
// in VS Code runs goimports/gopls on save). The Go compiler never adds imports itself.
```

Add above `fmt.Println`:

```go
	// Println writes the text followed by a newline to standard output.
```

> "I'm replacing that comment with one that credits the editor tooling, and adding a comment above `Println` explaining it writes text followed by a newline. I'm pressing Command-S to save. Now I'll open the terminal panel, type `go run .` and press Enter. The output is Hello, World!, which confirms the file compiles."

## 3. Lesson 02, variables (about 4 min)

**Show:** `02_variables/main.go`.

> "Now I'm opening `02_variables/main.go`. Near the top, `const token` is commented as public. In Go, visibility depends on the first letter. A lowercase name is unexported, so only this package can use it. I'm changing the comment to say unexported."

**Do:** change the comment on `token`:

```go
// Package-level constant. The lowercase first letter makes it unexported:
// it is visible only inside this package.
const token int = 134325
```

> "To show the other case, I'm typing a second constant named `Version`, with an uppercase V, and a comment saying it's exported."

**Do:** add below it:

```go
// Package-level constant. The uppercase first letter makes it exported:
// other packages can use it as variables.Version.
const Version string = "1.0"
```

Then add this as the last line of `main`:

```go
	fmt.Println("Version:", Version)
```

> "In `main` I'll add a line that prints `Version`. I'm also fixing the float32 comment. A float32 holds about seven significant digits, which is why 4592.028002853085 prints as 4592.028."

**Do:** replace the float comment with:

```go
	// float32 holds about 7 significant digits, so the extra digits below are rounded off
```

> "I'm saving with Command-S, running `go run .` in the terminal and tracing the output: string, int, bool, float32, the zero values, the inferred type, the walrus variable, and finally the two constants. Version prints 1.0."

## 4. Lesson 03, user input (about 4 min)

**Show:** `03_user_input/main.go`.

> "Next is `03_user_input/main.go`. This program reads a rating with `bufio` and `ReadString`. The old comment says the value is an integer, but `ReadString` always returns a string, and the value still has its trailing newline. I'll correct that comment."

**Do, chunk 1:** update imports:

```go
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
```

> "I'm editing the import block to add `strconv` and `strings`."

**Do, chunk 2:** replace the read section:

```go
	// ReadString returns (string, error). The "comma ok / err" idiom lets us check the error.
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input:", err)
		return
	}

	// ReadString always returns a string, even when the user types digits.
	// The value keeps the trailing newline, so trim it before converting.
	input = strings.TrimSpace(input)
	fmt.Printf("type of input: %T\n", input)
```

> "I'm changing `input, _` to `input, err` and adding an `if err != nil` check, since ignoring errors hides problems. Then `strings.TrimSpace` removes the newline."

**Do, chunk 3:** convert and print:

```go
	// Convert the string to a number with strconv.ParseFloat.
	rating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Rating must be a number:", err)
		return
	}

	fmt.Println("Thanks for rating,", rating)
	fmt.Printf("type of rating: %T\n", rating)
	fmt.Println("Rating plus 1:", rating+1)
```

> "`strconv.ParseFloat` converts the string to a number, again with an error check. Last, I'll print the rating and `rating+1` to prove it's a number."

**Run:**

> "I'm saving with Command-S and running `go run .`. When prompted, I'll type 4.5. The output shows type of input is string, rating is float64, and rating plus 1 is 5.5. Now I'll run it again and type abc to check the failure path. It prints a clear error and exits."

## 5. Lesson 04, time (about 3 min)

**Show:** `04_time/go.mod`, then `04_time/main.go`.

> "I'm opening `04_time/go.mod`. The module is named `time`, which collides with Go's standard `time` package and causes an 'ambiguous import' error. I'm renaming it to `timelesson`."

**Do:** change `module time` to `module timelesson`. Optionally run `go run .` first to show the error, then fix it.

**Do:** in `main.go` add:

```go
	// Go formats dates using the reference time Mon Jan 2 15:04:05 2006 (01-02-2006 = month-day-year).
```

above the `Format` line, and after printing `createdDate`:

```go
	fmt.Println(createdDate.Format("02 Jan 2006"))

	// Since returns a time.Duration measured from the given time until now.
	fmt.Println("Hours since created:", int(time.Since(createdDate).Hours()))
```

> "In `main.go` I'm adding a comment explaining that Go formats dates using the reference time, Monday January 2, 2006 at 15:04:05. I'm also adding a line that formats `createdDate`, and one that uses `time.Since` to print the hours since that date. I'm saving with Command-S and running `go run .`. The output shows the current time, the formatted time, the created date and the hours elapsed."

## 6. Final check and epilogue (3-4 min, 4:00 max)

**Show:** the VS Code Explorer with the four lesson folders, all saved (no unsaved dots on the tabs).

> "All four lesson files are saved. In the Explorer on the left, `01_hello`, `02_variables`, `03_user_input` and `04_time` each contain their `go.mod` and Go file. These refined lessons are my final output. After this recording I'll copy them into a `final output 1` folder with the Goal document and the original lessons, and zip everything."

**Epilogue:**

> "I verified three things. First, all four lessons compile and run after my edits. Second, the rating input works: 4.5 gives a float64 and 5.5, and abc gives an error. Third, the `time` module name now builds. The limitations are that the program only checks for a number, not a 1 to 5 range. The time output depends on my local clock and timezone. I tested on macOS only.
>
> My main advice for beginners is to read the compiler and run output, not just the comments. Comments can be wrong. Two of mine were. Always check what a function really returns, and check errors instead of discarding them with an underscore. Also avoid naming your module after a standard library package."

## After recording

Packaging is not part of the recording.

1. Replace the contents of `Deliverables/final output 1` with your final on-camera lesson folders (`01_hello` to `04_time`, each with `go.mod`).
2. Check that `Deliverables` has `input 1.docx`, `input 2/` and `final output 1/`.
3. Zip it: `zip -r Deliverables.zip Deliverables -x "*.DS_Store"`.
4. Run `unzip -l Deliverables.zip` and confirm every file is there.
5. Keep this script, the metadata RTF, `video.mp4` and the TaskRecorder archive outside the zip.
