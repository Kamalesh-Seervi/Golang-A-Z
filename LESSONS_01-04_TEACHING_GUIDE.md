# Go from Scratch: Lessons 01–04

## Recording plan

**Target runtime: about 22–26 minutes**, including pauses while switching tabs and running the examples. Speak at a comfortable pace, keep the editor and terminal visible, and explain what is on screen rather than saying “here” or “this.”

| Time | Segment |
|---|---|
| 00:00–02:30 | Intro and official Go documentation in Safari |
| 02:30–06:00 | Lesson 01: first Go program |
| 06:00–12:00 | Lesson 02: variables, types, constants, and zero values |
| 12:00–17:30 | Lesson 03: reading user input |
| 17:30–22:30 | Lesson 04: current time, formatting, and dates |
| 22:30–24:00 | Recap, practice task, and epilogue |

The time ranges are a guide, not a reason to rush. Pauses to type and run the code are included in the target.

## Before recording

- Open Safari and go to the official Go documentation: <https://go.dev/doc/>.
- Open this repository in the editor, with `01_hello/hello.go` ready.
- Keep a terminal available. Run each lesson from its own directory, for example:

  ```sh
  cd 01_hello
  go run .
  ```

- Repeat from `02_variables`, `03_user_input`, and `04_time`. Each directory has its own `go.mod` and is a separate Go module.
- For the input lesson, type a sample rating and press Enter.

## Narration and teaching notes

### 00:00 — Introduction and the Go documentation

**Narration**

“Welcome. In this lesson, we’ll learn Go from the beginning by working through the first four folders in this repository: `01_hello`, `02_variables`, `03_user_input`, and `04_time`. I’ll explain the purpose of each important line, run the programs, and point out a few details that are useful when you start writing your own Go code.

I’m starting in Safari at `go.dev/doc`. This is the official Go documentation website. Go is a programming language, and the documentation site is the reference and learning resource maintained for the Go project. It includes tutorials for getting started, language references, explanations of tools and packages, and guidance for writing Go programs. It is not the code for this repository; it is a place to learn how the language and its standard library work.

For a beginner, the tutorials and the Go Tour are useful guided introductions. The language specification is the precise reference for the rules of Go, but you don’t need to read it cover to cover to begin. When a function such as `fmt.Println` or `time.Now` appears in a lesson, the documentation can tell you what that function does and what values it accepts.

Now I’m switching to the project. The numbered folders are small, separate examples. We’ll start with the smallest possible program and gradually add variables, keyboard input, and dates.”

**Show:** Safari’s address bar and the Go documentation page, then the four project folders.

### 02:30 — Lesson 01: `01_hello`

**Narration**

“The first example is in `01_hello/hello.go`. It prints a greeting. Before we run it, let’s understand the structure.

The line `package main` says this source file belongs to the `main` package. In Go, a package groups related source files. A program that we want to run directly uses the special package name `main`. The `main` package also needs a `main` function, because that is where the program starts.

The import statement, `import "fmt"`, makes the standard library’s formatting package available to this file. `fmt` provides functions for printing text and formatted values. Go source files must import the packages they use.

Next, `func main()` defines the entry-point function. The braces mark the body of the function. Inside it, `fmt.Println("Hello, World!")` prints the text followed by a newline. The text between double quotes is a string literal: a fixed piece of text written directly in the source.

I’ll run this module from its folder with `go run .`. Go reads the module information and runs the program. The output is `Hello, World!`.

There’s a comment underneath the program suggesting that typing `fmt.Println` automatically imports `fmt`. That is not a Go language feature. An editor or IDE may offer an automatic import suggestion, but the Go compiler requires the import to be present in the source. It’s useful to know the difference between help from an editor and a rule of the language.

The `go.mod` file in this folder says the module path is `example/hello` and specifies Go version 1.20. A module is a unit for managing Go code and dependencies. For this tiny example, there are no external dependencies. The module file still gives Go the module’s identity and language/toolchain version information.

The key ideas from the first program are: a runnable program uses `package main`, execution starts at `func main()`, imports make packages available, and `fmt.Println` writes a line of output.”

**Show and do:** Read `01_hello/hello.go`, point to each construct as it is named, run `go run .`, and show the output.

