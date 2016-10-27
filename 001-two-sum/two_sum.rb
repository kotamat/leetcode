# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  until
    result = nums.index(target - nums.pop)
  end
  [result, nums.length]
end

nums = [3, 2, 5]
target = 7

p two_sum(nums, target)
