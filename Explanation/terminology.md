# Programming and Technical Terminology

Plain-language definitions for the [Two Sum guide](two-sum.md), its language implementations, and the [URL shortener design guide](url-shortener.md).

## 1. Understanding the problem

| Term | Meaning | Example |
| --- | --- | --- |
| Integer | A whole number, including negative numbers and zero. | `-3`, `0`, `7`. |
| Array | An ordered collection of elements addressed by their positions. A Java array has a fixed length. | `int[] nums = {2, 7, 11, 15};` |
| List | An ordered collection. The Python example uses a list, whose length can change. | `nums = [2, 7, 11, 15]`. |
| Element | One item in a collection. | `7` is an element of `nums`. |
| Index / indices | A position / multiple positions in a collection. Java and Python use zero-based indices. | Index `1` identifies the second element. |
| Zero-based indexing | Counting positions from `0` instead of `1`. | Four elements have indices `0`, `1`, `2`, and `3`. |
| Value | The data stored at a position. | The value of `nums[1]` is `7`. |
| Input | Data supplied to the program or method. | The array and target. |
| Output | The result produced by the program or method. | The returned indices `[0, 1]`. |
| Target | The desired total in this problem. | `9`. |
| Pair | Two selected items. | Values `2` and `7`. |
| Distinct elements | Elements at different positions; their values may be equal. | The two `3`s in `[3, 3]`. |
| Duplicate | A value that occurs more than once. | `3` occurs twice in `[3, 3]`. |
| Complement | The value needed to complete the current number's pair. | For current `7` and target `9`, the complement is `2`. |
| Constraint | A rule defining allowed input or valid behavior. | The same array position cannot be used twice. |
| Assumption / guarantee | A condition the problem states you may rely on. | Exactly one valid pair exists. |

## 2. Algorithms and data structures

| Term | Meaning | Connection to Two Sum |
| --- | --- | --- |
| Algorithm | A precise sequence of steps that solves a problem. | Compute a complement, look it up, then return or store. |
| Data structure | A way to organize data for particular operations. | The array holds input; the map holds earlier values and indices. |
| Brute force | Try all possible candidates until a valid one is found. | Check every distinct pair. |
| Loop | Code that repeats while a condition permits it. | Visit each array position. |
| Iteration | One repetition of a loop. | Process the element at index `2`. |
| Nested loops | A loop inside another loop. | For each `i`, try every later `j`. |
| One pass | A scan in which each input position is visited at most once. | The hash map solution moves left to right. |
| Map / dictionary | A structure associating a key with a stored value. | A number maps to an earlier index. Python calls this a dictionary. |
| Key | The item used to identify a map entry. | In `3 -> 0`, the key is the number `3`. |
| Map value | The data stored under a key. It need not mean an array value. | In `3 -> 0`, the stored value is index `0`. |
| Hash map / hash table | A structure that uses hashing to store and retrieve entries efficiently. Hash maps associate keys with stored values. | Java's `HashMap` and Python's dictionary provide the lookup used here. |
| Hash function | A calculation that turns a key into a hash value, which helps choose where an entry is stored. | The map uses the number's hash when finding its index. |
| Bucket | An internal storage location for entries assigned to the same part of a hash table. | A bucket can contain more than one key. |
| Hash collision | Different keys are assigned to the same hash or storage bucket. The map must still distinguish them. | Collisions can affect cost, but they do not make unequal keys the same key. |
| Lookup / membership check | Find stored information / check whether a key exists. | `get` retrieves an index; `containsKey` checks presence. |
| Insertion / update | Add an entry / replace the value for an existing key. | `seen.put(nums[i], i)` records the current index. |
| Set | A collection of distinct values used for membership checks. | A set can report that a complement exists, but the map also retrieves its index. |
| Pseudocode | Steps written without requiring the exact syntax of a particular language. | "If the complement is present, return the two positions." |
| Dry run / trace | Follow an algorithm by hand and record its changing state. | The example tables show the map before each lookup. |
| State | The values held by a program at a particular point. | The current index, complement, and map contents. |
| Loop invariant | A statement that stays true at the same point in each loop iteration. | Before checking index `i`, the map contains only earlier positions. |
| Correctness proof | Reasoning that a method finds a valid answer for every allowed input. | The invariant prevents index reuse and ensures a valid pair is found. |
| Edge case | An input near a boundary or with a special property that may expose mistakes. | Two elements, zeros, negatives, or equal values. |
| Two pointers | Two variables tracking positions, often at opposite ends of a sorted array. | Move the positions based on whether their sum is too small or too large. |

