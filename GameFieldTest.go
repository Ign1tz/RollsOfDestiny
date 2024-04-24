package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Column struct {
	first  int
	second int
	third  int
}

type Grid struct {
	left   Column
	middle Column
	right  Column
	value  int
}

type Playfield struct {
	hostGrid  Grid
	guestGrid Grid
}

type ColumnLogic interface {
	remove()
	add()
	clear()
	value()
}

func (x *Column) remove(s int) {
	if x.first == s {
		x.first = 0
	}
	if x.second == s {
		x.second = 0
	}
	if x.third == s {
		x.third = 0
	}
}

func (x *Column) add(s int) error {
	if s > 6 || s < 1 {
		return errors.New("Int invalid, value (" + strconv.Itoa(s) + ") out of bounds (1-6)!")
	}
	if x.first == 0 {
		x.first = s
	} else if x.second == 0 {
		x.second = s
	} else if x.third == 0 {
		x.third = s
	} else {
		return errors.New("Column is already full")
	}
	return nil
}

func (x *Column) clear() {
	x.first = 0
	x.second = 0
	x.third = 0
}

func (x Column) value() int {
	sum := 0
	numArr := [3]int{x.first, x.second, x.third}
	numbers := make(map[int]int)
	for _, num := range numArr {
		numbers[num] = numbers[num] + 1
	}
	for num, count := range numbers {
		sum += num * count * count
	}
	return sum
}

func main() {
	col := Column{0, 0, 0}
	col.add(2)
	fmt.Println(col)
	fmt.Println(col.value())
}
