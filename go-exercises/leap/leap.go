// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

func RemainderOf(year, n int) bool {
    return year % n == 0
} 

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
    if RemainderOf(year, 400) {
        return true
    }
    if RemainderOf(year, 100) {
        return false
    }
    if RemainderOf(year, 4) {
        return true
    }
	return false
}
