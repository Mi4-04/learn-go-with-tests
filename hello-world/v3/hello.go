package v3

const prefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return prefix + name
}
