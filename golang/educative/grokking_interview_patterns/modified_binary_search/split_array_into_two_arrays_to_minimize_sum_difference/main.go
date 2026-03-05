package main

import (
	"fmt"
	"sort"
	"strings"
)

// =============================================================================
// PROBLEM: Minimum Difference Between Sums of Two Equal-Size Groups
//
// Split nums (len = 2n) into two groups of exactly n elements each.
// Minimize |sum(groupA) - sum(groupB)|.
// =============================================================================

// =============================================================================
// KEY MATH INSIGHT:
//
//   S_A + S_B = totalSum       (they cover all elements)
//   S_A - S_B = difference     (what we want to minimize)
//
//   Solving: S_A = (totalSum + diff) / 2
//            S_B = (totalSum - diff) / 2
//
//   => Minimizing |S_A - S_B| is equivalent to finding S_A as close
//      to totalSum/2 as possible.
//
//   So we set target = floor(totalSum / 2) and try to get one group's
//   sum as close to target as possible.
// =============================================================================

// =============================================================================
// WHY MEET IN THE MIDDLE?
//
//   Brute force: each of 2n elements goes to group A or B => 2^(2n) combos.
//   For n=15 that's 1 billion+ — too slow.
//
//   Meet in the middle splits the array into two halves:
//     Left half:  2^n subsets
//     Right half: 2^n subsets
//   For n=15: ~32,000 each — totally fine.
//
//   We then combine one subset from each half to form group A (size n),
//   with the remaining elements automatically forming group B (size n).
// =============================================================================

// =============================================================================
// EQUAL SIZE CONSTRAINT:
//
//   We can't just pick any subset from left + any from right.
//   If we pick k elements from the left half,
//   we MUST pick exactly (n-k) from the right half,
//   so that group A has exactly n elements total.
//
//   That's why subsets are stored as (count, sum) pairs
//   and right subsets are grouped by count.
// =============================================================================

// =============================================================================
// WHY BINARY SEARCH AND HOW IT WORKS HERE:
//
//   For each left subset (count=k, sum=leftSum), we need a right subset
//   with exactly (halfLen-k) elements whose sum is as close to
//   `need = target - leftSum` as possible.
//
//   rightByCount[halfLen-k] is a SORTED slice of all such right sums.
//   Instead of scanning every element (O(m)), binary search finds the
//   closest value in O(log m).
//
//   We use upperBound to find `idx` = first position where cand[idx] > need:
//
//     index:  0   1   2  [idx-1] [idx]  ...
//     value:  ... ... ...  ≤need  >need  ...
//                           ↑       ↑
//                      left neighbor  right neighbor
//
//   The closest value must be one of these two neighbors — anything further
//   left is even smaller, anything further right is even larger.
//   We compute the miss (|neighbor - need|) for both and keep the minimum.
//
//   Visual example — need=8, cand=[2, 5, 7, 11, 14]:
//
//     idx = upperBound(cand, 8) = 3   (first value > 8 is cand[3]=11)
//
//     left  neighbor: cand[2] = 7,  miss = 8-7  = 1
//     right neighbor: cand[3] = 11, miss = 11-8 = 3
//
//     => best miss from this left subset = 1
// =============================================================================

// =============================================================================
// FINAL ANSWER FORMULA: minDiff*2 + parity
//
//   minDiff tracks how far S_A missed the target (totalSum/2).
//   The actual group difference is:
//     |S_A - S_B| = |S_A - (totalSum - S_A)| = |2*S_A - totalSum|
//                 = 2 * |S_A - totalSum/2|
//                 = 2 * minDiff  +  (totalSum % 2)
//
//   The parity term corrects for integer division flooring the target.
//   Example: totalSum=7, target=3 (floored). Best S_A=3.
//     actual diff = |3 - 4| = 1
//     formula     = 2*0 + 1 = 1 ✓
// =============================================================================

// floorDiv2 performs floor division by 2 for any integer (including negative).
// Go's built-in `/` truncates toward zero, which differs from floor for negatives.
// e.g. -7/2 in Go = -3 (truncated), but floor(-7/2) = -4.
func floorDiv2(a int64) int64 {
	q := a / 2
	r := a % 2
	if r != 0 && a < 0 {
		q-- // adjust truncation toward zero → floor
	}
	return q
}

// mod2NonNegative returns a % 2, always non-negative.
// Needed because Go's % can return negative values for negative inputs.
// e.g. -7 % 2 = -1 in Go, but we want 1 (the true remainder).
func mod2NonNegative(a int64) int64 {
	r := a % 2
	if r < 0 {
		r += 2
	}
	return r
}

