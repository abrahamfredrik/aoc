package utils

import "fmt"

func RemoveFromSlice[T string | int | rune](s []T, remove T) []T {
	for i, v := range s {
		if v == remove {
			s[i] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	panic(fmt.Sprintf("Could not remove %v from slice %v", remove, s))
}

func RemoveDuplicates[T string | int | rune](s []T) []T {
	m := map[T]bool{}
	list := []T{}
	for _, v := range s {
		if !m[v] {
			m[v] = true
			list = append(list, v)
		}
	}
	return list
}

func CopySlice[T string | int | rune](s []T) []T {
	n := make([]T, len(s))
	for i, v := range s {
		n[i] = v
	}
	return n
}
