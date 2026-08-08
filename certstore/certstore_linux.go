package certstore

import "errors"

// certstore isn't implemented on Linux. openStore returns an error so packages
// that import certstore will build on non-mac/windows platforms.
func openStore() (Store, error) {
	return nil, errors.New("certstore only works on macOS and Windows")
}
