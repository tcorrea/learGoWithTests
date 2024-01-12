package concurrency

type WebsiteChecker func(string) bool

// The new type, result has been made to associate the return value of the WebsiteChecker
// with the url being checked - it's a struct of string and bool.
// As we don't need either value to be named, each of them is anonymous within the struct;
// this can be useful in when it's hard to know what to name a value.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// chan result is the type of the channel - a chennel of result
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
