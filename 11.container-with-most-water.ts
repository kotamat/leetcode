/*
 * @lc app=leetcode id=11 lang=typescript
 *
 * [11] Container With Most Water
 */

// @lc code=start
function maxArea(height: number[]): number {
  let left = 0;
  let right = height.length - 1;
  let max = 0;
  while (left < right) {
    // The longer line should be the result
    if (height[left] < height[right]) {
      max = Math.max(max, (right - left) * height[left]);
      left++;
    } else {
      max = Math.max(max, (right - left) * height[right]);
      right--;
    }
  }
  return max;
}
// @lc code=end
