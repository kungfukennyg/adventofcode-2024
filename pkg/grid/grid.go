package grid

type Grid[T any] [][]T

func New[T any](data [][]T) *Grid[T] {
	g := Grid[T]{}
	for _, d := range data {
		g = append(g, d)
	}
	return &g
}

func (g Grid[T]) Get(x, y int) *T {
	if !g.boundsCheck(x, y) {
		return nil
	}

	return &g[y][x]
}

func (g Grid[T]) Set(x, y int, value T) {
	if !g.boundsCheck(x, y) {
		return
	}

	g[y][x] = value
}

func (g Grid[T]) boundsCheck(x, y int) bool {
	if y < 0 {
		return false
	}
	if x < 0 {
		return false
	}
	if y >= len(g) {
		return false
	}
	if x >= len(g[y]) {
		return false
	}

	return true
}

// Cardinal returns a slice containing pointers to the surrounding tiles to the coordinates given, placed directly above, below, left, and right.
func (g Grid[T]) Cardinal(x, y int) []*T {
	tiles := make([]*T, 0, 4)
	for _, coord := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		tiles = append(tiles, g.Get(x+coord[0], y+coord[1]))
	}
	return tiles
}

// Diagonals returns a slice containing pointers to the surrounding tiles to the coordinates given, above-left, above-right, below-left, and below-right
func (g Grid[T]) Diagonals(x, y int) []*T {
	tiles := make([]*T, 0, 4)
	for _, coord := range [][2]int{{-1, 1}, {-1, -1}, {1, 1}, {1, -1}} {
		tiles = append(tiles, g.Get(x+coord[0], y+coord[1]))
	}
	return tiles
}

// Adjacent returns a slice containing pointers to the surrounding tiles to the coordinates given, one step in all eight possible directions.
func (g Grid[T]) Adjacent(x, y int) []*T {
	tiles := make([]*T, 0, 8)
	tiles = append(tiles, g.Cardinal(x, y)...)
	tiles = append(tiles, g.Diagonals(x, y)...)
	return tiles
}
