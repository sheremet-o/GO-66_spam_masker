package masker

func Masker(message string) string {
	buffer := []byte(message)
	linkHttp := []byte("http://")

	for i := 0; i < len(buffer)-len(linkHttp); i++ {
		if string(buffer[i:i+len(linkHttp)]) == string(linkHttp) {
			j := i + len(linkHttp)
			for j < len(buffer) && buffer[j] != ' ' {
				buffer[j] = '*'
				j++
			}
			i = j
		}
	}
	return string(buffer)
}
