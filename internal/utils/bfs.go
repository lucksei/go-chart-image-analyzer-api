package utils

func ContainersSpecSearch(m map[string]interface{}, res *[]any) {
	for k, v := range m {
		if k == "containers" {
			*res = append(*res, v.([]any))
		}
		if m, ok := v.(map[string]interface{}); ok {
			ContainersSpecSearch(m, res)
		}
	}
}
