package concurrency

type Checker func(url string) bool

type Status struct {
	Address string
	IsUp    bool
}

// Receives a slice of website addresses then check if the passed website are up or down
// Returns a map that matches every website address to its status.
func CheckWebsites(checker Checker, websites []string) map[string]bool {
	statusChannel := make(chan Status)
	statuses := make(map[string]bool)
	for _, website := range websites {
		go func(url string) {
			statusChannel <- Status{Address: url, IsUp: checker(url)}
		}(website)

	}
	for i := 0; i < len(websites); i++ {
		websiteStatus := <-statusChannel
		statuses[websiteStatus.Address] = websiteStatus.IsUp
	}

	return statuses
}
