# Two Sum — Python

Read the [explanation](../../Explanation/two-sum.md) and [terminology guide](../../Explanation/terminology.md) for the problem, examples, diagrams, and technical words.

## Files

| File | Purpose |
| --- | --- |
| [practice.py](practice.py) | Your local practice file: write `Solution.twoSum` and run the built-in checks. |
| [solution.py](solution.py) | Completed reference implementations and examples to compare with your approach. |

## Practice locally

1. Open `practice.py` and replace the `NotImplementedError` placeholder in `Solution.twoSum` with your implementation.
2. Run it with VS Code's **Run Python File** button or from the repository root:

   ```sh
   python3 Python/two-sum/practice.py
   ```

3. Read the result for each case. The checker accepts either index order and requires two different valid indices whose original values sum to the target.
4. Compare your approach with the reference after working through the problem.

The first run reports eight `TODO` cases because the method is intentionally unfinished. Exit code `2` means unfinished, `1` means at least one answer failed or raised an error, and `0` means all checks passed.

The checks include the three supplied examples, zeros, negative numbers, two-element input, the input bounds, and a pair found at the end. Add custom cases to the `examples` list as `(nums, target, expected_indices)`.

Prerequisite: Python 3.9 or newer. No external packages are required.

To run the reference from the repository root:

```sh
python3 Python/two-sum/solution.py
```
