import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Two Sum examples that can be run directly from main(). */
public class TwoSum {

    /** Expected O(n) time and O(n) extra space. */
    public static int[] twoSum(int[] nums, int target) {
        // Each entry stores an earlier number and its index.
        Map<Integer, Integer> seen = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];

            // Check before storing so the same element cannot match itself.
            if (seen.containsKey(complement)) {
                return new int[] {seen.get(complement), i};
            }

            seen.put(nums[i], i);
        }

        throw new IllegalArgumentException("No valid pair exists");
    }

    /** O(n^2) worst-case time and O(1) extra space. */
    public static int[] twoSumBruteForce(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) {
                    return new int[] {i, j};
                }
            }
        }

        throw new IllegalArgumentException("No valid pair exists");
    }

    public static void main(String[] args) {
        int[][] examples = {
            {2, 7, 11, 15},
            {3, 2, 4},
            {3, 3}
        };
        int[] targets = {9, 6, 6};

        for (int i = 0; i < examples.length; i++) {
            int[] answer = twoSum(examples[i], targets[i]);
            System.out.printf(
                "nums=%s, target=%d -> %s%n",
                Arrays.toString(examples[i]),
                targets[i],
                Arrays.toString(answer)
            );
        }
    }
}
