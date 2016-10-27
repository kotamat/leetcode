# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  numsLength = nums.length
  i = 0
  while i < numsLength do
    result = nums[(i + 1)..-1].find_index(target - nums[i])
    return [i, result + i + 1] if result != nil
    i += 1
  end
end

nums = [3, 2, 5]
target = 7

puts two_sum(nums, target)
