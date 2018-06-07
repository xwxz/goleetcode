package twosum

func TwoSum(inputArr []int, target int) (int, int) {
	var maps = make(map[int]int)
	for i := 0; i < len(inputArr); i++ {
		r := target - inputArr[i]
		if _, ok := maps[r]; ok {
			return maps[r], i
		}
		num := inputArr[i]
		maps[num] = i
	}
	return 0, 0
}
