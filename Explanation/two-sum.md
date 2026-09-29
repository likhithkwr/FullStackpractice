# Two Sum

A step-by-step guide to finding two numbers that add up to a target, with Java and Python code, worked examples, diagrams, and a realistic industry scenario.

Read the essential terms below first. The [terminology guide](terminology.md) also explains Java syntax, execution tools, performance terms, and the words used in the industry scenario.

## 1. What the problem asks

Given an integer array `nums` and an integer `target`, find two **different positions** `i` and `j` such that:

```text
nums[i] + nums[j] = target
i != j
```

Return `[i, j]`. These are zero-based **indices**, not the values themselves. Either index order is accepted.

For example:

```text
Index:   0   1    2    3
Value:   2   7   11   15
         └───┘
         2 + 7 = 9

Answer: [0, 1]
```

The input guarantees exactly one valid pair of indices. Equal values can be used if they belong to different positions: `[3, 3]` contains two separate elements.

The constraints mean `2 <= n <= 10^4`, and values and the target lie between `-10^9` and `10^9`. The array is not guaranteed to be sorted.

Here, `n` means the number of elements. `10^4` means 10,000, and `10^9` means 1,000,000,000. These limits describe which inputs the problem allows.

### Essential terms

| Term | Plain-language meaning | Example |
| --- | --- | --- |
| Integer | A whole number, including zero and negative numbers. | `-3`, `0`, `7`. |
| Array | An ordered collection of elements. | `[2, 7, 11, 15]`. |
| Element | One item in the array. | `7` is an element. |
| Index | An element's position. Java and Python start counting positions at zero. | `7` is at index `1`. |
| Value | The number stored at a position. | `nums[1]` has value `7`. |
| Target | The total that the selected values must reach. | Target `9` means the pair must sum to `9`. |
| Pair | Two selected elements. Their positions must differ. | Values `2` and `7`, at indices `0` and `1`. |
| Complement | The one value needed to complete the current number's pair. | If the current value is `7`, `9 - 7 = 2`. |
| Algorithm | A precise sequence of steps for solving a problem. | Compute the complement, look it up, then return or store. |
| Hash map | A structure that associates a key with a stored value and supports fast lookup. | In this solution, number `2` maps to index `0`. |

**An index and a value are different things.** `nums[1] = 7` says that position `1` contains the value `7`. The answer reports the positions.

## 2. Start with the straightforward solution

Try every distinct pair. For each index `i`, try indices `j` starting at `i + 1`.

```python
def two_sum_brute_force(nums: list[int], target: int) -> list[int]:
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i, j]
    raise ValueError("No valid pair exists")
```

Starting `j` at `i + 1` prevents using one position twice and avoids checking both `(i, j)` and `(j, i)`.

For `[3, 2, 4]`, target `6`:

```text
Indices (0, 1): 3 + 2 = 5 → no
Indices (0, 2): 3 + 4 = 7 → no
Indices (1, 2): 2 + 4 = 6 → return [1, 2]
```

There are at most `n(n - 1) / 2` pairs. At `n = 10,000`, that is **49,995,000** pair checks. The worst-case time is `O(n^2)`; extra space is `O(1)`.

This is a useful baseline to explain first in an interview. The repeated search for a partner is what we can improve.

## 3. The key idea: look up the complement

If the current number is `x`, its required partner is fixed:

```text
x + partner = target
partner = target - x
```

This partner is called the **complement**.

For target `9`, when `x = 7`, the complement is `2`. We do not need to try every other number; we need to answer: **Have we already seen a 2?**

A hash map stores numbers we have already visited and their positions:

```text
seen = {value: index}

{2: 0} means "the value 2 was found at index 0."
```

In Python, a dictionary provides this map. Lookup and insertion take expected `O(1)` time. A map is useful here because we need both the matching value and its index.

### Think of the map as a notebook of earlier numbers

For every number you visit, remember its value and its position. When the next number arrives, calculate the partner it needs and look for that partner in your earlier records. The hash map makes that lookup fast.

For `[3, 2, 4]`, target `6`, just before reading `4`:

```text
Earlier records:

Key: number seen       Stored value: its index
        3          →               0
        2          →               1

Current index: 2
Current value: 4
Required complement: 6 - 4 = 2

Look up key 2 → retrieve index 1 → return [1, 2]
```

The map's **key** is an array value. The map's **stored value** is that number's index. This is why `seen.put(nums[i], i)` in Java puts the number first and its position second.

## 4. The algorithm

