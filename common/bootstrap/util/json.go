/**
** @创建时间: 2022/4/11 13:53
** @作者　　: return
** @描述　　:
 */

package util

import "reflect"

func JsonMerge(dst, src map[string]interface{}) map[string]interface{} {
	return jsMerge(dst, src, 0)
}

func jsMerge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if dst == nil {
		dst = make(map[string]interface{})
	}
	if depth > 32 {
		return dst
		// panic("too deep!")
	}

	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := jsMapify(srcVal)
			dstMap, dstMapOk := jsMapify(dstVal)

			if srcMapOk && dstMapOk {
				srcVal = jsMerge(dstMap, srcMap, depth+1)
			}
		}

		dst[key] = srcVal
	}

	return dst
}

func jsMapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
