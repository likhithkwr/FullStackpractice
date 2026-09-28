# Two Sum — Java

Find the indices of two different array elements whose values add up to a target. Read the [shared explanation](../../Explanation/two-sum.md) for walkthroughs, diagrams, complexity, and an industry scenario, and the [terminology guide](../../Explanation/terminology.md) for plain-language definitions.

## Files

- [TwoSum.java](TwoSum.java): the hash map solution, a brute-force baseline, and a runnable `main()` with the three supplied examples.
- [TwoSumTest.java](TwoSumTest.java): checks for both approaches, including duplicates, zero, negative values, the input bounds, and defensive no-solution behavior.

## Run the examples

With JDK 11 or newer, run this single-file command from the repository root:

```sh
java Java/two-sum/TwoSum.java
```

Expected output:

```text
nums=[2, 7, 11, 15], target=9 -> [0, 1]
nums=[3, 2, 4], target=6 -> [1, 2]
nums=[3, 3], target=6 -> [0, 1]
```

Edit the arrays and targets inside `main()` to try another input.

## Compile and run the checks

From the repository root:

```sh
mkdir -p Java/two-sum/build
javac -d Java/two-sum/build Java/two-sum/TwoSum.java Java/two-sum/TwoSumTest.java
java -cp Java/two-sum/build TwoSum
java -cp Java/two-sum/build TwoSumTest
```

No external packages or test framework are required. Generated files stay in the ignored `build` directory.

## Run in an editor

- **VS Code:** use the Extension Pack for Java, select an installed JDK, open `TwoSum.java`, and click **Run** above `main()`.
- **IntelliJ IDEA:** set the project SDK to an installed JDK, open `TwoSum.java`, and use the Run icon beside `main()`.
- Run `TwoSumTest.main()` in the same way to check both approaches.

## Java syntax to understand

| Code | Meaning |
| --- | --- |
| `int[] nums` | The input array of integers. |
| `Map<Integer, Integer>` | A map from a number to its earlier index. Generic types use `Integer`, the object wrapper for `int`. |
| `new HashMap<>()` | The hash map implementation used for expected constant-time lookup and insertion. |
| `seen.containsKey(complement)` | Check whether the required partner occurred earlier. |
| `seen.get(complement)` | Retrieve its stored index. |
| `seen.put(nums[i], i)` | Remember the current value for later elements. |
| `new int[] {j, i}` | Return the two indices in an integer array. |
| `static` | The examples can call the method without creating a `TwoSum` object. |

The hash map method takes expected `O(n)` time and `O(n)` extra space. The brute-force method takes `O(n^2)` worst-case time and `O(1)` extra space. Both leave the input unchanged and check distinct positions.

For the supplied bounds, the sums and complements fit in Java's `int` type.

## Using the solution on LeetCode

Copy the `HashMap` and `Map` imports and the `twoSum` method into LeetCode's `class Solution`. Use the method signature `public int[] twoSum(int[] nums, int target)`. The local `main()` and test runner provide example inputs when practicing in an editor.
