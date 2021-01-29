package sort

func qsort(a []int, start, end int) {
	if start >= end {
		return
	}
	p := partition(a, start, end)
	qsort(a, start, p-1)
	qsort(a, p+1, end)
}

func partition(a []int, start, end int) int {
	p := a[end]
	i := start
	//遍历，找到所有比partition小的，交换到a[0:i-1]里面去
	//,则a[0:i-1]全部小于partition。
	//,a[i:end-1] > partition
	//最后交换 a[i] 和 a[end]
	for j := start; j < end; j++ {
		if a[j] < p {
			//比partition小
			//加入到a[0:i-1]的尾部 也就是a[i]
			//swap a[i] a[j]
			temp := a[j]
			a[j] = a[i]
			a[i] = temp
			i++
		}
	}
	//swap a[i] a[end]
	a[end] = a[i]
	a[i] = p
	return i
}