## 3. Performance terminology

| Term | Meaning | Example |
| --- | --- | --- |
| `n` | The number of input elements. | `[2, 7, 11, 15]` has `n = 4`. |
| Time complexity | How the amount of work grows as input size grows. | Scanning once has linear growth. |
| Space complexity | How memory requirements grow as input size grows. Here we discuss extra memory beyond the input. | Remembering earlier values can require `O(n)` extra space. |
| Big O | Notation for an upper bound on growth, ignoring constant factors and lower-order terms. | `O(n)` describes linear growth. |
| `O(1)` / constant | Growth bounded independently of input size in the relevant operation. | The returned two-index array always holds two indices. |
| `O(n)` / linear | Growth proportional to input size as an upper bound. | One full scan. |
| `O(n^2)` / quadratic | Growth proportional to the square of input size as an upper bound. | Checking all pairs. |
| `O(n log n)` | Growth from a linear factor and a logarithmic factor. A logarithm grows slowly; repeated halving helps explain it. | Sorting value/index pairs before a two-pointer scan. |
| Worst case | The largest cost among allowed inputs for the chosen analysis model. | The matching pair is found only after many other pairs are checked. |
| Expected complexity | Cost under stated assumptions about how a structure behaves, such as suitable hash distribution. It is not a guarantee for every operation. | Expected constant-time hash map lookup. |
| Amortized complexity | Cost averaged over a sequence of operations, including occasional expensive operations. | Many cheap insertions share the cost of an occasional map resize. |
| Tradeoff | A benefit that comes with a cost elsewhere. | The map reduces search work while using more memory. |

Big O does not state an exact number of seconds. Doubling `n` roughly doubles a full linear scan's work; the number of all-pairs checks grows by roughly four times. Early returns, machine speed, and implementation details also affect actual timing.

## 4. Reading the Java code