1. Create an empty map called `seen`.
2. Visit the array from left to right.
3. Compute `complement = target - x`.
4. If the complement is in `seen`, return its stored index and the current index.
5. Otherwise, store `seen[x] = i` and continue.

**Check first, store second.**

```mermaid
flowchart TD
    A["Start with an empty seen map"] --> B{"Another element?"}
    B -- Yes --> C["Read index i and value x"]
    C --> D["complement = target - x"]
    D --> E{"Is complement in seen?"}
    E -- Yes --> F["Return stored index and i"]
    E -- No --> G["Store seen[x] = i"]
    G --> B
    B -- No --> H["No solution: outside the given guarantee"]
```

## 5. Implementations

### Python

This class can be used directly for the supplied interview problem:

```python
class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        seen: dict[int, int] = {}

        for i, x in enumerate(nums):
            complement = target - x

            if complement in seen:
                return [seen[complement], i]

            seen[x] = i

        raise ValueError("No valid pair exists")
```

#### Python syntax

| Code | Meaning |
| --- | --- |
| `seen = {}` | Remember earlier values and their indices. |
| `enumerate(nums)` | Read each index `i` together with its value `x`. |
| `target - x` | Compute the only value that can complete the pair. |
| `complement in seen` | Check whether that value occurred earlier. |
| `[seen[complement], i]` | Return the earlier position and current position. |
| `seen[x] = i` | Make this number available to later elements. |

The final exception is defensive behavior for an input without a solution. Under the problem's guarantee, the function returns before reaching it.

### Java

The runnable [TwoSum.java](../Java/two-sum/TwoSum.java) uses the same complement lookup. Its core method is:

```java
public static int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> seen = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];

        if (seen.containsKey(complement)) {
            return new int[] {seen.get(complement), i};
        }

        seen.put(nums[i], i);
    }

    throw new IllegalArgumentException("No valid pair exists");
}
```

The file imports `java.util.Map` and `java.util.HashMap`. The map stores each earlier value as a key and its index as the value. Java generics use the wrapper type `Integer` rather than primitive `int`.

`containsKey` performs the membership check, `get` retrieves the earlier index, and `put` stores the current number. `new int[] {...}` creates the array returned to the caller. The local method is `static` so the example `main()` can call it directly.

Checking before `put` gives the same protection against reusing an index as the Python version. The time and space complexity and all walkthroughs below apply to both implementations. See the [Java guide](../Java/two-sum/README.md) for editor instructions and the LeetCode method signature.

#### Read the Java method one line at a time

| Code | What it does |
| --- | --- |
| `public static int[] twoSum(int[] nums, int target)` | Declares a method named `twoSum`. It receives an integer array and a target, and returns an integer array. `public` makes it accessible to callers; `static` allows a call such as `TwoSum.twoSum(...)` without creating an object. |
| `Map<Integer, Integer> seen = new HashMap<>();` | Creates an empty map. The first `Integer` describes its keys; the second describes its stored values. `seen` is the variable name. |
| `int i = 0` | Start with the first array position. |
| `i < nums.length` | Continue while `i` is a valid position. `length` is the number of elements. |
| `i++` | Increase the position by one after each iteration. |
| `int complement = target - nums[i];` | Calculate the required partner for the current value. |
| `seen.containsKey(complement)` | Ask whether that partner occurred at an earlier position. The result is a boolean: `true` or `false`. |
| `seen.get(complement)` | Retrieve the earlier partner's index. |
| `return new int[] {seen.get(complement), i};` | Create the two-index answer and immediately end the method. |
| `seen.put(nums[i], i);` | Store the current number and position so later elements can find it. |
| `throw new IllegalArgumentException(...)` | Signal that the input has no valid pair. The problem's guarantee means this does not occur for valid inputs. |

An **iteration** is one repetition of the loop. A **method** is a named piece of code you call to do a task. Its **parameters** are the named inputs in its declaration; its **arguments** are the actual values supplied in a call.

For example, in `TwoSum.twoSum(new int[] {2, 7}, 9)`, the arguments are the array containing `2` and `7`, and the target `9`.

## 6. Walk through every supplied example

The tables show `seen` **before** checking the current number.

### Example 1: `[2, 7, 11, 15]`, target `9`

| Index | Value | Complement | Map before lookup | Action |
| --- | --- | --- | --- | --- |
| 0 | 2 | 7 | `{}` | No 7; store `{2: 0}`. |
| 1 | 7 | 2 | `{2: 0}` | Found 2 at index 0; return `[0, 1]`. |

