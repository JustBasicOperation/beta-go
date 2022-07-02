package utils

// 空结构体
type void struct {
}

// DistinctSliceInt 切片类型去重
func DistinctSliceInt(slice []int64) []int64 {
	var value void
	var res []int64
	tempMap := make(map[int64]void)
	for _, i := range slice {
		l := len(tempMap)
		tempMap[i] = value
		if len(tempMap) > l {
			res = append(res, i)
		}
	}
	return res
}

// DistinctSliceString 切片类型去重
func DistinctSliceString(slice []string) []string {
	var value void
	var res []string
	tempMap := make(map[string]void)
	for _, i := range slice {
		l := len(tempMap)
		tempMap[i] = value
		if len(tempMap) > l {
			res = append(res, i)
		}
	}
	return res
}