| Term or syntax | Plain-language meaning |
| --- | --- |
| Data type | Describes what kind of data a variable can hold. `int` is a whole-number type; `int[]` is an integer-array type. |
| Variable | A named place to hold a value or reference. `complement` holds the required partner. |
| Reference | A way to access an object. `seen` refers to the map created by `new HashMap<>()`. |
| Local variable / scope | A variable declared inside a method or block / the part of the program where it can be used. Each method call has its own local `seen` variable. |
| Primitive `int` | Java's 32-bit signed integer type. The supplied problem's numbers, sums, and complements fit within its range. |
| `Integer` | An object wrapper for `int`, used where Java generic types require reference types. |
| Boxing / unboxing | Conversion between a primitive and its wrapper object. Java performs it automatically here: the code stores `int` values as `Integer` entries and converts a retrieved index back to `int`. |
| Class | A named definition grouping data and behavior. `TwoSum` groups the methods in this example. |
| Object / instance | A particular object created from a class. `new HashMap<>()` creates a map object. |
| Constructor | Code used to initialize a new object. The expression `new HashMap<>()` calls a map constructor. |
| Interface | A specification of operations that an implementing class provides. `Map` describes the map operations; `HashMap` implements them. |
| Method | A named block of code called to perform a task. Java calls the operations declared in a class methods. `twoSum` is a method. |
| Function | A callable piece of code. The Python example defines `twoSum` as a method within a class and `two_sum_brute_force` as a separate function. |
| Parameter | A named input in a method declaration. `nums` and `target` are parameters. |
| Argument | An actual value supplied when calling a method. An array and `9` are arguments in the example call. |
| Method signature | The method's name and parameter types, used to identify it in Java. A declaration also states its return type and modifiers. |
| Modifier | A keyword that controls a declaration's behavior. `public` controls access, and `static` associates this method with the class. |
| Return type / return value | The type the method promises to produce / the actual result it produces. Here these are `int[]` / an array such as `{0, 1}`. |
| `public` | Allows callers in other parts of the program to access this method. |
| `static` | A method or field belongs to the class rather than a particular instance. Call this method as `TwoSum.twoSum(...)`; each call still creates its own local map. |
| `void` | A method does not return a value to its caller. `main` prints output but has no returned answer. |
| `main` | The conventional entry method used to start this standalone Java application. |
| `String[] args` | An array of text arguments supplied when starting the program. The example does not use these arguments. |
| `String` | Java's type for text, such as an error message or a printed sentence. |
| `int[]` / `int[][]` | An array of integers / an array whose elements are integer arrays. The examples variable holds several input arrays. |
| `new` | Creates an object or array. `new int[] {0, 1}` creates the returned answer array. |
| Generics | Type parameters describing the kinds of items a structure accepts. `Map<Integer, Integer>` specifies integer-object keys and values. |
| Diamond syntax `<>` | Lets Java infer generic type arguments for the constructor from context. `new HashMap<>()` uses the types declared on the left. |
| Package / import | A grouping of related Java types / a declaration letting this file use a type's short name. `import java.util.HashMap;` lets the code write `HashMap`. |
| Boolean | A value that is either `true` or `false`. `containsKey` returns a boolean. |
| Condition / `if` | A true-or-false check / code that branches based on such a check. Return the pair if the complement exists. |
| `=` | Assigns a value to a variable. `int i = 0` gives `i` its initial value. |
| `==` / `!=` | Tests equality / inequality. `nums[i] + nums[j] == target` checks the sum. |
| `<` | Tests whether the left value is less than the right value. `i < nums.length` keeps the index within the array. |
| `+` / `-` | Addition / subtraction for the integer values used here. The complement calculation subtracts the current value from the target. |
| `i++` | Increases `i` by one. It advances to the next position after each loop iteration. |
| `nums.length` | The number of elements in the Java array. |
| `nums[i]` | The value stored at array position `i`. |
| Dot `.` | Accesses a field or method. `nums.length` reads the size; `seen.get(...)` calls a map operation. |
| Braces `{ }` | Group a block of code or enclose initial array values, depending on the context. |
| Semicolon `;` | Ends many Java statements and separates the three parts of a traditional `for` loop header. |
| `containsKey` / `get` / `put` | Check whether a key exists / retrieve its value / add or update its value. |
| `return` | Immediately exits the method and sends its result to the caller. It does not merely end the current loop iteration. |
| Exception | An object signaling a problem that interrupts normal execution unless handled. `IllegalArgumentException` signals that supplied arguments do not satisfy the method's requirements. |
| `throw` | Signals an exception. This code throws when no pair exists, which is outside the problem's guarantee. |
| `try` / `catch` | Run code that might throw an exception / handle a particular exception. The test runner checks the documented no-solution behavior this way. |
| `Runnable` / lambda | An interface describing an operation with no arguments and no returned value / a compact way to supply such an operation here. The test passes `() -> TwoSum.twoSum(...)` to its no-solution check and ignores the method's result. |
| `clone` / `Arrays.equals` | Copy the integer array / compare two arrays' contents. The test uses these to check that the input was preserved. |
| `Arrays.toString` | Converts an array's contents into readable text, such as `[0, 1]`, for output. |
| `System.out.printf` | Prints formatted output. In this example `%s` inserts text, `%d` inserts an integer, and `%n` inserts a newline. |
| Comment | Text explaining code to a reader that does not execute as code. Java comments include `// ...` and `/* ... */`. |

The test also uses **logical operators**: `&&` means both conditions must be true, `||` means at least one must be true, and `!` reverses a boolean value.

## 5. Running and checking programs

| Term | Meaning |
| --- | --- |
| Source code | The code a person writes, such as `TwoSum.java` or `solution.py`. |
| Compiler / compilation | A tool / process that translates source code into another form for execution. `javac` compiles Java source into class files. |
| Bytecode / `.class` file | Java instructions for the JVM / a file containing compiled Java class data. |
| JVM | Java Virtual Machine: the execution environment that runs Java bytecode. |
| JDK | Java Development Kit: the tools for developing Java programs, including the compiler and launcher. |
| Runtime / interpreter | An execution environment / a tool that executes program code. The Python example needs Python; Java runs on a JVM. |
| Java launcher `java` | Starts a Java program. It can run a compiled class; with JDK 11+, the single-file form also compiles source as part of launching it. |
| Entry point | Where execution begins for an application. Here it is `main`, which calls `twoSum` with example inputs. |
| Classpath | The locations where Java searches for classes to load. `-cp Java/two-sum/build` identifies the compiled output directory. |
| Terminal / console | The interface for entering commands / the area where the example prints its results. |
| Standard output | The program's ordinary output stream. `System.out` writes to it. |
| Debugging | Investigating how code behaves and finding the cause of a problem. |
| Breakpoint | A place where a debugger pauses execution so you can inspect variables. Put one at the complement calculation to watch the map change. |
| Test case | An input and expected behavior used to check a program. `[3, 3]`, target `6`, must return indices `0` and `1`. |
| Assertion | A check that reports a failure if a required condition is false. The test runner throws `AssertionError` for an incorrect answer. |
| Input preservation | Leaving the supplied data unchanged while computing an answer. Both implementations keep the original array's contents. |

