# Go Fundamentals: Recording Script

Four lessons: Hello World, variables, user input, time. Target 16-18 minutes, epilogue 3-4 minutes (4:00 max).

The comments in the lesson files are already corrected. On camera you point at them and explain them, then type code additions only.

## Rules while recording

- Keep Safari or VS Code large, with the menu bar, clock and dock visible.
- Move the cursor, pause, point, then act. Pause about 1 second after every click.
- Announce every keyboard shortcut (for example, "Command-S to save").
- Keep each typed chunk under 45 seconds. Between chunks, point at the code or run it.
- Say What, Where and Why for each step. Avoid "this" and "that".
- If something fails, say what happened, make one fix, and confirm the result.

## Before recording

Open `go.dev/doc` in Safari and the Google Docs Goal document in a second tab. Open the workspace in VS Code with the terminal panel visible. Do not show the metadata or any PII.

Current starting state of the files:

- `01_hello/hello.go`: complete, comments corrected.
- `02_variables/main.go`: original code, comments corrected. No `Version` constant yet.
- `03_user_input/main.go`: original code, comments corrected. It ignores the error and prints the raw string.
- `04_time/main.go`: prints now and a formatted time. `04_time/go.mod` still says `module time`, which breaks `go run`.

## 1. Intro (Safari, about 1 min)

**Show:** the Google Docs Goal document for a few seconds, then the Go docs at go.dev/doc.

> "This is the Goal document for today's lesson. I'm now moving to Safari, to the official Go documentation at go.dev/doc. I'll use it as my beginner-friendly reference. Go is a compiled, statically typed language from Google, known for simple syntax and fast builds. Today I'll build on four small programs in Visual Studio Code: Hello World, variables, user input and time. I'll explain their comments, add code to each, save every change, and run every lesson."

## 2. Lesson 01, Hello World (about 2 min)

**Show:** `01_hello/hello.go`.

> "I'm switching to Visual Studio Code and opening `01_hello/hello.go` in the Explorer on the left. The program prints Hello, World. The comment above `Println` says it writes the text followed by a newline. The comment at the bottom explains that the import line can be added by editor tooling, which is the Go extension running goimports and gopls on save. The Go compiler itself never adds imports."

**Code addition:** type this on a new line under the first `Println`:

```go
	fmt.Println("Welcome to Go fundamentals")
```

> "I'm adding a second `fmt.Println` inside `main` that prints a welcome line. I'm pressing Command-S to save. Now I'll open the terminal panel, type `go run .` and press Enter. Both lines print, so the program compiles and runs."

## 3. Lesson 02, variables (about 4 min)

**Show:** `02_variables/main.go`.

> "Now I'm opening `02_variables/main.go`. At the top, the comment on `token` says a lowercase first letter makes it unexported, so only this package can use it. The old comment called it public, which was wrong. The comment also says an uppercase name would be exported. Lower down, the float32 comment explains that it holds about seven significant digits, which is why the long number is rounded when printed. The walrus comment says `:=` only works inside a function."

**Code addition, chunk 1:** under the `token` constant, type:

```go
const Version string = "1.0"
```

> "To show the exported case, I'm typing a second constant under `token`. It's named `Version` with an uppercase V, so it is exported."

**Code addition, chunk 2:** as the last line of `main`, type:

```go
	fmt.Println("Version:", Version)
```

> "At the bottom of `main` I'm printing `Version`. I'm saving with Command-S, running `go run .` in the terminal and tracing the output: string, int, bool, float32, the zero values, the inferred type, the walrus variable, and finally the constants. Version prints 1.0."

## 4. Lesson 03, user input (about 4 min)

**Show:** `03_user_input/main.go`.

> "Next is `03_user_input/main.go`. This program reads a rating with `bufio` and `ReadString`. The comments explain that `ReadString` returns a string and an error, that Go has no try/catch because errors are returned as values, and that the underscore discards the error. The old comment said the input was an integer, which was wrong: it always returns a string, even when the user types digits, and it keeps the trailing newline. Let me run it first with `go run .` and type 4.5. It prints the value and says the type is string."

