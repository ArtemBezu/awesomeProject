package history

type BrowserHistory struct {
	history      []string
	currentIndex int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (browser *BrowserHistory) Visit(url string) {
	if len(browser.history) > browser.currentIndex+1 {
		browser.history = browser.history[:browser.currentIndex+1]
	}
	browser.history = append(browser.history, url)
	browser.currentIndex++
}

func (browser *BrowserHistory) Back(steps int) string {
	if steps > browser.currentIndex {
		browser.currentIndex = 0
	} else {
		browser.currentIndex = browser.currentIndex - steps
	}
	return browser.history[browser.currentIndex]
}

func (browser *BrowserHistory) Forward(steps int) string {
	maxForward := len(browser.history) - browser.currentIndex - 1
	if steps > maxForward {
		browser.currentIndex += maxForward
	} else {
		browser.currentIndex += steps
	}
	return browser.history[browser.currentIndex]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
