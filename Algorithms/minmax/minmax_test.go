package minmax

import (
	"math/rand"
	"testing"
	"time"
)

var (
	min int
	max int
)

func TestFindNaive(t *testing.T) {
	var nNum int
	var xNum int

	nums := []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33}
	nNum, xNum = FindNaive(nums)

	if nNum != -34 ||
		xNum != 1029438 {
		t.Errorf("Min is %v should be -34; max is %v should be 1029438", nNum, xNum)
	}
}

func BenchmarkFindNaive(b *testing.B) {

	rand.Seed(time.Now().Unix())

	nums := rand.Perm(9999990)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		min, max = FindNaive(nums)
	}
}

func TestFindNaiveDouble(t *testing.T) {
	var nNum int
	var xNum int

	nums := []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33}
	nNum, xNum = FindNaiveDouble(nums)

	if nNum != -34 ||
		xNum != 1029438 {
		t.Errorf("Min is %v should be -34; max is %v should be 1029438", nNum, xNum)
	}
}

func BenchmarkFindNaiveDouble(b *testing.B) {

	rand.Seed(time.Now().Unix())

	nums := rand.Perm(9999990)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		min, max = FindNaiveDouble(nums)
	}
}

func TestFindOptim(t *testing.T) {
	var nNum int
	var xNum int

	nums := []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33}
	nNum, xNum = FindOptim(nums)

	if nNum != -34 ||
		xNum != 1029438 {
		t.Errorf("Min is %v should be -34; max is %v should be 1029438", nNum, xNum)
	}
}

func BenchmarkFindOptim(b *testing.B) {

	rand.Seed(time.Now().Unix())

	nums := rand.Perm(9999990)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		min, max = FindOptim(nums)
	}
}
