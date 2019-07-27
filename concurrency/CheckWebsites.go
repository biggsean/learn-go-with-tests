package concurrency

// WebsiteChecker type
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks a list of urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for range urls {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}