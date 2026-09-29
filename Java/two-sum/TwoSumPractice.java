import java.util.Arrays;

/**
 * Your local Two Sum practice file.
 * Find two different indices whose values add up to target.
 * Exactly one valid pair exists; either index order is accepted.
 * Implement twoSum(), then run main() to check your answers.
 */
public class TwoSumPractice {

    public static int[] twoSum(int[] nums, int target) {
        // TODO: Replace the placeholder below with your solution. Return the two indices.
        throw new UnsupportedOperationException("TODO: implement twoSum()");
    }

    // The checker validates the indices against the original input.
    private static boolean isValidAnswer(int[] nums, int target, int[] answer) {
        if (answer == null || answer.length != 2) {
            return false;
        }

        int first = answer[0];
        int second = answer[1];
        return first >= 0 && first < nums.length
            && second >= 0 && second < nums.length
            && first != second
            && nums[first] + nums[second] == target;
    }

    public static void main(String[] args) {
        int[][] examples = {
            {2, 7, 11, 15},
            {3, 2, 4},
            {3, 3},
            {0, 4, 3, 0},
            {-3, 4, 3, 90},
            {2, 7},
            {-1_000_000_000, 1_000_000_000},
            {5, 8, 1, 12}
        };
        int[] targets = {9, 6, 6, 0, 0, 9, 0, 20};
        int[][] expected = {
            {0, 1}, {1, 2}, {0, 1}, {0, 3},
            {0, 2}, {0, 1}, {0, 1}, {1, 3}
        };

        int passed = 0;
        int failed = 0;
        int todo = 0;
        System.out.println("Two Sum practice: implement twoSum(), then run this file.");

        for (int i = 0; i < examples.length; i++) {
            String input = "nums=" + Arrays.toString(examples[i]) + ", target=" + targets[i];

            try {
                int[] answer = twoSum(examples[i].clone(), targets[i]);
                if (isValidAnswer(examples[i], targets[i], answer)) {
                    passed++;
                    System.out.printf("[PASS] %s -> %s%n", input, Arrays.toString(answer));
                } else {
                    failed++;
                    System.out.printf(
                        "[FAIL] %s -> %s; expected %s (either order)%n",
                        input, Arrays.toString(answer), Arrays.toString(expected[i])
                    );
                }
            } catch (UnsupportedOperationException unfinished) {
                todo++;
                System.out.printf(
                    "[TODO] %s; expected %s (either order)%n",
                    input, Arrays.toString(expected[i])
                );
            } catch (RuntimeException error) {
                failed++;
                System.out.printf(
                    "[ERROR] %s -> %s: %s%n",
                    input, error.getClass().getSimpleName(), error.getMessage()
                );
            }
        }

        System.out.printf("Summary: %d PASS, %d FAIL, %d TODO%n", passed, failed, todo);
        if (todo > 0) {
            System.out.println("Replace the TODO in twoSum() with your implementation.");
        }
        if (failed > 0) {
            System.exit(1);
        }
        if (todo > 0) {
            System.exit(2);
        }
    }
}
