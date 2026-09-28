"""Two Sum: a hash map solution and a brute-force reference."""


class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        seen: dict[int, int] = {}

        for i, x in enumerate(nums):
            complement = target - x

            if complement in seen:
                return [seen[complement], i]

            # Store after looking up so the same position cannot match itself.
            seen[x] = i

        raise ValueError("No valid pair exists")


def two_sum_brute_force(nums: list[int], target: int) -> list[int]:
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i, j]

    raise ValueError("No valid pair exists")


if __name__ == "__main__":
    solver = Solution()
    examples = [
        ([2, 7, 11, 15], 9),
        ([3, 2, 4], 6),
        ([3, 3], 6),
    ]

    for nums, target in examples:
        answer = solver.twoSum(nums, target)
        print(f"nums={nums}, target={target} -> {answer}")
