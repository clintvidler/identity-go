package util

import (
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

// Returns true if the slice contains the element
func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func GetCookie(md metadata.MD, name string) (string, error) {
	if md["grpcgateway-cookie"] == nil {
		return "", errors.New("no cookies in the metadata")
	}

	for _, e := range strings.Split(md.Get("grpcgateway-cookie")[0], "; ") {
		parts := strings.Split(e, "=")
		if parts[0] == name {
			return parts[1], nil
		}
	}

	return "", errors.New("cookie not found")
}
