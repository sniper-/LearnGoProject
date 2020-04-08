package leetcode

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		//index和value互换，互换后m[target-value]值若存在，即可返回index和i
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		//m[target-value]值不存在，则将当前index和value互换
		m[v] = i
	}
	return nil
}
