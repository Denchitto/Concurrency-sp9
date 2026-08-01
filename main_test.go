package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	valid := []int{1000, 302, 2, 3, 4, 23}
	for _, v := range valid {
		sl := generateRandomElements(v)
		assert.NotNil(t, sl)
		assert.Len(t, sl, v)
		assert.NotEmpty(t, sl)
	}
	notValid := []int{0, 1, -1, -2, -100}
	for _, v := range notValid {
		sl := generateRandomElements(v)
		assert.Nil(t, sl)
	}
}

func TestMaximum(t *testing.T) {
	valid := [][]int{
		{1, 3, 4, 2, 12},
		{0, 1, 30, 100},
		{1000, 200, 300, 1000000},
		{0, 0, 0, 0},
	}
	for _, v := range valid {
		sl := maximum(v)
		assert.NotEqual(t, -1, sl)
		assert.Equal(t, slices.Max(v), sl)
	}
	notValid := [][]int{
		{1},
		{},
		nil,
		make([]int, 5, 10),
		{-1, 3, 34, -50},
	}
	for _, v := range notValid {
		sl := maximum(v)
		assert.Equal(t, -1, sl)
	}
}
