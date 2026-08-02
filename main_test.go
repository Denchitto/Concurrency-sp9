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
		{-90, 1, 30, 0, -20, 100},
		{1000, 200, -300, 1000000},
		{0, 0, 0, 0},
		{-1, -3, -6, -100},
	}
	for _, v := range valid {
		sl, err := maximum(v)
		assert.NoError(t, err)
		assert.Equal(t, slices.Max(v), sl)
	}
	notValid := [][]int{
		{1},
		{},
		nil,
		//Для прошлого коммита, если отрицательные числа все таки не подходят
		//{-1, 3, 34, -50},
		//{-2, -100, -1000, -999},
	}
	for _, v := range notValid {
		sl, err := maximum(v)
		assert.Error(t, err)
		assert.Equal(t, 0, sl)
	}
}
