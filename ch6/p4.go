package main


// The containsX function's time complexity is O(N) since it must always iterate through the entire string.
func containsX(s string) (foundX bool){
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "X" {
			foundX = true
		}
	}

	return
}

// To improve the containsX function's efficiency for best- and average-case scenarios, we can return as soon as we find an "X" and avoid any unnecessary iterations.
func improvedContainsX(s string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "X" {
			return true
		}
	}

	return false
}

func main() {
	str := "abXcdefg"
	containsX(str)
	improvedContainsX(str)
}