## 6. Words in the industry scenario

| Term | Meaning in the warehouse example |
| --- | --- |
| SKU | Stock keeping unit: an identifier for a kind of product. Candidate packs must contain the same required product. |
| Inventory | The available stock records used to find candidate packs. |
| Snapshot | The records observed when the search begins. Availability may change after those records were read. |
| Stable ID | An identifier for a particular record or pack that stays attached to it when array positions change. |
| API | Application programming interface: the operations through which software components interact. The selection service might return pack IDs through an API. |
| API contract | The agreed requirements for inputs, outputs, and failures. It specifies what happens when no suitable pack pair exists. |
| Concurrency | Requests overlapping in time. Two workers may attempt to select the same pack. |
| Atomic operation | A change that succeeds as a whole or has no partial effect. Reserving two packs must reserve both or neither. |
| Transaction | A group of database operations completed together. Availability checks also need suitable locking or conditional writes to prevent conflicting reservations. |
| Reservation | Claiming a pack for one order so it is no longer available for another order. |

The algorithm finds candidate pairs. The surrounding service uses these additional rules to select compatible items and reserve them safely.

## 7. Local practice terminology

| Term | Meaning |
| --- | --- |
| Reference solution | A completed implementation you can study and compare with your own approach. |
| Practice starter / template | A file with a method for you to implement and code for running example checks. |
| TODO | A comment or message marking work that is unfinished. Replace the placeholder in `twoSum()` with your implementation. |
| Test runner / checker | Code that supplies inputs to your method and validates its returned answer. |
| `PASS` / `FAIL` | The returned indices satisfy the problem / the answer does not satisfy it. |
| `NotImplementedError` / `UnsupportedOperationException` | The Python / Java placeholder exception used here to report the unfinished method as `TODO`. |
| Exit code | A number reported when a program finishes. In these practice files, `0` means all checks passed, `1` means a failure or error, and `2` means unfinished code without a reported failure. |
| `raise` / `except` in Python | Signal an exception / handle an exception. The runner handles the unfinished-method exception separately from other errors. |
| `self` in Python | The instance passed to a method when it is called on an object. This is the conventional name for that first parameter. |
| Main guard in Python | `if __name__ == "__main__":` runs the example checks when this file is launched directly. |
| Git branch | A named line of work. Each learner can develop their practice implementation on a separate branch. |
| Pull request | A proposal to merge changes from one branch into another, with an opportunity for review. |

## 8. Go terms used in Two Sum

| Term or syntax | Plain-language meaning |
| --- | --- |
| `package main` | Declares a runnable Go program. The reference and practice programs live in different folders, so both can define `main` and `twoSum`. |
| `func` / `func main()` | Introduces a Go function / marks where a runnable Go program starts. |
| `[]int` / slice | A Go slice whose elements are integers. A slice has a variable length; `[2]int` would name a fixed-length, two-element array type. |
| `map[int]int` | A map with integer keys and integer stored values. Here each key is a number from `nums`, and each stored value is its earlier index. |
| `make(map[int]int)` | Creates a writable empty map. A declared but uninitialized (`nil`) map cannot accept new entries. |
| `:=` / `=` | Declare a new local variable while inferring its type / assign to an existing variable or map entry. |
| `for i, value := range nums` | Visit a slice from left to right, receiving the position and its value on each iteration. |
| `if earlierIndex, found := seen[complement]; found` | Read both the map value and a boolean reporting whether the key exists. Even an index of `0` has `found == true`. The semicolon separates the lookup from the condition. |
| `[]int{earlierIndex, i}` | A slice literal containing the two returned indices. |
| `nil` | The absence of a slice value. The reference returns it only if the promised solution does not exist. |
| `struct` | A Go type that groups named fields. The practice runner groups each array, target, and expected result into one case. |
| `slices.Clone(nums)` | Creates a copy for the practice checker, preserving the original input for validation. |
| `panic`, `defer`, `recover` | Signal an abrupt failure, schedule cleanup on exit, and catch a panic in scheduled code. The practice runner uses these only to distinguish its unfinished TODO marker from an unexpected error. |
| `go.mod` / module | The file naming a Go module and its minimum Go language version / the collection of packages in this exercise. |
| `gofmt` | Go's standard source formatter. It uses tabs for indentation. |

