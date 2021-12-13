package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	height int
}

func (p *Point) String() string {
	return fmt.Sprintf("(x: %v, y: %v) = %v", p.x, p.y, p.height)
}

func read_heightmap(filename string) [][]*Point {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	heightmap := make([][]*Point, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]*Point, 0)
		for _, h := range strings.Split(line, "") {
			height, err := strconv.Atoi(h)
			if err != nil {
				panic(err)
			}
			p := &Point{x: len(row), y: len(heightmap), height: height}
			row = append(row, p)
		}
		heightmap = append(heightmap, row)
	}

	return heightmap
}

func get_adjacent_heights(heightmap [][]*Point, x int, y int) []*Point {
	adjacent_heights := make([]*Point, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i != 0 && j != 0 {
				continue
			}
			if x+i < 0 || x+i >= len(heightmap[0]) || y+j < 0 || y+j >= len(heightmap) {
				continue
			}
			adjacent_heights = append(adjacent_heights, heightmap[y+j][x+i])
		}
	}
	return adjacent_heights
}

// find points that are lower than all adjacent points
func part1(heightmap [][]*Point) []*Point {
	low_points := make([]*Point, 0)
	for y, row := range heightmap {
		for x, p := range row {
			adjacent_heights := get_adjacent_heights(heightmap, x, y)
			fmt.Printf("Points adjacent to %v: %v\n", p, adjacent_heights)
			is_low_point := true
			for _, adj_point := range adjacent_heights {
				if p.height >= adj_point.height {
					is_low_point = false
					break
				}
			}
			if (is_low_point) {
				low_points = append(low_points, p)
				fmt.Printf("Point %v is a low point\n", p)
			}
		}
	}
	return low_points
}

// do recursive stuff for part 2
func find_basin_points(low_point *Point, heightmap [][]*Point) map[int]*Point {
	basin_points := make(map[int]*Point)
	adjacent_points := get_adjacent_heights(heightmap, low_point.x, low_point.y)
	for _, adj_point := range adjacent_points {
		if adj_point.height > low_point.height && adj_point.height != 9 {
			for _, v := range find_basin_points(adj_point, heightmap) {
				basin_points[v.Hash()] = v
			}
			basin_points[adj_point.Hash()] = adj_point
		}
	}
	return basin_points
}

type basinList [][]*Point

func (p *Point) Hash() int {
	if p.x < p.y {
		return p.y * p.y + p.x
	} else {
		return p.x * p.x + p.y + p.x
	}
}

func find_basins(low_points []*Point, heightmap [][]*Point) [][]*Point {
	basins := make([][]*Point, 0)
	for _, p := range low_points {
		basin_points := find_basin_points(p, heightmap)
		fmt.Printf("Basin points for %v: %v\n", p, basin_points)
		p_hash := p.Hash()
		fmt.Printf("hash of p: %v\n", p_hash)
		basin_points[p_hash] = p
		points_slice := make([]*Point, 0)
		for _, v := range basin_points {
			points_slice = append(points_slice, v)
		}
		basins = append(basins, points_slice)
	}
	return basins
}

func (b basinList) Len() int {
	return len(b)
}

func (b basinList) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b basinList) Less(i, j int) bool {
	return len(b[i]) > len(b[j])
}

func main() {
	const filename = "input.txt"

	heightmap := read_heightmap(filename)

	fmt.Printf("%v\n", heightmap)

	sum := 0
	low_points := part1(heightmap)
	for _, p := range low_points {
		sum += p.height + 1
	}

	fmt.Printf("%v\n", sum)

	basin_points := find_basins(low_points, heightmap)
	
	sort.Sort(basinList(basin_points))
	fmt.Printf("%v\n", basin_points)

	for _, basin := range basin_points {
		fmt.Printf("basin size: %v\n", len(basin))
	}

	product_of_3_biggest_sizes := len(basin_points[0]) * len(basin_points[1]) * len(basin_points[2])
	fmt.Printf("Product of 3 biggest sizes: %v\n", product_of_3_biggest_sizes)
}