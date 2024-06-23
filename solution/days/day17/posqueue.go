package day17

// purpose-built binary heap to see which of the traversed positions currently has the lowest heat loss
type positionPriorityQueue struct {
	heap []*lavaTraveralPosition
	size int `default:"0"`
}

func (q *positionPriorityQueue) push(position *lavaTraveralPosition) {
	q.heap = append(q.heap, position)

	idx := q.size

	for idx > 0 && q.heap[(idx-1)/2].heatLoss > q.heap[idx].heatLoss {
		q.swap((idx-1)/2, idx)
		idx = (idx - 1) / 2
	}

	q.size++
}

func (q *positionPriorityQueue) heapify(idx int) {
	if q.size <= 1 {
		return
	}

	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2
	smallestIdx := idx

	if leftIdx < q.size && q.heap[leftIdx].heatLoss < q.heap[smallestIdx].heatLoss {
		smallestIdx = leftIdx
	}
	if rightIdx < q.size && q.heap[rightIdx].heatLoss < q.heap[smallestIdx].heatLoss {
		smallestIdx = rightIdx
	}

	if smallestIdx != idx {
		q.swap(idx, smallestIdx)
		q.heapify(smallestIdx)
	}
}

func (q *positionPriorityQueue) pop() *lavaTraveralPosition {
	if q.size == 0 {
		return nil
	}
	lowest := q.heap[0]
	q.size--
	q.heap[0] = q.heap[q.size]
	q.heap = q.heap[:q.size]
	q.heapify(0)

	return lowest
}

func (q *positionPriorityQueue) swap(i, j int) {
	temp := q.heap[i]
	q.heap[i] = q.heap[j]
	q.heap[j] = temp
}
