class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
            #set creation for previous numbers
            prev = set()
            for num in nums:
                #if number was already added to the set, return True
                if num in prev:
                    return True
                #otherwise add the number
                prev.add(num)
            return False