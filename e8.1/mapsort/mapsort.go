// Package mapsort contains interfaces to sort maps by either their keys or values.
// Keys or values must be of uint, int, float, or string types.
// When sorting by Keys, value types are not checked,
// when sorting by Values, key types are not checked.
// See mapsortDemo.go for usage instructions.
package mapsort

import "sort"

// MsFmtMap creates a map of interface{} k/v pairs,
// formatted for use with MapSort functions
type MsFmtMap map[interface{}]interface{}

// Entry represents a single key/value pair
// Field values are of type interface{} to allow dynamic comparisons
type Entry struct {
	Key interface{}
	Val interface{}
}

// KeySort is an iterface type for sorting a map by Key
type KeySort []Entry

func (s KeySort) Len() int { return len(s) }
func (s KeySort) Less(i, j int) bool {
	switch s[i].Key.(type) {
	case int:
		if ii, ok := s[i].Key.(int); ok {
			if ji, ok := s[j].Key.(int); ok {
				return ii < ji
			}
		}
	case float32:
		if ii, ok := s[i].Key.(float32); ok {
			if ji, ok := s[j].Key.(float32); ok {
				return ii < ji
			}
		}
	case float64:
		if ii, ok := s[i].Key.(float64); ok {
			if ji, ok := s[j].Key.(float64); ok {
				return ii < ji
			}
		}
	case uint8:
		if ii, ok := s[i].Key.(uint8); ok {
			if ji, ok := s[j].Key.(uint8); ok {
				return ii < ji
			}
		}
	case uint16:
		if ii, ok := s[i].Key.(uint16); ok {
			if ji, ok := s[j].Key.(uint16); ok {
				return ii < ji
			}
		}
	case uint32:
		if ii, ok := s[i].Key.(uint32); ok {
			if ji, ok := s[j].Key.(uint32); ok {
				return ii < ji
			}
		}
	case uint64:
		if ii, ok := s[i].Key.(uint64); ok {
			if ji, ok := s[j].Key.(uint64); ok {
				return ii < ji
			}
		}
	case int8:
		if ii, ok := s[i].Key.(int8); ok {
			if ji, ok := s[j].Key.(int8); ok {
				return ii < ji
			}
		}
	case int16:
		if ii, ok := s[i].Key.(int16); ok {
			if ji, ok := s[j].Key.(int16); ok {
				return ii < ji
			}
		}
	case int32:
		if ii, ok := s[i].Key.(int32); ok {
			if ji, ok := s[j].Key.(int32); ok {
				return ii < ji
			}
		}
	case int64:
		if ii, ok := s[i].Key.(int64); ok {
			if ji, ok := s[j].Key.(int64); ok {
				return ii < ji
			}
		}
	case string:
		if ii, ok := s[i].Key.(string); ok {
			if ji, ok := s[j].Key.(string); ok {
				return ii < ji
			}
		}
	default:
		panic("Unknown/Type not comparable")
	}
	return false
}
func (s KeySort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// MapSort sorts the formatted map passed to it as an argument
func (s KeySort) MapSort(m MsFmtMap) KeySort {
	var es KeySort
	for k, v := range m {
		es = append(es, Entry{Key: k, Val: v})
	}
	sort.Sort(es)
	return es
}

// ValSort is an iterface type for sorting a map by Value
type ValSort []Entry

func (s ValSort) Len() int { return len(s) }
func (s ValSort) Less(i, j int) bool {
	switch s[i].Val.(type) {
	case int:
		if ii, ok := s[i].Val.(int); ok {
			if ji, ok := s[j].Val.(int); ok {
				return ii < ji
			}
		}
	case float32:
		if ii, ok := s[i].Val.(float32); ok {
			if ji, ok := s[j].Val.(float32); ok {
				return ii < ji
			}
		}
	case float64:
		if ii, ok := s[i].Val.(float64); ok {
			if ji, ok := s[j].Val.(float64); ok {
				return ii < ji
			}
		}
	case uint8:
		if ii, ok := s[i].Val.(uint8); ok {
			if ji, ok := s[j].Val.(uint8); ok {
				return ii < ji
			}
		}
	case uint16:
		if ii, ok := s[i].Val.(uint16); ok {
			if ji, ok := s[j].Val.(uint16); ok {
				return ii < ji
			}
		}
	case uint32:
		if ii, ok := s[i].Val.(uint32); ok {
			if ji, ok := s[j].Val.(uint32); ok {
				return ii < ji
			}
		}
	case uint64:
		if ii, ok := s[i].Val.(uint64); ok {
			if ji, ok := s[j].Val.(uint64); ok {
				return ii < ji
			}
		}
	case int8:
		if ii, ok := s[i].Val.(int8); ok {
			if ji, ok := s[j].Val.(int8); ok {
				return ii < ji
			}
		}
	case int16:
		if ii, ok := s[i].Val.(int16); ok {
			if ji, ok := s[j].Val.(int16); ok {
				return ii < ji
			}
		}
	case int32:
		if ii, ok := s[i].Val.(int32); ok {
			if ji, ok := s[j].Val.(int32); ok {
				return ii < ji
			}
		}
	case int64:
		if ii, ok := s[i].Val.(int64); ok {
			if ji, ok := s[j].Val.(int64); ok {
				return ii < ji
			}
		}
	case string:
		if ii, ok := s[i].Val.(string); ok {
			if ji, ok := s[j].Val.(string); ok {
				return ii < ji
			}
		}
	default:
		panic("Unknown")
	}
	return false
}
func (s ValSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// MapSort sorts the formatted map passed to it as an argument
func (s ValSort) MapSort(m MsFmtMap) ValSort {
	var es ValSort
	for k, v := range m {
		es = append(es, Entry{Key: k, Val: v})
	}
	sort.Sort(es)
	return es
}
