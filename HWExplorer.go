package hwarchiver

import (
	"bufio"
)

// HWExplorer is a struct to help us go trough files. We can scan the line
// but we only clean it from the buffer once we consume it
type HWExplorer struct {
	scanner *bufio.Scanner
	state   bool
	line    int
}

// newHWExplorer returns a new empty HWExplorer
func newHWExplorer(scanner *bufio.Scanner) *HWExplorer {
	return &HWExplorer{
		scanner: scanner,
		state:   false,
		line:    0,
	}
}

// scan conditionally advances the scanner if the current line has been consumed.
func (exp *HWExplorer) scan() bool {
	if !exp.state {
		if exp.scanner.Scan() {
			exp.state = true
			return true
		}
		return false
	}
	return true
}

// peek returns the current line without advancing.
func (exp *HWExplorer) peek() string {
	if exp.state {
		return exp.scanner.Text()
	}
	return ""
}

// consume marks the current line as used and returns it.
func (exp *HWExplorer) consume() string {
	exp.state = false
	exp.line++
	return exp.scanner.Text()
}
