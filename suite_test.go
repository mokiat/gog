package gog_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gog Suite")
}
