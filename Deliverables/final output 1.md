# Go from Scratch: Arrays and Slices

## Recording plan

**Target runtime: about 20–24 minutes**, including switching between Safari, the editor, and the terminal, plus running both examples. Speak continuously at a comfortable pace. When performing an action, name what you are doing, where it is on screen, and why.

| Time | Segment |
|---|---|
| 00:00–02:30 | Intro and official Go documentation |
| 02:30–09:00 | Lesson 07: arrays |
| 09:00–20:30 | Lesson 08: slices |
| 20:30–23:00 | Recap, practice, and epilogue |

## Before recording

- In Safari, open the official Go documentation at <https://go.dev/doc/>. The standard library references are at <https://pkg.go.dev/std>.
- Open `07_arrays/main.go` and `08_slices/main.go` in the editor.
- Keep a terminal visible. Run each lesson from its own directory:

  ```sh
  cd 07_arrays
  go run .
  ```

  Then run:

  ```sh
  cd ../08_slices
  go run .
  ```

- This repository’s slices module uses a module path that avoids colliding with Go’s standard-library `slices` package.
- The packaged source copies are `input 1.go` (arrays) and `input 2.go` (slices). In the `Deliverables` directory, each can also be run directly with `go run "input 1.go"` or `go run "input 2.go"`.

## Narration script

### 00:00 — Introduction and the Go documentation

**Narration**

“Welcome. Today we’re continuing our Go lessons with arrays and slices. The examples are in this repository’s `07_arrays` and `08_slices` folders. We’ll start with arrays, which have a fixed length, and then move to slices, which are the more flexible sequence type used in many Go programs.

I’m opening Safari at `go.dev/doc`. This is the official Go documentation site. It provides beginner tutorials, language references, and explanations of the Go tools. The Go standard library documentation is also searchable at `pkg.go.dev/std`. That is where I can look up packages such as `fmt` and `sort`, including what their functions accept and return.

Documentation is especially useful when learning collections because the names can sound similar while the behavior is different. An array has a length that is part of its type. A slice is a small, flexible view over an array. We’ll see these differences in the source and verify them by running both examples.

I’m switching from Safari to the editor and opening `07_arrays/main.go`. Then I’ll run each lesson from its own module directory so we can connect the code to its output.”

**Show:** The Go documentation in Safari, then the `07_arrays` and `08_slices` folders in the editor.

### 02:30 — Lesson 07: arrays

**Narration**

“The first code example is `07_arrays/main.go`. At the top, `package main` identifies this as a runnable Go program, and `import "fmt"` makes Go’s formatted printing functions available. The program starts in `func main()`. Its first statement prints a welcome message so we can tell which example produced the terminal output.

Now look at `var names [3]string` in the editor. This declares an array named `names`. The type is `[3]string`: the number three is part of the type and means the array contains exactly three string positions. The element type is `string`, so each position stores text. Unlike a slice, this array cannot grow or shrink.

Go indexes arrays starting at zero. The next three assignments fill the positions: `names[0]` gets `kamy`, `names[1]` gets `kamy1`, and `names[2]` gets `kamy2`. The last valid index is two because the three positions are numbered zero, one, and two. Trying to use index three would be outside the array and cause a bounds error.

The first `fmt.Println` prints the contents of `names`. The next call prints `len(names)`. The built-in `len` function reports how many elements the array contains, which is three. The array’s length is fixed whether or not we have assigned every position. If a string array position had not been assigned, it would contain the empty string, because that is the zero value for a string.

The source has comments pointing out this fixed length. The important idea is that an array’s capacity is determined by its type declaration, not by how many values we happen to display. Arrays are useful when a fixed-size collection is part of the problem or the type itself needs to express a fixed number of elements.

The next declaration uses a composite literal: `var places = [3]string{"USA","UK","Russia"}`. This creates another three-string array and initializes all three positions in one expression. Go infers the type of the variable from the array literal, so we don’t have to write a separate type after `places`.

