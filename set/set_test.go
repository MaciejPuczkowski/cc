package set

import (
	"testing"
)

func TestSetUnion(t *testing.T) {
	s1 := From([]int{1, 2, 3})
	s2 := From([]int{3, 4, 5})
	s3 := From([]int{1, 2, 3, 4, 5})
	result := Union(s1, s2, s3)
	expected := From([]int{1, 2, 3, 4, 5})
	if !AreEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSetProduct(t *testing.T) {
	s1 := From([]int{1, 2, 3})
	s2 := From([]int{3, 4, 5})
	s3 := From([]int{1, 2, 3, 4, 5})
	result := Product(s1, s2, s3)
	expected := From([]int{3})
	if !AreEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSetXor(t *testing.T) {
	s1 := From([]int{1, 2, 3})
	s2 := From([]int{3, 4, 5})
	s3 := From([]int{1, 2, 3, 4, 5})
	result := Xor(s1, s2, s3)
	expected := From([]int{1, 2, 4, 5})
	if !AreEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
