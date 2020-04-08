package leetcode

func MovingCount(m int, n int, k int) int {
	count := 0

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			row := i
			col := j
			sum := 0
			for row!=0 {
				sum += row%10
				row /= 10
			}
			for col!=0{
				sum += col%10
				col /= 10
			}
			if sum <= k{
				count++
			}
		}
	}
	return count
}