### 06:00 — Lesson 02: `02_variables`

**Narration**

“The second example is `02_variables/main.go`. It introduces values and types. A variable is a named place to store a value so the program can use it later.

At the top, `const token int = 134325` declares a constant named `token` with type `int`. A constant’s value is fixed rather than reassigned while the program runs. This declaration is at package level, outside `main`, so the function can use it. The comment calls it public, but that wording is not correct in Go. Go determines whether an identifier is exported from a package by its first letter: an uppercase name such as `Token` is exported, while lowercase `token` is not. Here, `token` is package-level but unexported.

Inside `main`, `var name string = "Kamalesh"` declares a variable, states its type as `string`, and gives it an initial value. A string stores text. `fmt.Println("Hello, ", name)` prints the greeting and the variable. `fmt.Printf` is different from `Println`: it lets us choose a format. The `%T` verb means ‘print the type of this value,’ and `\n` adds a newline. So this line prints that `name` has type `string`.

The integer example declares `num` as an `int` and prints it. An integer stores a whole number. The exact size of `int` depends on the platform, so use a specific integer type only when your program needs that exact size.

The Boolean example declares `case1` as `bool`. A Boolean value is either `true` or `false`, useful for representing a condition.

Next, `flt` is declared as `float32`, which stores a floating-point number. Floating-point values are approximate representations, not exact decimal arithmetic. The source literal has more decimal digits than a `float32` can preserve, so the printed result is rounded. Go also has `float64`, which generally gives more precision and is the common default for floating-point literals when Go infers their type.

The next two declarations demonstrate zero values. `var new_num int` has no explicit initial value, so Go initializes it to the integer zero. `var named string` is initialized to the empty string, which prints as an empty value. Go initializes variables rather than leaving them with arbitrary uninitialized contents. Other common zero values are `false` for a Boolean and `nil` for certain reference-like types.

Then `var nametyp = "Kamalesh007"` uses type inference. We wrote an initial string value, so the compiler infers that `nametyp` is a string. We could also write `nametyp := "Kamalesh007"`. The `:=` short variable declaration is concise and also infers the type. It can be used inside a function, such as `main`, but not as a standalone package-level declaration. At package level, use a declaration such as `var` or `const`.

Finally, the program prints `token`, which was declared outside the function. This shows that a function can refer to a name declared at package level in the same package.

I’ll run the program and read the output alongside the source. Notice that `%T` reports the type, while `Println` prints the value. They answer different questions.

As a quick practice task, change `name`, then run the program again. Add a Boolean variable set to `false` and print both its value and type. Think about why an explicitly initialized declaration and a zero-valued declaration can both be useful.”

**Show and do:** Explain each variable group in the source, run `go run .`, and point out the printed values and `%T` output.

### 12:00 — Lesson 03: `03_user_input`

**Narration**

“The third example, `03_user_input/main.go`, reads a line from the keyboard. It brings in three packages: `bufio`, `fmt`, and `os`.

`fmt` prints the welcome message and prompt. `os` gives access to operating-system resources; `os.Stdin` represents the program’s standard input, which is usually the keyboard in a terminal. `bufio` provides buffered input and output helpers. A buffer collects input efficiently and lets us read it in useful chunks.

The line `welcome := "hi Users"` creates a string variable using short declaration. `fmt.Println(welcome)` displays it.

Next, `bufio.NewReader(os.Stdin)` creates a buffered reader connected to standard input. We store it in `reader` so the program can use it to read user input. Then the program prints a prompt asking for a pizza rating.

The next line is `input, _ := reader.ReadString('\n')`. `ReadString` reads characters until it finds the delimiter provided as an argument. The single-quoted `'\n'` is the newline character, so the program waits for the user to type a line and press Enter. `ReadString` returns two values: the text it read and an error value. The underscore is Go’s blank identifier; it intentionally discards a value. In this example, the code discards the error, which keeps the demonstration short, but production code should check errors so it can handle end-of-input or another read failure deliberately.

The returned input is a string. That is true even if the user types digits like `5`. Text that looks like a number is not automatically an integer. To do arithmetic with a rating, we would validate and convert it, for example with `strconv.Atoi`, and handle the possible conversion error.