func minimumDifference(nums []int) int {
	n := len(nums)
	halfLen := n / 2 // each group must have exactly halfLen elements

	// Compute totalSum to derive the target each group should aim for.
	var totalSum int64
	for _, x := range nums {
		totalSum += int64(x)
	}

	// target = floor(totalSum / 2)
	// One group's sum should be as close to this as possible.
	target := floorDiv2(totalSum)

	// parity = totalSum % 2 (0 if even, 1 if odd).
	// Used in the final answer formula: minDiff*2 + parity.
	parity := mod2NonNegative(totalSum)

	// Split the array into two halves for meet-in-the-middle.
	left := nums[:halfLen]
	right := nums[halfLen:]

	// pair holds (count of elements chosen, their sum) for a subset.
	type pair struct {
		c int   // how many elements selected
		s int64 // their sum
	}

	// --- Generate all 2^halfLen subsets of the left half ---
	// Start with the empty subset: 0 elements chosen, sum = 0.
	leftSubsets := []pair{{0, 0}}
	for _, num := range left {
		next := make([]pair, 0, len(leftSubsets)*2)
		for _, p := range leftSubsets {
			next = append(next, p)                               // don't pick num
			next = append(next, pair{p.c + 1, p.s + int64(num)}) // pick num
		}
		leftSubsets = next
	}

	// --- Generate all 2^halfLen subsets of the right half ---
	rightSubsets := []pair{{0, 0}}
	for _, num := range right {
		next := make([]pair, 0, len(rightSubsets)*2)
		for _, p := range rightSubsets {
			next = append(next, p)
			next = append(next, pair{p.c + 1, p.s + int64(num)})
		}
		rightSubsets = next
	}

	// --- Group right subsets by their element count ---
	// rightByCount[k] = sorted list of sums achievable by picking exactly k
	// elements from the right half.
	//
	// Why sort? So we can binary search for the sum closest to a target.
	// Why group by count? Because if left picks k elements, right must
	// pick exactly (halfLen - k) to keep both groups at size halfLen.
	rightByCount := make([][]int64, halfLen+1)
	for _, p := range rightSubsets {
		if p.c >= 0 && p.c <= halfLen {
			rightByCount[p.c] = append(rightByCount[p.c], p.s)
		}
	}
	for i := range rightByCount {
		sort.Slice(rightByCount[i], func(a, b int) bool {
			return rightByCount[i][a] < rightByCount[i][b]
		})
	}

	minDiff := int64(1 << 62) // initialize to a very large value

	// --- For each left subset, binary search for the best right complement ---
	for _, lp := range leftSubsets {
		// Right half must contribute exactly (halfLen - lp.c) elements
		// so that group A totals halfLen elements.
		rc := halfLen - lp.c
		if rc < 0 || rc > halfLen {
			continue
		}

		cand := rightByCount[rc] // sorted right sums with correct count
		if len(cand) == 0 {
			continue
		}

		// We want rightSum as close to (target - leftSum) as possible,
		// so that leftSum + rightSum ≈ target = totalSum/2.
		need := target - lp.s

		// upperBound returns first index where cand[i] > need.
		// So cand[idx-1] <= need  (left neighbor, could be exact or under)
		//    cand[idx]   >  need  (right neighbor, over)
		// We check both and keep the smaller miss.
		idx := upperBound(cand, need)

		// Left neighbor: rightSum <= need, miss = need - rightSum >= 0
		if idx-1 >= 0 {
			val := cand[idx-1]
			diff := need - val
			if diff < minDiff {
				minDiff = diff
			}
		}
		// Right neighbor: rightSum > need, miss = rightSum - need > 0
		if idx < len(cand) {
			val := cand[idx]
			diff := val - need
			if diff < minDiff {
				minDiff = diff
			}
		}
	}

	// Convert "miss from half" back to actual group difference.
	// |S_A - S_B| = 2 * |S_A - totalSum/2| + (totalSum % 2)
	//             = 2 * minDiff + parity
	answer := minDiff*2 + parity
	return int(answer)
}

// upperBound returns the first index i where a[i] > key.
// Equivalent to bisect_right in Python / upper_bound in C++.
//
// How sort.Search works (standard binary search):
//
//	It finds the smallest index i in [0, n) where f(i) is true,
//	assuming f is false for low indices and true for high ones.
//
//	lo, hi := 0, len(a)
//	for lo < hi {
//	    mid := lo + (hi-lo)/2
//	    if f(mid) { hi = mid } else { lo = mid + 1 }
//	}
//	return lo
//
//	Here f(i) = (a[i] > key), so it finds the first position
//	where the value exceeds key — exactly what we need to locate
//	both the left neighbor (idx-1, value ≤ key) and
//	the right neighbor (idx, value > key).
func upperBound(a []int64, key int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > key })
}

// Driver Code

func main() {
	testCases := [][]int{
		{3, 9, 7, 3},
		{-36, 36},
		{2, -1, 0, 4, -2, -9},
		{1000, -500, -400, -100},
		{5, 5, 5, 10, 15, 20},
	}

	for i, nums := range testCases {
		result := minimumDifference(nums)
		fmt.Printf("%d. nums: %v\n", i+1, nums)
		fmt.Printf("   result: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
