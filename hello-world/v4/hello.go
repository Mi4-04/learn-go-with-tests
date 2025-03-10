package v4

const (
	Spanish = "Spanish"
	French  = "French"
	English = "English"
)

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case English:
		prefix = englishPrefix
	case Spanish:
		prefix = spanishPrefix
	case French:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
