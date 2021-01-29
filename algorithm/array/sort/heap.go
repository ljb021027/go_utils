package sort

func heapSort(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		minHeapDown(a, i, len(a))
	}
	for j := len(a) - 1; j >= 0; j-- {
		a[j], a[0] = a[0], a[j]
		minHeapDown(a, 0, j)
	}
}
func minHeapDown(a []int, i, l int) {
	temp := a[i]
	for k := 2*i + 1; k < l; k = 2*k + 2 {
		if k+1 < l && a[k] < a[k+1] {
			k++
		}
		if a[k] > temp {
			a[i] = a[k]
			//向下
			i = k
		}
	}
	//最终位置
	a[i] = temp
}