There is another detail: because the delimiter is a newline, the input string usually includes the newline that the user entered. The print statement may therefore appear with extra spacing. A common next step is to remove surrounding whitespace with `strings.TrimSpace`.

The program prints a thank-you message, then uses `%T` to confirm that `input` is a string. I’ll run it now. I’ll type `5` and press Enter. The output shows the text we typed and confirms the value’s type.

The important distinction is between reading and interpreting input. This program reads a line as text. It does not yet verify that the text is a valid rating or convert it to a number. Those are separate steps we can add when the program needs numeric behavior.”

**Show and do:** Run `go run .`, type `5`, press Enter, and explain the newline and `string` type shown in the output.

### 17:30 — Lesson 04: `04_time`

**Narration**

“The fourth example is `04_time/main.go`. It uses Go’s standard `time` package, which contains types and functions for representing times, dates, durations, and time zones.

The import block includes `fmt` for printing and `time` for working with dates and times. The first print statement displays a welcome message. Then `mytime := time.Now()` asks the system for the current time and stores the result in `mytime`. The type is `time.Time`. Printing it directly displays a default human-readable representation, often including a time zone and offset. The exact value changes each time we run the program.

The next line calls `mytime.Format("01-02-2006 15:04:05 Monday")`. Go’s time formatting uses a specific reference date rather than familiar placeholders such as `YYYY` and `MM`. In the reference layout, the number `01` represents the month, `02` the day, `2006` the year, `15` the hour in 24-hour time, `04` the minute, and `05` the second. `Monday` represents the weekday name. The chosen order in this string produces a result like month-day-year followed by time and weekday. This layout can look unusual at first, so remember that these digits are examples from Go’s reference time, not arbitrary formatting codes.

The final declaration calls `time.Date`. Its arguments are year, month, day, hour, minute, second, nanosecond, and location. In this example, the date is December 20, 2002, at 20:20:00 in the local time zone. The month is written as `time.December`, which is clearer than guessing a numeric month value. The final zero is the nanosecond field, and `time.Local` says to interpret the date in the computer’s local time zone.

`time.Date` constructs a time value; it does not mean that this is the current time. We print `createdDate` to show the constructed date. If you run the program on another computer or in another time zone, the displayed local-time details may differ.

I’ll run the example. The first time is the current time, while the last one is the fixed date we constructed in the source. That contrast is the point: `time.Now()` asks for now, and `time.Date(...)` builds a particular date and time.

One practical note: dates and times can be affected by time zones and daylight-saving rules. For real applications, be explicit about whether a time represents local time or a globally consistent instant, and check the `time` package documentation when handling conversions.”

**Show and do:** Run `go run .`, identify the current-time line, formatted line, and constructed date in the output.

### 22:30 — Recap and epilogue

**Narration**

“Let’s recap what we learned in the first four Go examples.

In `01_hello`, we saw a runnable `main` package, the `main` function, an import, and a call to `fmt.Println`. In `02_variables`, we declared strings, integers, Booleans, and floating-point values; looked at inferred types, constants, and zero values; and used `%T` to inspect types. In `03_user_input`, we connected a buffered reader to standard input and learned that a line from the keyboard is text, even when it contains digits. In `04_time`, we used `time.Now`, Go’s reference-date formatting layout, and `time.Date` to construct a specific date.

For practice, make a small program that asks for a name, reads it, and prints a greeting with the current date. Build it one idea at a time: print a prompt, read the input, then use the `time` package. Check the official Go documentation when you want to understand a package or function more deeply.

That’s the end of this introduction. Thanks for learning Go with me. In the next lessons, we can build on these fundamentals with conversions, arrays, and other features.”

**Epilogue:** Leave the final recap on screen briefly, then end the recording.

## Accuracy notes for the instructor

- Go does not automatically add imports. Editors may offer to insert them.
- `token` is not exported because its name begins with a lowercase letter; package-level does not mean public.
- `ReadString` returns both a string and an error. The sample ignores the error and keeps the newline in the string; mention both limitations.
- The sample asks for a pizza rating but does not parse or validate a number.
- Go’s `time.Time.Format` uses the reference timestamp `Mon Jan 2 15:04:05 MST 2006` as its layout model.
