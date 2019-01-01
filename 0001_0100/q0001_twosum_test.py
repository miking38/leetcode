def twoSum(nums, target):
    numMap = {}

    for idx1, num in enumerate(nums):
        numToFind = target - num
        if numToFind in numMap.keys():
            return [numMap[numToFind], idx1]
        else:
            numMap[num] = idx1 
        print(numMap)

    return [-1, -1]

result = twoSum([1,2,4], 6)
print(result)
