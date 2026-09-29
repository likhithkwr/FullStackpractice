# Two Sum — Go

Read the [Two Sum explanation](../../Explanation/two-sum.md) for the algorithm, diagrams, and a comparison with Python and Java. The [terminology guide](../../Explanation/terminology.md) explains the Go syntax used here.

## Files

| File | Purpose |
| --- | --- |
| [practice/main.go](practice/main.go) | Your local practice file. Replace the `twoSum` TODO and run its eight checks. |
| [reference/main.go](reference/main.go) | Completed hash map and brute-force solutions, with three runnable examples. |
| [reference/main_test.go](reference/main_test.go) | Checks both reference approaches on examples and edge cases. |
| [go.mod](go.mod) | Declares this exercise's Go module and minimum Go version. |

The two programs live in separate folders so each can define its own `twoSum` function and `main` entry point.

## Practice locally

Install Go 1.21 or newer. From the repository root, run:

```sh
go -C Go/two-sum run ./practice
```

Open `practice/main.go` in your editor and replace the `panic(unfinished)` placeholder inside `twoSum` with your implementation. Keep the return type `[]int`. The problem guarantees a pair for valid input; a final `return nil` after your loop can handle input outside that guarantee.

The checker accepts either index order, validates positions against the original input, and requires two distinct indices. The first run reports eight `TODO` cases because your function is intentionally unfinished. The practice program uses exit code `2` for unfinished work, `1` for a failed answer or unexpected panic, and `0` when all checks pass. The `go run` wrapper prints `exit status 2` for the initial unfinished program and returns a nonzero command status; read the printed summary for your results.

Add custom cases to the `cases` slice in `practice/main.go`, supplying `nums`, `target`, and the expected indices.

## Run the reference

From the repository root:

```sh
go -C Go/two-sum run ./reference
go -C Go/two-sum test ./...
```

The test command exercises the completed reference. Your practice file reports its own results when you run it. All code uses the Go standard library; there are no packages to download.

## Go syntax to notice

| Go code | Meaning |
| --- | --- |
| `package main` and `func main()` | Define a runnable program and its starting point. |
| `nums []int` | `nums` is a slice of integers. |
| `seen := make(map[int]int)` | Create a writable map from a number to its earlier index. |
| `for i, value := range nums` | Visit each index and value. |
| `if earlierIndex, found := seen[complement]; found` | Retrieve the stored index and check whether the key actually exists, including when the stored index is `0`. |
| `seen[value] = i` | Remember the current number and position after looking for its complement. |
| `return []int{earlierIndex, i}` | Return a two-element integer slice. |

The completed hash map approach takes expected `O(n)` time and `O(n)` extra space. The brute-force approach takes `O(n²)` worst-case time and `O(1)` extra space.