There are other array literal forms in Go, including a form where the compiler counts the elements, but this example explicitly writes `[3]string` so the fixed length is easy to see. The array still has the same properties: three elements, indices from zero through two, and a type that includes its length.

The final comment says arrays are not used as often as some other collection types in typical Go code. That’s a useful observation, but arrays are still important: slices are built on top of arrays, so understanding fixed-size storage helps explain how slices behave.

I’m moving to the terminal panel and running `go run .` from the `07_arrays` directory. I’m doing that to execute this module and compare the printed arrays and length with the declarations in the editor.

The output shows `names` with three strings, then the length three, and then the `places` array. The output confirms that the array keeps all three positions and that `len` reports the number of positions.”

**Show and do:** Point to the declaration and assignments in `07_arrays/main.go`, then run `go run .` from `07_arrays` and compare the output to the source.

### 09:00 — Lesson 08: slices

**Narration**

“Now I’m opening `08_slices/main.go` in the editor. This example imports `fmt` for printing and `sort` for sorting integers. The `main` function begins by printing a welcome message.

The first collection is declared as `var names = []string{"Kamy", "Kamy1", "kamy2"}`. The brackets are empty before `string`, which means this is a slice of strings, not a three-element array. The literal initializes the slice with three values. Go infers the slice type as `[]string`.

The next print statement displays the values. The `fmt.Printf` line uses `%T`, which prints the type of the value. The output is `[]string`. That is a quick way to confirm that the code created a slice rather than an array.

A slice is a descriptor over an underlying array. You can think of it as keeping track of where its elements begin, how many elements are in its current view, and how much of the underlying storage is available before it needs more storage. The length is what `len` reports. The capacity, reported by `cap`, is the number of elements from the start of the slice to the end of its available underlying array. This example doesn’t print capacity, but the distinction explains why slices can be extended with `append`.

Next the program calls `append(names, "Kamy3", "Kamy4")` and assigns the result back to `names`. `append` returns a slice containing the original elements plus the new ones. We assign the returned value because the resulting slice may have a different length and may refer to a different underlying array if more space was needed. The output now has five names.

After that, the source says `names = append(names[1:])`. Let’s unpack it carefully. `names[1:]` is a slice expression: start at index one and continue through the end. Slice ranges include the starting index and exclude the ending index; when the end is omitted, the slice goes to the end. Here, `names[1:]` selects all names except the first one.

There are no new elements passed to `append` in this expression. In this example, `append` simply returns that subslice, so the assignment makes `names` refer to the tail beginning at the second name. Writing `names = names[1:]` would express the intention more directly. I’m explaining the exact source first because it’s useful to distinguish the slice operation from the append operation.

The next statement similarly assigns `append(names[1:3])`. At this point, `names` is already the shortened slice. The range `[1:3]` selects elements at indices one and two of that current slice; it excludes index three. Again, because no elements are supplied to append, the call effectively returns the selected subslice. The output shows the resulting shorter view. In new code, use the direct expression `names = names[1:3]` when the goal is only to take a subslice.

Now we create `nums` with `make([]int, 3)`. The built-in `make` function allocates and initializes a slice of integers with length three. Those three integer positions start at their zero value, zero. The program assigns 100, 200, and 300 at indices zero, one, and two. This is valid because the slice length is three.

The commented-out line `nums[3] = 400` is a useful warning. At this point, the valid indices are zero, one, and two. Index three is not in the slice’s current length, so uncommenting that assignment would cause a runtime index-out-of-range panic. Having spare capacity would not make index three directly assignable; indexing is bounded by the current length.

To add elements, the program instead uses `nums = append(nums, 350, 650, 250)`. This adds three values and assigns the returned slice back to `nums`. The length grows from three to six. Go manages the slice’s available backing storage as needed; programs should rely on the slice’s length and append behavior, not assume a particular capacity-growth formula.

