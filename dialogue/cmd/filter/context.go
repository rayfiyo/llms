package filter

const errNum = 235

func Context(ctx []int) []int {
	var filtered []int
	for _, v := range ctx {
		if v != errNum {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
