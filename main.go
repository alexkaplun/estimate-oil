package main

/*
  To estimate how much oil can be stored at any position, we should know what the max level on the left and on the right is from this position.
  Then this position can store up to min(maxLeft, maxRight) minus level on this position.

  To count overall capacity, we will move from left and from right to the center (actually to the highest level, should it be in the middle
  or on the edges), and sum up to the total how much current position can hold.
*/
func estimate(levels []int) int {
	if len(levels) == 0 {
		return 0
	}

	// initial positions are at the edges of the slice
	left, right := 0, len(levels)-1

	// start with empty max values on the left and on the right
	maxOnTheLeft, maxOnTheRight := 0, 0
	total := 0

	// move to the center
	for left < right {

		// decide which direction to go. If the level on the left is lower than on the right - move left, otherwise move right
		if levels[left] < levels[right] {
			// check if the left level is higher than maxOnTheLeft
			if levels[left] >= maxOnTheLeft {
				// found the higher left level - just move to the next position and save max left level
				maxOnTheLeft = levels[left]
			} else {
				// left level is lower than maxOnTheLeft - sum up to the total that the current position can hold up to maxOnTheLeft
				total += maxOnTheLeft - levels[left]
			}
			// go from left to right
			left++
		} else {
			// check if the right level is higher than maxOnTheRight
			if levels[right] >= maxOnTheRight {
				// found the higher right level - just move to the next position and save max right level
				maxOnTheRight = levels[right]
			} else {
				// right level is lower than maxOnTheRight - sum up to the total that the current position can hold up to maxOnTheRight
				total += maxOnTheRight - levels[right]
			}
			// go from right to left
			right--
		}
	}
	return total
}