The algorithm stops immediately. It never needs to visit `11` or `15`.

### Example 2: `[3, 2, 4]`, target `6`

| Index | Value | Complement | Map before lookup | Action |
| --- | --- | --- | --- | --- |
| 0 | 3 | 3 | `{}` | No earlier 3; store `{3: 0}`. |
| 1 | 2 | 4 | `{3: 0}` | No 4; store `{3: 0, 2: 1}`. |
| 2 | 4 | 2 | `{3: 0, 2: 1}` | Found 2 at index 1; return `[1, 2]`. |

Although `3 + 3 = 6`, there is only one `3` in this array. Using index `0` twice would be invalid. The map is empty when index `0` is checked, so the algorithm correctly continues.

#### Diagram: the array scan talks to the map

This sequence diagram reads from top to bottom. Each arrow represents one lookup, response, or insertion.

```mermaid
sequenceDiagram
    participant S as Array scan
    participant M as Map of earlier values
    S->>M: Index 0, value 3: look for complement 3
    M-->>S: Not found
    S->>M: Store number 3 with index 0
    S->>M: Index 1, value 2: look for complement 4
    M-->>S: Not found
    S->>M: Store number 2 with index 1
    S->>M: Index 2, value 4: look for complement 2
    M-->>S: Found number 2 at index 1
    Note over S,M: Return indices 1 and 2
```

### Example 3: `[3, 3]`, target `6`

| Index | Value | Complement | Map before lookup | Action |
| --- | --- | --- | --- | --- |
| 0 | 3 | 3 | `{}` | No earlier 3; store `{3: 0}`. |
| 1 | 3 | 3 | `{3: 0}` | Found 3 at index 0; return `[0, 1]`. |

```text
First 3, index 0                    Second 3, index 1
Check empty map → no match          Need 6 - 3 = 3
Store 3 → 0                        Map contains 3 → 0
                                   Return [0, 1]
```

The values are equal, but the indices are different. Returning `[3, 3]` would be wrong because the requested answer contains indices.

The map stores only earlier positions. At the second `3`, the stored `3` refers to index `0`, and the current element is at index `1`.

## 7. Why checking before storing matters

Here is an incorrect order:

```python
seen[x] = i  # Too early!
if target - x in seen:
    return [seen[target - x], i]
```

For `[3, 2, 4]`, target `6`, this inserts `3: 0` and then finds that same entry. It incorrectly returns `[0, 0]`, even though the valid answer is `[1, 2]`.

In the correct algorithm, the map only contains positions strictly before `i`. Therefore any returned index `j` satisfies `j < i`, which guarantees two different elements.

## 8. Why the algorithm is correct

A **proof of correctness** explains why the method must produce a valid answer for every allowed input. A useful tool is a **loop invariant**: a statement that stays true at the same point in every repetition of the loop.

**Invariant:** Before processing index `i`, the map contains every distinct value from earlier positions, mapped to an earlier index where that value appeared.

- If the complement is present, its value is `target - nums[i]`, so the returned values sum to `target`.
- Its stored position is earlier than `i`, so the same element is never reused.
- If the solution is at indices `a < b`, then when we reach `b`, the value at `a` is already in the map. The complement lookup finds it.
- Storing the current value after an unsuccessful lookup preserves the invariant for the next iteration.

This proves both that a returned answer is valid and that an existing answer will be found.

## 9. Time and space complexity

**Complexity** describes how the work or memory grows as the input gets larger. It is not an exact duration in seconds or an exact count of bytes.

- **Time complexity:** how much work the algorithm performs as `n` grows.
- **Space complexity:** how much extra memory it needs as `n` grows.
- **Big O notation:** a way to describe an upper bound on that growth while ignoring constant factors.
- **`O(1)`:** constant growth; the cost does not grow with the number of array elements in the relevant operation.
- **`O(n)`:** linear growth; scanning twice as many elements involves roughly twice as much scanning work.
- **`O(n^2)`:** quadratic growth; trying all pairs in an array twice as large involves roughly four times as many pair checks in the worst case.

These comparisons describe growth, not guaranteed wall-clock timing. An early match lets either solution stop sooner.

| Approach | Time | Extra space | Notes |
| --- | --- | --- | --- |
| Try all pairs | `O(n^2)` worst case | `O(1)` | Simple; no map required. |
| One pass with a hash map | `O(n)` expected | `O(n)` worst case | Keeps original indices and avoids sorting. |
| Sort `(value, original_index)` pairs, then use two pointers | `O(n log n)` | `O(n)` for the pairs | Must preserve original indices. |

