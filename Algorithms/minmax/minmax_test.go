package minmax

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	min int
	max int
)

func TestFindNaive(t *testing.T) {
	var samples = []struct {
		min    int
		max    int
		sample []int
	}{
		{
			min:    -34,
			max:    1029438,
			sample: []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33},
		},
		{
			min:    -34,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, 33},
		},
		{
			min:    -35,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, -35},
		},
		{
			min:    -999,
			max:    -8,
			sample: []int{-999, -99, -34, -10, -8, -10},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{1111, 222, 33, 2, -1, -22},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{-22, -1, 2, 33, 222, 1111},
		},
		{
			min:    0,
			max:    0,
			sample: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range samples {
		t.Run(fmt.Sprintf("For min %v max %v", tt.min, tt.max), func(t *testing.T) {
			nNum, xNum := FindNaive(tt.sample)
			if nNum != tt.min ||
				xNum != tt.max {
				t.Errorf("Min is %v should be %v; max is %v should be %v", nNum, tt.min, xNum, tt.max)
			}
		})
	}
}

func BenchmarkFindNaive(b *testing.B) {

	rand.Seed(time.Now().Unix())

	var m = []struct {
		cnt int
	}{
		{cnt: 144},
		{cnt: 14400},
		{cnt: 1440000},
		{cnt: 14400000},
	}

	for _, mgntd := range m {

		b.Run(fmt.Sprintf("input_size_%d", mgntd.cnt), func(b *testing.B) {
			nums := rand.Perm(mgntd.cnt)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				min, max = FindNaive(nums)
			}
		})
	}
}

func TestFindNaiveDouble(t *testing.T) {
	var samples = []struct {
		min    int
		max    int
		sample []int
	}{
		{
			min:    -34,
			max:    1029438,
			sample: []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33},
		},
		{
			min:    -34,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, 33},
		},
		{
			min:    -35,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, -35},
		},
		{
			min:    -999,
			max:    -8,
			sample: []int{-999, -99, -34, -10, -8, -10},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{1111, 222, 33, 2, -1, -22},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{-22, -1, 2, 33, 222, 1111},
		},
		{
			min:    0,
			max:    0,
			sample: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range samples {
		t.Run(fmt.Sprintf("For min %v max %v", tt.min, tt.max), func(t *testing.T) {
			nNum, xNum := FindNaiveDouble(tt.sample)
			if nNum != tt.min ||
				xNum != tt.max {
				t.Errorf("Min is %v should be %v; max is %v should be %v", nNum, tt.min, xNum, tt.max)
			}
		})
	}
}

func BenchmarkFindNaiveDouble(b *testing.B) {

	rand.Seed(time.Now().Unix())

	var m = []struct {
		cnt int
	}{
		{cnt: 144},
		{cnt: 14400},
		{cnt: 1440000},
		{cnt: 14400000},
	}

	for _, mgntd := range m {

		b.Run(fmt.Sprintf("input_size_%d", mgntd.cnt), func(b *testing.B) {
			nums := rand.Perm(mgntd.cnt)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				min, max = FindNaiveDouble(nums)
			}
		})
	}
}

func TestFindOptim(t *testing.T) {
	var samples = []struct {
		min    int
		max    int
		sample []int
	}{
		{
			min:    -34,
			max:    1029438,
			sample: []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33},
		},
		{
			min:    -34,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, 33},
		},
		{
			min:    -35,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, -35},
		},
		{
			min:    -999,
			max:    -8,
			sample: []int{-999, -99, -34, -10, -8, -10},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{1111, 222, 33, 2, -1, -22},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{-22, -1, 2, 33, 222, 1111},
		},
		{
			min:    0,
			max:    0,
			sample: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range samples {
		t.Run(fmt.Sprintf("For min %v max %v", tt.min, tt.max), func(t *testing.T) {
			nNum, xNum := FindOptim(tt.sample)
			if nNum != tt.min ||
				xNum != tt.max {
				t.Errorf("Min is %v should be %v; max is %v should be %v", nNum, tt.min, xNum, tt.max)
			}
		})
	}
}

func BenchmarkFindOptim(b *testing.B) {

	rand.Seed(time.Now().Unix())

	var m = []struct {
		cnt int
	}{
		{cnt: 144},
		{cnt: 14400},
		{cnt: 1440000},
		{cnt: 14400000},
	}

	for _, mgntd := range m {

		b.Run(fmt.Sprintf("input_size_%d", mgntd.cnt), func(b *testing.B) {
			nums := rand.Perm(mgntd.cnt)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				min, max = FindOptim(nums)
			}
		})
	}
}

func TestFindNaiveCPU(t *testing.T) {
	var samples = []struct {
		min    int
		max    int
		sample []int
	}{
		{
			min:    -34,
			max:    1029438,
			sample: []int{-34, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, 1029438, 4, 33},
		},
		{
			min:    -34,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, 33},
		},
		{
			min:    -35,
			max:    1029438,
			sample: []int{1029438, 54, 2, 592, 0, 11, 11, -11, 432, 3, -5, 23, -6, -23, 65, -34, 4, -35},
		},
		{
			min:    -999,
			max:    -8,
			sample: []int{-999, -99, -34, -10, -8, -10},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{1111, 222, 33, 2, -1, -22},
		},
		{
			min:    -22,
			max:    1111,
			sample: []int{-22, -1, 2, 33, 222, 1111},
		},
		{
			min:    0,
			max:    0,
			sample: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range samples {
		t.Run(fmt.Sprintf("For min %v max %v", tt.min, tt.max), func(t *testing.T) {
			nNum, xNum := FindNaiveCPU(tt.sample)
			if nNum != tt.min ||
				xNum != tt.max {
				t.Errorf("Min is %v should be %v; max is %v should be %v", nNum, tt.min, xNum, tt.max)
			}
		})
	}
}

func BenchmarkFindNaiveCPU(b *testing.B) {

	rand.Seed(time.Now().Unix())

	var m = []struct {
		cnt int
	}{
		{cnt: 144},
		{cnt: 14400},
		{cnt: 1440000},
		{cnt: 14400000},
	}

	for _, mgntd := range m {

		b.Run(fmt.Sprintf("input_size_%d", mgntd.cnt), func(b *testing.B) {
			nums := rand.Perm(mgntd.cnt)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				min, max = FindNaiveCPU(nums)
			}
		})
	}
}
