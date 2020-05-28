package examples

import (
	"sort"
	"testing"
)

var persons Persons

func init() {
	persons = Persons{
		{
			Name:   "test123",
			Age:    20,
			Height: 170,
		},
		{
			Name:   "test1",
			Age:    22,
			Height: 170,
		},
		{
			Name:   "test12",
			Age:    21,
			Height: 175,
		},
	}
}

func TestSortByName(t *testing.T) {
	sort.Sort(SortByName{persons})
	for i, _ := range persons {
		if i == 0 {
			continue
		}
		if len(persons[i].Name) > len(persons[i-1].Name) {
			t.Errorf("SortByName invalid, %v", persons)
		}
	}
}

func TestSortByAge(t *testing.T) {
	sort.Sort(SortByAge{persons})
	for i, _ := range persons {
		if i == 0 {
			continue
		}
		if persons[i].Age > persons[i-1].Age {
			t.Errorf("SortByAge invalid, %v", persons)
		}
	}
}

func TestSortByHeight(t *testing.T) {
	//sort.Sort(SortByHeight{persons})
	sort.Stable(SortByHeight{persons})
	for i, _ := range persons {
		if i == 0 {
			continue
		}
		if persons[i].Height > persons[i-1].Height {
			t.Errorf("SortByHeight invalid, %v", persons)
		}
	}
	t.Log(persons)
}
