package cooonfig

func deepMerge(dst, src map[string]any) {
	for key, srcValue := range src {
		dstValue, exists := dst[key]
		if !exists {
			dst[key] = srcValue
			continue
		}
		dstMap, dstIsMap := dstValue.(map[string]any)
		srcMap, srcIsMap := srcValue.(map[string]any)

		if dstIsMap && srcIsMap {
			deepMerge(dstMap, srcMap)
		} else {
			// new value, overrides
			dst[key] = srcValue
		}

	}

}
