func kClosest(points [][]int, k int) (result [][]int) {
	que := NewPriorityQueue[[]int]()

	for _, p := range points {
		que.Push(p, p[0]*p[0] + p[1]*p[1])
	}

	for k > 0 {
		k--
		result = append(result, que.Pop())
	}

	return
}

// ----------------------------------Priority queue---------------------------------------
type PriorityQueue[TKey any] interface {
	Push(key TKey, priority int)
	Peek() TKey
	Pop() TKey
	Len() int
}

type priorityQueue[TKey any] struct {
	items []*item[TKey]
}

type item[TKey any] struct {
	key      TKey
	priority int
}

func NewPriorityQueue[TKey any]() PriorityQueue[TKey] {
	return &priorityQueue[TKey]{items: make([]*item[TKey], 0)}
}

func (this *priorityQueue[TKey]) Push(key TKey, priority int) {
	this.items = append(this.items, &item[TKey]{key, priority})

	heapify_up(this, len(this.items)-1)
}

func (this *priorityQueue[TKey]) Peek() TKey {
	return this.items[0].key
}

func (this *priorityQueue[TKey]) Pop() TKey {
	result := this.items[0]
	this.items[0] = this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]

	if len(this.items) > 1 {
		this.heapify_down(0)
	}

	return result.key
}

func (this *priorityQueue[TKey]) Len() int {
	return len(this.items)
}

func heapify_up[TKey any](this *priorityQueue[TKey], i int) {
	for i != 0 {
		parent := (i - 1) / 2
		if this.less(i, parent) {
			this.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

func (this *priorityQueue[TKey]) heapify_down(i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i

	if left < len(this.items) && this.less(left, smallest) {
		smallest = left
	}

	if right < len(this.items) && this.less(right, smallest) {
		smallest = right
	}

	if smallest != i {
		this.swap(i, smallest)
		this.heapify_down(smallest)
	}
}

func (this *priorityQueue[TKey]) less(i int, j int) bool {
	return this.items[i].priority < this.items[j].priority
}

func (this *priorityQueue[TKey]) swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

// ----------------------------------Disjoint set---------------------------------------
type DisjointSet interface {
	GetParent(x int) int
	Unite(x int, y int)
}

type disjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(size int) DisjointSet {
	var ds = &disjointSet{make([]int, size+1), make([]int, size+1)}

	for i := 1; i <= size; i++ {
		ds.parent[i] = i
	}

	return ds
}

func (this *disjointSet) GetParent(x int) int {
	if this.parent[x] == x {
		return x
	}

	newParent := this.GetParent(this.parent[x])

	this.parent[x] = newParent

	return newParent
}

func (this *disjointSet) Unite(x int, y int) {
	parentX := this.GetParent(x)
	parentY := this.GetParent(y)

	if parentX == parentY {
		return
	}

	if this.rank[parentX] < this.rank[parentY] {
		parentX, parentY = parentY, parentX
	}

	this.parent[parentY] = parentX

	if this.rank[parentX] == this.rank[parentY] {
		this.rank[parentX]++
	}
}

// ----------------------------------Queue---------------------------------------
type Queue[T any] interface {
	Push(item T)
	Peek() T
	Pop() T
	Len() int
}

type queue[T any] struct {
	items []T
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{make([]T, 0)}
}

func (this *queue[T]) Push(item T) {
	this.items = append(this.items, item)
}

func (this *queue[T]) Peek() T {
	return this.items[0]
}

func (this *queue[T]) Pop() T {
	result := this.items[0]
	this.items = this.items[1:]
	return result
}

func (this *queue[T]) Len() int {
	return len(this.items)
}

// ----------------------------------Stack---------------------------------------
type Stack[T any] interface {
	Push(item T)
	Peek() T
	Pop() T
	Len() int
}

type stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{make([]T, 0)}
}

func (this *stack[T]) Push(item T) {
	this.items = append(this.items, item)
}

func (this *stack[T]) Peek() T {
	return this.items[len(this.items)-1]
}

func (this *stack[T]) Pop() T {
	result := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return result
}

func (this *stack[T]) Len() int {
	return len(this.items)
}

// ----------------------------------Set---------------------------------------
type Set[T comparable] interface {
	Add(item T) bool
	Remove(item T) bool
	Enumerate(func(item T))
	Len() int
}

type set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{make(map[T]bool)}
}

func (this *set[T]) Add(item T) bool {
	_, ok := this.items[item]

	if ok {
		return false
	}

	this.items[item] = true

	return true
}

func (this *set[T]) Remove(item T) bool {
	_, ok := this.items[item]

	if ok {
		delete(this.items, item)
		return true
	}

	return false
}

func (this *set[T]) Enumerate(action func(item T)) {
	for k := range this.items {
		action(k)
	}
}

func (this *set[T]) Len() int {
	return len(this.items)
}

// ----------------------------------Math---------------------------------------
func min[T Orderable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Orderable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func abs[T Orderable](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

type Orderable interface {
	~float32 | ~float64 |
		~int | ~int8 |
		~int16 | ~int32 |
		~int64 | ~uint |
		~uint8 | ~uint16 |
		~uint32 | ~uint64
}
