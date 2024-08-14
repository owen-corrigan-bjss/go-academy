package main

import (
	"reflect"
	"testing"
)

func TestAssignment4(t *testing.T) {

	t.Run("sort a slice into ascending order", func(t *testing.T) {
		arr := []int{2, 1, 3}
		got := SortSliceAscending(arr)
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v got %v", want, got)
		}
	})
	t.Run("sort a slice into descending order", func(t *testing.T) {
		arr := []int{1, 3, 2}
		got := SortArrayDescending(arr)
		want := []int{3, 2, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v got %v", want, got)
		}
	})

	t.Run("sort a slice into descending order but dont mutate the array", func(t *testing.T) {
		arr := []int{2, 1, 3}
		got := SortArrayDescending(arr)
		want := []int{3, 2, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v got %v", want, got)
		}
		if (arr[0] != 2) && (arr[1] != 1) && (arr[2] != 3) {
			t.Errorf("original array mutated: %v\n", arr)
		}
	})

	t.Run("will sort an array into two different odd and even arrays", func(t *testing.T) {
		arr := []int{4, 5, 3, 2, 7}
		gotOdds, gotEvens := SortToOddsAndEvens(arr)
		wantOdds := []int{3, 5, 7}
		wantEvens := []int{4, 2}

		if !reflect.DeepEqual(gotOdds, wantOdds) {
			t.Errorf("wanted %v got %v\n", wantOdds, gotOdds)
		}
		if !reflect.DeepEqual(gotEvens, wantEvens) {
			t.Errorf("wanted %v got %v\n", wantEvens, gotEvens)
		}
	})

	t.Run("will sort an array into two different odd and even arrays with longer original array", func(t *testing.T) {
		arr := []int{4, 5, 3, 2, 7, 81, 32, 12, 13, 100}
		gotOdds, gotEvens := SortToOddsAndEvens(arr)
		wantOdds := []int{3, 5, 7, 13, 81}
		wantEvens := []int{100, 32, 12, 4, 2}

		if !reflect.DeepEqual(gotOdds, wantOdds) {
			t.Errorf("wanted %v got %v\n", wantOdds, gotOdds)
		}
		if !reflect.DeepEqual(gotEvens, wantEvens) {
			t.Errorf("wanted %v got %v\n", wantEvens, gotEvens)
		}
	})
}
