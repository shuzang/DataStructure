package set

import (
	"fmt"
	"testing"
)

func newSet(count int, start int) *Set {
	set := Set{}
	for i := start; i < (start + count); i++ {
		set.Add(fmt.Sprintf("item%d", i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := newSet(3, 0)
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item1") //should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item4") //should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestClear(t *testing.T) {
	set := newSet(3, 0)
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestDelete(t *testing.T) {
	set := newSet(3, 0)
	set.Delete("item2")
	if size := set.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	set := newSet(3, 0)
	has := set.Has("item2")
	if !has {
		t.Errorf("expected item2 to be there")
	}
	set.Delete("item2")
	has = set.Has("item2")
	if has {
		t.Errorf("expected item2 to be removed")
	}
	set.Delete("item1")
	has = set.Has("item1")
	if has {
		t.Errorf("expected item1 to be removed")
	}
}

func TestTraversal(t *testing.T) {
	set := newSet(3, 0)
	items := set.Traversal()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	set = newSet(520, 0)
	items = set.Traversal()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestSize(t *testing.T) {
	set := newSet(3, 0)
	items := set.Traversal()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = newSet(0, 0)
	items = set.Traversal()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = newSet(10000, 0)
	items = set.Traversal()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
}

func TestUnion(t *testing.T) {
	set1 := newSet(3, 0)
	set2 := newSet(2, 3)

	set3 := set1.Union(set2)

	if len(set3.Traversal()) != 5 {
		t.Errorf("wrong count, expected 5 and got %d", set3.Size())
	}
	//don't edit original sets
	if len(set1.Traversal()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Traversal()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestIntersection(t *testing.T) {
	set1 := newSet(3, 0)
	set2 := newSet(2, 0)

	set3 := set1.Intersection(set2)

	if len(set3.Traversal()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set3.Size())
	}
	//don't edit original sets
	if len(set1.Traversal()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Traversal()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestDifference(t *testing.T) {
	set1 := newSet(3, 0)
	set2 := newSet(2, 0)

	set3 := set1.Difference(set2)

	if len(set3.Traversal()) != 1 {
		t.Errorf("wrong count, expected 2 and got %d", set3.Size())
	}
	//don't edit original sets
	if len(set1.Traversal()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Traversal()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestSubset(t *testing.T) {
	set1 := newSet(3, 0)
	set2 := newSet(2, 0)

	if set1.Subset(set2) {
		t.Errorf("expected false and got true")
	}

	//don't edit original sets
	if len(set1.Traversal()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Traversal()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}

	//try real subsets
	set1 = newSet(2, 0)
	if !set1.Subset(set2) {
		t.Errorf("expected true and got false")
	}

	set1 = newSet(1, 0)
	if !set1.Subset(set2) {
		t.Errorf("expected true and got false")
	}
}
