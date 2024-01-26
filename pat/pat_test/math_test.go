package pat_test

import (
	"testing"

	"github.com/Tom-abc/nynygo/pat"

	"errors"
)

type DetCase struct {
	m   [][]float64
	r   float64
	err error
}

func runDetTests(t *testing.T, tests []DetCase) {
	t.Helper()
	t.Parallel()
	for _, tt := range tests {
		r, err := pat.Determinant(tt.m)
		if !errors.Is(err, tt.err) {
			t.Errorf("expected error %v, got %v", tt.err, err)
		}
		if r != tt.r {
			t.Errorf("expected %f, got %f", tt.r, r)
		}
	}
}

func TestDeterminant(t *testing.T) {
	normalTests := []DetCase{
		{
			m: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			r: 0,
		},
		{
			m: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 8},
			},
			r: 3,
		},
	}
	errTests := []DetCase{
		{
			m: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
			},
			err: pat.ErrIllegalMatrix,
		},
	}
	t.Run("Normal", func(t *testing.T) {
		runDetTests(t, normalTests)
	})
	t.Run("Error", func(t *testing.T) {
		runDetTests(t, errTests)
	})
}