**Code addition, chunk 1:** add `"strconv"` and `"strings"` to the import block, after `"os"`:

```go
	"strconv"
	"strings"
```

> "I'm adding `strconv` and `strings` to the import block. One converts text to numbers, the other trims whitespace."

**Code addition, chunk 2:** replace the `input,_ :=` line with:

```go
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input:", err)
		return
	}
	input = strings.TrimSpace(input)
```

> "I'm replacing the underscore with `err` and checking it, since ignoring errors hides problems. Then `strings.TrimSpace` removes the trailing newline."

**Code addition, chunk 3:** replace the last two print lines with:

```go
	rating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Rating must be a number:", err)
		return
	}

	fmt.Println("Thanks for rating, ", rating)
	fmt.Printf("type of rating: %T\n", rating)
	fmt.Println("Rating plus 1:", rating+1)
```

> "`strconv.ParseFloat` converts the string to a 64-bit number and returns an error if it can't. If there's an error, I print a message and return. Otherwise I print the rating, its type, and the rating plus one to prove it's a number."

**Run:**

> "I'm saving with Command-S and running `go run .`. When prompted, I'll type 4.5. The output shows type of rating is float64, and rating plus 1 is 5.5. Now I'll run it again and type abc to check the failure path. It prints a clear error and exits."

## 5. Lesson 04, time (about 3 min)

**Show:** `04_time/main.go`, then `04_time/go.mod`.

> "I'm opening `04_time/main.go`. It prints the current time, then a formatted version. The comment explains that Go formats dates using a reference time, Monday January 2, 2006 at 15:04:05, so the layout 01-02-2006 means month, day, year. Let me run it with `go run .`."

**Fix on camera:** run `go run .` to show the error, then edit `04_time/go.mod`:

```
module timelesson
```

> "The run fails with an ambiguous import error. The module is named `time`, which collides with Go's standard `time` package. I'm opening `go.mod` in the Explorer and renaming the module to `timelesson`, then saving with Command-S."

**Code addition:** in `main.go`, under `fmt.Println(createdDate)`:

```go
	fmt.Println(createdDate.Format("02 Jan 2006"))
	fmt.Println("Hours since created:", int(time.Since(createdDate).Hours()))
```

> "Now I'll add two lines under the created date. The first formats it as day, short month, year. The second uses `time.Since` to get the duration since that date, converts it to hours and prints it as a whole number. I'm saving with Command-S and running `go run .`. The output shows the current time, the formatted time, the created date, the short date 20 Dec 2002 and the hours elapsed."

## 6. Final check and epilogue (3-4 min, 4:00 max)

**Show:** the VS Code Explorer with the four lesson folders, all saved (no unsaved dots on the tabs).

> "All four lesson files are saved. In the Explorer on the left, `01_hello`, `02_variables`, `03_user_input` and `04_time` each contain their `go.mod` and Go file. These refined lessons are my final output. After this recording I'll copy them into a `final output 1` folder with the Goal document and the original lessons, and zip everything."

**Epilogue:**

> "I verified three things. First, all four lessons compile and run after my edits. Second, the rating input works: 4.5 gives a float64 and 5.5, and abc gives an error. Third, the `time` module name now builds. The limitations are that the program only checks for a number, not a 1 to 5 range. The time output depends on my local clock and timezone. I tested on macOS only.
>
> My main advice for beginners is to read the compiler and run output, not just the comments. Comments can be wrong. Always check what a function really returns, and check errors instead of discarding them with an underscore. Also avoid naming your module after a standard library package."

## After recording

Packaging is not part of the recording.

1. Replace the contents of `Deliverables/final output 1` with your final on-camera lesson folders (`01_hello` to `04_time`, each with `go.mod`).
2. Check that `Deliverables` has `input 1.docx`, `input 2/` and `final output 1/`.
3. Zip it: `zip -r Deliverables.zip Deliverables -x "*.DS_Store"`.
4. Run `unzip -l Deliverables.zip` and confirm every file is there.
5. Keep this script, the metadata RTF, `video.mp4` and the TaskRecorder archive outside the zip.