The hash map solution visits each element at most once. Each visit performs expected constant-time map operations. It may store a number of entries proportional to `n`.

The word **expected** matters: fast hash map lookup assumes keys are distributed suitably through its internal storage. Different keys can end up in the same storage location, called a **hash collision**. Maps handle collisions so the keys remain distinguishable, but the cost depends on the implementation. Some hash table implementations can degrade to quadratic total time under poor collision handling. The usual interview claim for this solution is **expected `O(n)` time and `O(n)` extra space**.

An insertion may also occasionally resize the map and rearrange existing entries. That occasional work is spread over many insertions; averaging such cost over a sequence is called **amortized analysis**. This differs from the expectation used for hash distribution.

If the input is already sorted, two pointers can solve the problem in `O(n)` time and `O(1)` extra space. Two pointers are not valid on the arbitrary unsorted array in this prompt without first sorting it.

## 10. Where this pattern is useful

Use this approach when:

- You need two distinct records whose integer quantities sum to an exact target.
- You need the actual positions or record identifiers of the match.
- The input is unsorted and you can use extra memory.
- You need to detect a match while processing records in order.

The broader pattern is: **compute what would complete the current item, then look it up among earlier items.** Hash maps also support related interview problems such as counting pairs and finding a subarray with a target sum. Those problems need different stored information: for example, frequencies for counting pairs and prefix sums for subarrays.

Two Sum answers an exact two-item matching question. It does not directly solve choosing any number of items, finding the closest total, or maximizing a total below a capacity; those have different requirements.

## 11. Industry scenario: fulfillment pack selection

**Hypothetical scenario:** A warehouse keeps sealed packs of components. For one order, a workstation must choose exactly two available packs of the same SKU that together contain **600 components**, without opening either pack.

**SKU** means stock keeping unit: an identifier for a type of product. An **inventory snapshot** is the set of records observed when choosing candidates. A **pack ID** identifies one particular physical pack, even if its array position changes.

After filtering by SKU, warehouse, and availability, it has this inventory snapshot:

| Array index | Pack ID | Component count |
| --- | --- | --- |
| 0 | PK-A | 300 |
| 1 | PK-B | 200 |
| 2 | PK-C | 400 |

Run Two Sum on `[300, 200, 400]` with target `600`:

1. Pack PK-A needs another 300-component pack. None has been seen.
2. Pack PK-B needs a 400-component pack. None has been seen.
3. Pack PK-C needs a 200-component pack. PK-B matches.
4. Return indices `[1, 2]`, then resolve them to stable pack IDs **PK-B and PK-C**.

```mermaid
flowchart TD
    A["Order: 600 components in exactly 2 packs"] --> B["Filter compatible available packs"]
    B --> C["Scan counts with a hash map"]
    C --> D["Candidate: PK-B and PK-C"]
    D --> E["Atomically reserve both packs if available"]
    E --> F["Send reservation to fulfillment"]
```

### Production considerations

- **Stable identity:** Array positions are local to the snapshot. Return pack IDs to other services.
- **Compatibility:** Filter by SKU and other required attributes before matching counts. Equal counts alone do not establish compatibility.
- **Concurrent requests:** Requests can overlap in time. Another worker might reserve a pack after the snapshot was read. Recheck availability and reserve both packs together, with availability checks and suitable locking or conditional writes. An **atomic** reservation succeeds for both packs or neither; a database transaction can group the changes. If the reservation fails, refresh candidates and retry according to the service's policy.
- **No or multiple matches:** Real input may not satisfy the interview's uniqueness guarantee. Define the no-match response and any selection policy, such as preferring packs that expire sooner. The basic algorithm returns the first pair encountered and does not optimize such a policy.
- **Changed requirements:** If an order can use three or more packs, or partial packs, this algorithm no longer solves the full selection problem.

The map solves the candidate search. Filtering, identity, selection policy, and reservation make it usable inside the larger workflow.

## 12. Edge cases and common mistakes

| Case | Example | Expected result or lesson |
| --- | --- | --- |
| Equal values at different positions | `[3, 3]`, target `6` | `[0, 1]` is valid. |
| A value is half the target but appears once | `[3, 2, 4]`, target `6` | Use 2 and 4, not the same 3 twice. |
| Zero values | `[0, 4, 3, 0]`, target `0` | `[0, 3]`. |
| Negative values | `[-3, 4, 3, 90]`, target `0` | `[0, 2]`. |
| Smallest allowed array | `[2, 7]`, target `9` | `[0, 1]`. |
| Match is stored at index 0 | `[2, 7]`, target `9` | Use membership, not the truthiness of the stored index. |

