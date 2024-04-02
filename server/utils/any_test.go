package utils

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// Point 表示一个点的结构体，包含时间和水平线
type Point struct {
	Time  int
	Level int
}

func TestAny(t *testing.T) {
	var x int64 = -10
	var y uint = uint(x) // 将 int 类型的值转换为 uint 类型

	fmt.Println("Original int:", x)
	fmt.Println("Converted uint:", y)

	var z int64 = int64(y) // 将 uint 类型的值转换回 int 类型

	fmt.Println("Converted back to int:", z)
}

func main() {
	// 定义 p1 和 p2 两条曲线
	p1 := []Point{
		{0, 1},
		{2, 3},
		{4, 1},
		{6, 7},
	}

	p2 := []Point{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 6},
	}

	// 合并并排序两条曲线的时间
	allTimes := mergeAndSortTimes(p1, p2)

	// 获取两条曲线中较低的一条
	lowerCurve := getLowerCurve(allTimes, p1, p2)

	// 打印结果
	fmt.Println("\np1:")
	for _, point := range p1 {
		fmt.Printf("(%d, %d) ", point.Time, point.Level)
	}

	fmt.Println("\np2:")
	for _, point := range p2 {
		fmt.Printf("(%d, %d) ", point.Time, point.Level)
	}

	fmt.Println("\nLower curve:")
	for _, point := range lowerCurve {
		fmt.Printf("(%d, %d) ", point.Time, point.Level)
	}
}

// mergeAndSortTimes 函数用于合并并排序两条曲线的时间
func mergeAndSortTimes(p1, p2 []Point) []int {
	timeSet := make(map[int]bool)
	for _, point := range p1 {
		timeSet[point.Time] = true
	}
	for _, point := range p2 {
		timeSet[point.Time] = true
	}

	allTimes := make([]int, 0, len(timeSet))
	for time := range timeSet {
		allTimes = append(allTimes, time)
	}
	sort.Ints(allTimes)
	return allTimes
}

// getLowerCurve 函数用于获取两条曲线中的较低一条
func getLowerCurve(allTimes []int, p1, p2 []Point) []Point {
	var lowerCurve []Point
	for _, time := range allTimes {
		level1 := getLevelAtTime(time, p1)
		level2 := getLevelAtTime(time, p2)
		lowerLevel := math.Min(level1, level2)
		lowerCurve = append(lowerCurve, Point{Time: time, Level: int(lowerLevel)})
	}
	return lowerCurve
}

// getLevelAtTime 函数用于获取指定时间点的纵坐标值
func getLevelAtTime(time int, points []Point) float64 {
	// 倒序遍历所有点，找到第一个时间小于等于time的点
	for i := len(points) - 1; i >= 0; i-- {
		if points[i].Time <= time {
			return float64(points[i].Level)
		}
	}

	return 0 // 如果没有对应的时间点，则返回负无穷
}
