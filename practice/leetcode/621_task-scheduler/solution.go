package taskscheduler

// LeetCode #621.
// https://leetcode.com/problems/task-scheduler/

import "container/heap"

type Solver struct {
	tasks []byte
	n     int
}

func NewSolver(tasks []byte, n int) *Solver {
	return &Solver{tasks: tasks, n: n}
}

func (s *Solver) Solve() int {
	// - Scan for character frequencies
	// - Construct a max heap from character frequencies
	// - Loop until the heap is exhausted
	//   - Check if we can pop from the queue based on current time
	//     - If yes, push the popped value into the heap
	//   - Pop a value from the heap
	//   - Decrement frequency and set next available time
	//   - If frequency > 0, push value onto queue
	//   - Increment counter by 1
	// - Return total count
	h := s.makeHeap()
	q := queue([]*task{})
	time := 0
	for h.Len() > 0 || len(q) > 0 {
		if len(q) > 0 && time-q.Peek().lastExec > s.n {
			heap.Push(h, q.Pop())
		}
		if h.Len() == 0 {
			// No tasks available to run, idle.
			time++
			continue
		}
		cf := heap.Pop(h).(*task)
		cf.dec()
		if cf.count > 0 {
			cf.executedAt(time)
			q.Push(cf)
		}
		time++
	}
	return time
}

func (s *Solver) makeHeap() *maxHeap {
	fm := map[byte]*task{}
	for _, name := range s.tasks {
		if _, ok := fm[name]; !ok {
			fm[name] = newTask()
		}
		fm[name].inc()
	}
	fs := make([]*task, 0, len(fm))
	for _, freq := range fm {
		fs = append(fs, freq)
	}
	h := maxHeap(fs)
	heap.Init(&h)
	return &h
}

type task struct {
	count, lastExec int
}

func newTask() *task {
	return &task{}
}

func (t *task) executedAt(time int) {
	t.lastExec = time
}

func (t *task) inc() {
	t.count++
}

func (t *task) dec() {
	t.count--
}

type maxHeap []*task

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(v any) {
	*h = append(*h, v.(*task))
}

func (h *maxHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type queue []*task

func (q *queue) Peek() *task {
	return (*q)[0]
}

func (q *queue) Push(v *task) {
	*q = append(*q, v)
}

func (q *queue) Pop() *task {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
