package main

func containsDuplicate(nums []int) bool {
	//making a map in go where num is the key and bool is the value
	hash := make(map[int]bool)
	for _, num := range nums {
		// if key is present in the map and the value is true then return true
		if _, ok := hash[num]; ok {
			return true
		}
		//otherwise the value for that key is set to true
		hash[num] = true
	}
	return false
}
