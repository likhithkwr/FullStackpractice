"""Your local Two Sum practice file: implement Solution.twoSum and run it.

Return two different indices whose values add up to target.
Exactly one valid pair exists, and either index order is accepted.
"""


class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        # TODO: Replace the placeholder below with your solution. Return the two indices.
        seen = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in seen:
                return [seen[complement], i]
            seen[nums[i]] = i
        raise ValueError("No value pair")


def is_valid_answer(nums: list[int], target: int, answer: object) -> bool:
    if not isinstance(answer, (list, tuple)) or len(answer) != 2:
        return False

    first, second = answer
    if type(first) is not int or type(second) is not int:
        return False

    return (
        0 <= first < len(nums)
        and 0 <= second < len(nums)
        and first != second
        and nums[first] + nums[second] == target
    )


def run_examples() -> int:
    examples = [
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
        ([0, 4, 3, 0], 0, [0, 3]),
        ([-3, 4, 3, 90], 0, [0, 2]),
        ([2, 7], 9, [0, 1]),
        ([-1_000_000_000, 1_000_000_000], 0, [0, 1]),
        ([5, 8, 1, 12], 20, [1, 3]),
    ]
    solver = Solution()
    passed = failed = todo = 0
    print("Two Sum practice: implement twoSum(), then run this file.")

    for nums, target, expected in examples:
        description = f"nums={nums}, target={target}"
        try:
            answer = solver.twoSum(nums.copy(), target)
        except NotImplementedError:
            todo += 1
            print(f"[TODO] {description}; expected {expected} (either order)")
        except Exception as error:
            failed += 1
            print(f"[ERROR] {description} -> {type(error).__name__}: {error}")
        else:
            if is_valid_answer(nums, target, answer):
                passed += 1
                print(f"[PASS] {description} -> {answer}")
            else:
                failed += 1
                print(f"[FAIL] {description} -> {answer}; expected {expected} (either order)")

    print(f"Summary: {passed} PASS, {failed} FAIL, {todo} TODO")
    if todo:
        print("Replace the TODO in twoSum() with your implementation.")
    if failed:
        return 1
    if todo:
        return 2
    return 0


if __name__ == "__main__":
    raise SystemExit(run_examples())
