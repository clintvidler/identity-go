package util

// Returns true if the slice contains the element
func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}
