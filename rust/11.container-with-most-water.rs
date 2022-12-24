/*
 * @lc app=leetcode id=11 lang=rust
 *
 * [11] Container With Most Water
 */

// @lc code=start
use std::cmp;

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut max = 0;
        let mut left = 0;
        let mut right = height.len() - 1;
        while left < right {
            if height[right] < height[left] {
                max = cmp::max(max, (right - left) as i32 * height[right]);
                right -= 1;
            } else {
                max = cmp::max(max, (right - left) as i32 * height[left]);
                left += 1;
            }
        }
        max
    }
}
// @lc code=end

struct Solution {}
