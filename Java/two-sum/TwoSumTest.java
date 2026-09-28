import java.util.Arrays;

/** Dependency-free checks; failures throw AssertionError. */
public class TwoSumTest {

    private static void check(int[] nums, int target, int first, int second) {
        int[] original = nums.clone();
        checkAnswer(TwoSum.twoSum(nums, target), first, second);
        checkAnswer(TwoSum.twoSumBruteForce(nums, target), first, second);

        if (!Arrays.equals(nums, original)) {
            throw new AssertionError("A solution changed the input array");
        }
    }

    private static void checkAnswer(int[] actual, int first, int second) {
        boolean forward = actual.length == 2
            && actual[0] == first && actual[1] == second;
        boolean reverse = actual.length == 2
            && actual[0] == second && actual[1] == first;

        if (!forward && !reverse) {
            throw new AssertionError("Unexpected indices: " + Arrays.toString(actual));
        }
    }

    private static void expectNoSolution(Runnable operation) {
        try {
            operation.run();
        } catch (IllegalArgumentException expected) {
            return;
        }
        throw new AssertionError("Expected an exception for an input without a pair");
    }

    public static void main(String[] args) {
        check(new int[] {2, 7, 11, 15}, 9, 0, 1);
        check(new int[] {3, 2, 4}, 6, 1, 2);
        check(new int[] {3, 3}, 6, 0, 1);
        check(new int[] {0, 4, 3, 0}, 0, 0, 3);
        check(new int[] {-3, 4, 3, 90}, 0, 0, 2);
        check(new int[] {2, 7}, 9, 0, 1);
        check(new int[] {-1_000_000_000, 1_000_000_000}, 0, 0, 1);
        check(new int[] {5, 8, 1, 12}, 20, 1, 3);

        expectNoSolution(() -> TwoSum.twoSum(new int[] {1, 2}, 10));
        expectNoSolution(() -> TwoSum.twoSumBruteForce(new int[] {1, 2}, 10));

        System.out.println("Passed 8 input cases for both approaches and 2 no-solution checks.");
    }
}
