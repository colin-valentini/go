package matrix

type Matrix2D[T any] struct {
	mat [][]T
}

func NewMatrix2D[T any](n, m int) *Matrix2D[T] {
	mat := make([][]T, n)
	for r := 0; r < n; r++ {
		mat[r] = make([]T, m)
	}
	return &Matrix2D[T]{mat: mat}
}

func NewMatrix2DFromSlice[T any](mat [][]T) *Matrix2D[T] {
	return &Matrix2D[T]{mat: mat}
}

func (m Matrix2D[T]) Value(row, col int) (T, bool) {
	var value T
	if m.OutOfRange(row, col) {
		return value, false
	}
	return m.mat[row][col], true
}

func (m Matrix2D[T]) SetValue(row, col int, value T) bool {
	if m.OutOfRange(row, col) {
		return false
	}
	m.mat[row][col] = value
	return true
}

func (m Matrix2D[T]) OutOfRange(row, col int) bool {
	return !m.InRange(row, col)
}

func (m Matrix2D[T]) InRange(row, col int) bool {
	numRows, numCols := m.Shape()
	return 0 <= row && row < numRows && 0 <= col && col < numCols
}

func (m Matrix2D[T]) Shape() (numRows, numCols int) {
	numRows = len(m.mat)
	if numRows > 0 {
		numCols = len(m.mat[0])
	}
	return numRows, numCols
}

func (m Matrix2D[T]) Iter() *Matrix2DIterator[T] {
	return newMatrix2DIterator(m)
}

type Matrix2DIterator[T any] struct {
	row, col int
	mat      Matrix2D[T]
}

func newMatrix2DIterator[T any](matrix Matrix2D[T]) *Matrix2DIterator[T] {
	return &Matrix2DIterator[T]{row: 0, col: -1, mat: matrix}
}

func (it *Matrix2DIterator[T]) HasNext() bool {
	_, _, ok := it.peekNext()
	return ok
}

func (it *Matrix2DIterator[T]) Next() bool {
	r, c, ok := it.peekNext()
	if ok {
		it.row, it.col = r, c
	}
	return ok
}

func (it *Matrix2DIterator[T]) peekNext() (r, c int, ok bool) {
	numRows, numCols := it.mat.Shape()
	if it.col < numCols-1 {
		return it.row, it.col + 1, true
	}
	if it.row < numRows-1 {
		return it.row + 1, 0, true
	}
	return r, c, false
}

func (it *Matrix2DIterator[T]) Value() (T, bool) {
	return it.mat.Value(it.row, it.col)
}
