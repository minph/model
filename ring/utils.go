package ring

// replace 重置指针位置
func replace(index, count int) int {
	if index < 0 {
		return count + index%count
	}

	if count <= index {
		return index % count
	}

	return index

}
