package main

import "maps"

func merge(m1, m2 map[string]any) map[string]any {
	return maps.Clone(m2)
}
