package usort

import (
	"reflect"
	"testing"
)

func TestUSort(t *testing.T) {
	s1 := []int{1, 2, 4, 6, 7, 10, 4, 2}
	USort(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})

	if !reflect.DeepEqual(s1, []int{10, 7, 6, 4, 4, 2, 2, 1}) {
		t.Fail()
	}

	type People struct {
		Name string
		Age  int
	}

	tom := People{"tom", 31}
	libi := People{"libi", 30}
	jerry := People{"jerry", 25}

	s2 := []People{jerry, tom, libi}

	USort(s2, func(i, j int) bool {
		return s2[i].Age > s2[j].Age
	})

	if !reflect.DeepEqual(s2, []People{tom, libi, jerry}) {
		t.Fail()
	}
}
