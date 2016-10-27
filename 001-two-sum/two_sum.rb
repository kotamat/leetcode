# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  i = 0
  while n = nums.pop do
    result = nums.find_index(target - n)
    return [result, nums.length] if result != nil
    i += 1
  end
end

nums = [3, 2, 5]
target = 7

p two_sum(nums, target)
