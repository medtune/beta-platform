// +build cicd

package capsul

// This file is compiled only when build is taged cicd
// example: go build -tags=cicd cmd/main.go
// the reason is that gocv library for golang is built by
// binding opencv3 c++ library. Until now we havent
// implemented gocv build on dockerfiles.

// Convert images bytes to float32 array
// return an empty array
func bytesToFloat32(ib []byte) ([]float32, error) {
	return []float32{}, nil
}
