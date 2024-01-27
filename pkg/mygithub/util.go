package mygithub

// splitAndMapFromSlice takes a slice of strings, splits each string at the first period,
// and returns a map where each key is the first part of the split and each value is the original string.
func splitAndMapFromSlice(slice []string) map[string]string {
	m := make(map[string]string)
	for _, s := range slice {
		parts := strings.SplitN(s, ".", 2) // Split at the first ".", resulting in 2 parts
		key := s                           // Default key is the full string
		if len(parts) > 0 {
			key = parts[0] // Use the first part of the split as the key
		}
		m[key] = s // Map the key to the original string
	}
	return m
}


// it should take a slice of string and return a map[string]string
func mapFromSlice(slice []string) map[string]string {
	m := make(map[string]string)
	for _, s := range slice {
		m[s] = s
	}
	return m
}