Next, `sort.Ints(nums)` sorts the integer elements into ascending order. This function modifies the slice’s elements rather than returning a separate sorted slice. The following print displays the sorted values. Then `sort.IntsAreSorted(nums)` checks whether the values are in ascending order and returns a Boolean. The variable `sort_nums` stores that Boolean, and the next print should show `true` after the sort.

The last part demonstrates removing an element from a slice by index. The `frds` slice contains five names, and `index` is set to two. Since indexing starts at zero, index two identifies the third name, `Patrick`.

The expression `frds[:index]` selects everything before index two. The expression `frds[index+1:]` selects everything after the element at index two. The three dots after that second slice expand its elements so they can be passed as individual values to `append`. Appending the later names to the earlier names creates a slice that omits `Patrick`.

The result is assigned back to `frds` because `append` returns the updated slice. The final print shows the list without the name at index two. This is a common slice-removal pattern, but remember that slices can share an underlying array. If another slice refers to the same storage, changes to elements may be visible through that other slice too.

I’m switching to the terminal panel and running `go run .` from the `08_slices` directory. I’m running the example now so we can check the values after appending, slicing, sorting, and removing an element.

The output starts with the original names and their `[]string` type. It then shows the longer slice after append, the shorter views after the slice expressions, the integer values before and after append and sorting, the `true` result from the sorted check, and finally the friends list without the element at index two.”

**Show and do:** Point to each operation as it is discussed. Run `go run .` from `08_slices`, then trace the output in order. Do not uncomment `nums[3] = 400` during the recording; explain why it would fail instead.

### 20:30 — Recap and epilogue

**Narration**

“Let’s recap the difference between arrays and slices. An array’s length is fixed and is part of its type, such as `[3]string`. A slice has a flexible length and refers to an underlying array. Use `len` to check the current number of elements. Use `append` to add elements, and always use the slice returned by `append`. Use slice expressions to select a range, remembering that the start is included and the end is excluded.

We also used `make` to create a slice with a chosen length, `sort.Ints` to sort a slice of integers, and `sort.IntsAreSorted` to check the ordering. Finally, we removed an element by appending the portion before it to the portion after it.

For practice, create a slice of five favorite foods. Print its length, append one more food, and print it again. Then remove the element at index one and explain which value was removed. After that, create a fixed array of three strings and compare its type with the slice type using `%T`.

When you want to confirm the exact behavior of a built-in or library function, consult the official Go documentation and the standard library reference. That’s the end of this arrays and slices lesson. Thanks for learning with me.”

**Epilogue:** Leave the recap or practice task visible briefly, then end the recording.

## Action narration cues

Say these while the action is happening. Each one identifies what is changing, where it is on screen, and why:

- “In the editor, I’m opening `07_arrays/main.go` so we can start with the fixed-length array example.”
- “In the terminal panel, I’m changing into `07_arrays` and running `go run .` so we can verify the array values and length.”
- “In the editor, I’m opening `08_slices/main.go` to compare a flexible slice with the fixed-length array we just ran.”
- “In the terminal panel, I’m changing into `08_slices` and running `go run .` to check the append, sort, and removal results.”
- “In the editor on the `nums[3]` example, I’m leaving the line commented out because index three is outside the slice’s current length and would panic.”

For any additional command or edit, narrate what you are doing, identify its location on screen, and explain why. Avoid unexplained typing or edits during the recording.

## Instructor accuracy notes

- Array lengths are part of array types; arrays are values and assignment copies the array.
- A slice is a descriptor over an underlying array. Its length and capacity are distinct.
- `append` may reuse existing storage or allocate new storage; do not promise a specific capacity-growth rule.
- The expressions `append(names[1:])` and `append(names[1:3])` add no elements; they return the provided subslice. Direct slicing expresses this intent more clearly.
- A slice index must be less than its current length, even when its capacity is larger.
- `sort.Ints` sorts the provided elements in place; `sort.IntsAreSorted` returns a Boolean.
- The removal expression uses `...` to expand a slice into the variadic arguments accepted by `append`.
- Slice aliases may share backing storage, so mutations can be visible through multiple slices.