See the [Go Two Sum guide](../Go/two-sum/README.md) for the runnable programs. The [Go language specification](https://go.dev/ref/spec) describes the precise map lookup and `range` behavior.

## 9. System design and URL shortening

| Term | Plain-language meaning and example |
| --- | --- |
| System design | Choosing a system's responsibilities, data, components, request flows, and trade-offs to meet stated requirements. |
| Scope | The boundary of what a version handles. This version creates and redirects links. |
| Functional requirement | An observable capability: submitting a valid destination produces a usable short link. |
| Non-goal | A feature deliberately excluded from this version, such as click reporting. |
| Non-functional requirement | A required quality such as speed, availability, or durability. |
| URL / destination URL | A web address / the original address a short link leads to. |
| Absolute URL | An address that includes its scheme and host, such as `https://shop.example.com/products/42`. |
| Scheme / host / path | The protocol (`https`), server name (`shop.example.com`), and resource path (`/products/42`) in a URL. |
| Query string / fragment | Parameters after `?` / a part after `#` used by the browser. Preserve both in the destination; the fragment is not sent in the HTTP request to the destination server. |
| Short code / mapping | The identifier in the short URL / its association with a stored destination. |
| Client / server | A component sending a request / one receiving it and returning a response. |
| Endpoint | An operation at an address, such as `POST /links`. |
| HTTP | The request and response protocol used in these web interactions. |
| POST / GET | The HTTP methods used here to create a link / open an existing short link. |
| Request / response | Information sent to a service / the result it sends back. |
| JSON | A text format for structured data, such as `{"url":"https://example.com"}`. |
| Status code | A number describing an HTTP response's outcome. This design uses `201` for creation, `302` for redirect, `400` for invalid input, `404` for a missing mapping, and `503` for unavailable service work. |
| Redirect / Location | A response directing a browser to another address / the response header carrying that address. |
| Header | A named piece of metadata in a request or response. `Location` is one response header. |
| Latency | Elapsed time for an operation, measured between specified start and end points. |
| Percentile / p95 | A position in a distribution / the duration at approximately the 95% point of sorted latency observations. It is not the average or a maximum. |
| Workload / throughput | The type and amount of work supplied / the amount completed per unit of time, such as requests per second. |
| Availability | How reliably the service can perform the required work when requested. Fast successes alone do not establish it. |
| Persistence / durability | Data remaining available across sessions or process restarts / committed data surviving failures covered by the storage system's guarantees. |
| Database / row | A data storage and retrieval system / one stored record. |
| Primary key / unique constraint | A non-null unique identifier for a row / a database rule preventing duplicate values in the constrained fields. |
| Database index | A structure that helps the database locate records. This meaning differs from an array position. |
| Commit | Complete a database write or transaction under its configured guarantees. It is separate from committing a file change in Git. |
| Code collision | Two attempts generate the same short code. Database uniqueness enforcement prevents replacing an existing mapping. |
| Retry / bounded retry | Attempt work again / limit the number of attempts so a failure cannot cause endless work. |
| Idempotency | Repeating the same logical operation has the intended effect only once. For link creation, an agreed request identifier can allow a retry to reuse its previous result. |
| Cache / stale data | A stored copy used for faster reuse / a copy that no longer reflects the source's current state. |
| Cache-Control: no-store | An HTTP instruction telling caches not to store the response. A `302` alone does not prohibit caching. |
| Failure path | The sequence of actions when a dependency or operation fails, such as returning a service error after a database timeout. |
| Markdown / Mermaid | A text format for documents / diagram syntax inside a Markdown code block that GitHub can render. |

See [HTTP Semantics](https://www.rfc-editor.org/rfc/rfc9110.html), [MDN Cache-Control](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Cache-Control), and [PostgreSQL constraints](https://www.postgresql.org/docs/current/ddl-constraints.html) for the protocol and database rules. The [URL shortener guide](url-shortener.md) explains the exercise's chosen policies and assumptions.
