# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  pool = {}
  i = nums.length - 1
  while n = nums.pop
    break if pool[n]
    pool[target - n] = i
    i -= 1
  end
  [nums.length, pool[n]]
end

nums = [0, 3, 2, 5, 0, 10]
target = 7

p two_sum(nums, target)
