package hwarchiver

import (
	"bufio"
)

type HWExplorer struct {
	scanner *bufio.Scanner
	state   bool
	line    int
}

func NewHWExplcorer(scanner *bufio.Scanner) *HWExplorer {
	return &HWExplorer{
		scanner: scanner,
		state:   false,
		line:    0,
	}
}

// Scan conditionally advances the scanner if the current line has been consumed.
func (exp *HWExplorer) Scan() bool {
	if !exp.state {
		if exp.scanner.Scan() {
			exp.state = true
			return true
		}
		return false
	}
	return true
}

// Peek returns the current line without advancing.
func (exp *HWExplorer) Peek() string {
	if exp.state {
		return exp.scanner.Text()
	}
	return ""
}

// Consume marks the current line as used and returns it.
func (exp *HWExplorer) Consume() string {
	exp.state = false
	exp.line++
	return exp.scanner.Text()
}
