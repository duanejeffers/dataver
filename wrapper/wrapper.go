package wrapper

import "io"

// Wrapper interface for using custom or specific format wrapper
type Wrapper interface {
	New(data []byte) (Wrapper, error)
	NewFromString(data string) (Wrapper, error)
	NewFromReader(reader io.Reader) (Wrapper, error)
}
