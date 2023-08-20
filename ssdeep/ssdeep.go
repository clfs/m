// Package ssdeep implements the ssdeep fuzzy hashing algorithm.
package ssdeep

type Hash struct{}

// Write adds more data to the running hash.
// It never returns an error.
func (h *Hash) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h *Hash) Sum(b []byte) []byte {
	return nil
}

// Reset resets the Hash to its initial state.
func (h *Hash) Reset() {

}

// New returns a new Hash.
func New() Hash {
	return Hash{}
}

// MatchScore returns the matching score of two ssdeep hashes.
// The result is between 0 and 100 inclusive.
func MatchScore(a, b []byte) int {
	return 0
}
