package Types

import (
	"errors"
	"strconv"
)

type Column struct {
	First     int
	Second    int
	Third     int
	GridId    string
	Placement int
}

type ColumnLogic interface {
	Remove()
	Add()
	Clear()
	Value()
	shift()
	IsFull()
}

func (c *Column) Remove(s int) int {
	numberOfRemoved := 0
	removed := false
	if c.First == s {
		c.First = 0
		numberOfRemoved++
		removed = true
	}
	if c.Second == s {
		c.Second = 0
		numberOfRemoved++
		removed = true
	}
	if c.Third == s {
		c.Third = 0
		numberOfRemoved++
	}
	if removed {
		c.shift()
	}
	return numberOfRemoved
}

func (c *Column) shift() {
	if c.Second != 0 {
		if c.First == 0 {
			c.First = c.Second
			c.Second = 0
		}
	}
	if c.Third != 0 {
		if c.First == 0 {
			c.First = c.Third
			c.Third = 0
		} else if c.Second == 0 {
			c.Second = c.Third
			c.Third = 0
		}
	}
}

func (c *Column) Add(s int) error {
	if s > 6 || s < 1 {
		return errors.New("Int invalid, Value (" + strconv.Itoa(s) + ") out of bounds (1-6)!")
	}
	if c.First == 0 {
		c.First = s
	} else if c.Second == 0 {
		c.Second = s
	} else if c.Third == 0 {
		c.Third = s
	} else {
		return errors.New("Column is already full")
	}
	return nil
}

func (c *Column) Clear() {
	c.First = 0
	c.Second = 0
	c.Third = 0
}

func (c Column) Value() int {
	sum := 0
	numArr := [3]int{c.First, c.Second, c.Third}
	numbers := make(map[int]int)
	for _, num := range numArr {
		numbers[num] = numbers[num] + 1
	}
	for num, count := range numbers {
		sum += num * count * count
	}
	return sum
}
func (c *Column) IsFull() bool {
	return c.First != 0 && c.Second != 0 && c.Third != 0
}