Avoid these mistakes:

- Returning values instead of indices.
- Inserting the current number before looking up its complement.
- Using `if seen.get(complement):`; a matching index of `0` is false in Python. Use `if complement in seen:`.
- Rejecting duplicates. The rule forbids reusing a position, not using equal values.
- Sorting only the values and then returning positions from the sorted array.
- Applying two pointers directly to unsorted input.

## 13. Interview-ready explanation

> I can try every pair in quadratic time. To improve that, for each number x I compute target minus x and check a hash map of previously visited values. If its complement is present, I return the earlier index and the current index. Otherwise, I store the current value and index. Checking before storing prevents using the same element twice. This takes expected linear time and linear extra space.

### Common follow-up questions

**Why a map instead of a set?** A set can tell us whether a value exists, but the map also retrieves its index.

**How do duplicates work?** The earlier copy is stored before the later copy is processed. If the value is its own complement, the later copy finds the earlier index.

**What if no solution exists?** Follow the agreed API contract, such as raising an exception or returning an optional result. This prompt guarantees a solution, so that branch is not reached for valid input.

**What if more than one pair exists?** This algorithm returns the first valid pair it encounters. Returning every pair requires a modified algorithm that retains enough indices for duplicates; the output itself can be quadratic in size.

**What if the array is sorted?** Move a left and a right pointer inward: increase the left pointer when the sum is too small and decrease the right pointer when it is too large.

**Can I preserve the input?** Yes. The hash map approach does not modify `nums`.

**Is linear time optimal?** For arbitrary unsorted input, there are cases where the relevant pair is only established after reading the final element. Worst-case inspection therefore needs linear work. The expected linear-time hash map approach meets that bound under expected constant-time map operations.

## 14. Run the reference example code

Both the [Python solution](../Python/two-sum/solution.py) and [Java solution](../Java/two-sum/TwoSum.java) contain the optimized solution, the brute-force baseline, and the three supplied examples.

From the repository root:

```sh
python3 Python/two-sum/solution.py
```

For Java, with JDK 11 or newer:

```sh
java Java/two-sum/TwoSum.java
```

The [Java guide](../Java/two-sum/README.md) also describes compiling and running its checks.

### Why the Java example has `main()`

Defining a method describes what it should do. Calling the method makes it execute. The `main()` method is the **entry point** for this standalone Java program: it supplies each example array and target, calls `twoSum`, and prints the returned indices.

For the usual compiled workflow, `javac` compiles Java source into `.class` files containing **bytecode**, and the Java Virtual Machine (**JVM**) runs that bytecode. The Java Development Kit (**JDK**) includes the tools used for this workflow. The JDK 11+ single-file command above compiles the source as part of launching it.

## 15. Terminology reference

See [terminology.md](terminology.md) for the complete glossary, including `int`, `Integer`, `Map`, `HashMap`, `public`, `static`, `void`, `main`, `new`, generics, lookup, insertion, runtime, compilation, tests, and the production terms in this guide.

To recall the approach quickly: **complement → lookup → return or store**.

## 16. Write your own solution locally

Use the separate practice files to implement the algorithm in your editor:

| Language | Your practice file | Completed reference |
| --- | --- | --- |
| Java | [TwoSumPractice.java](../Java/two-sum/TwoSumPractice.java) | [TwoSum.java](../Java/two-sum/TwoSum.java) |
| Python | [practice.py](../Python/two-sum/practice.py) | [solution.py](../Python/two-sum/solution.py) |

1. Read the problem statement in the practice file and implement only the `twoSum()` method marked `TODO`.
2. Run that file using the editor's Run button or the commands below.
3. Use the feedback to fix your implementation. The checker calls your method and accepts either valid index order.
4. Compare the completed approach with the reference and explain its time and space complexity.

From the repository root:

```sh
java Java/two-sum/TwoSumPractice.java
python3 Python/two-sum/practice.py
```

The initial run shows eight `TODO` cases because the method is unfinished. After implementing it, correct answers show `PASS`; invalid answers show `FAIL`, and unexpected exceptions show `ERROR`. Exit code `0` means all checks passed, `1` means a failure or error, and `2` means there is unfinished code without a reported failure.

See the [Java practice guide](../Java/two-sum/README.md) or [Python practice guide](../Python/two-sum/README.md) for editor instructions and adding custom cases. Each learner can implement the starter on their own Git branch.
