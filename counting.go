package counting

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

// findMinMax finds the minimum and maximum integers in a slice.
func findMinMax(a []int) (min, max int) {
	max = minInt
	min = maxInt
	for _, n := range a {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return min, max
}

// Sort sorts a slice of integers using the counting sort algorithm.
func Sort(a []int) []int {
	if len(a) == 0 || a == nil {
		return a
	}

	min, max := findMinMax(a)

	K := (max - min) + 1
	N := len(a)

	b := make([]int, N)
	c := make([]int, K*2+1)

	for i := 0; i < N; i++ {
		c[a[i]+K]++
	}

	for i := 1; i < K*2+1; i++ {
		c[i] = c[i-1] + c[i]
	}

	i := N - 1
	for i >= 0 {
		b[c[a[i]+K]-1] = a[i]
		c[a[i]+K]--
		i--
	}

	return b
}
