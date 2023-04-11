package quicksort

// QuickSort sorts the given slice using the comparator to determine
// if two items are less than another.
func QuickSort(in []any, less func(i, j int) bool) {
	newQuickSorter(in, less).sort(0, len(in)-1)
}

type quickSorter struct {
	in   []any
	less func(i, j int) bool
}

func newQuickSorter(in []any, less func(i, j int) bool) *quickSorter {
	return &quickSorter{in: in, less: less}
}

func (s *quickSorter) sort(left, right int) {
	if right <= left {
		return
	}
	p := s.partition(left, right)
	s.sort(left, p-1)
	s.sort(p+1, right)
}

func (s *quickSorter) partition(left, right int) int {
	p := left
	for i := left; i < right; i++ {
		if s.less(i, right) {
			s.swap(i, p)
			p++
		}
	}
	s.swap(p, right)
	return p
}

func (s *quickSorter) swap(i, j int) {
	s.in[i], s.in[j] = s.in[j], s.in[i]
}
