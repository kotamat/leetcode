require 'benchmark'
# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  pool = {}
  i = nums.length - 1
  while n = nums.pop
    return [nums.length, pool[n]] if pool[n]
    pool[target - n] = i
    i -= 1
  end
end

iterate_num = 1000000

nums = [0, 3, 2, 5, 0, 10]
target = 7

Benchmark.bm do |r|
  r.report "two sum" do
    iterate_num.times do
      two_sum(nums, target)
    end
  end
end
nums = [0, 3, 2, 5, 0, 10]
target = 7
p two_sum(nums, target)
