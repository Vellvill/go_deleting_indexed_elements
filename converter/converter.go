package converter

func Map(mapFunc func(s string) string) []string {
	strs := make([]string, 5, 6)
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}

	return mapped
}
