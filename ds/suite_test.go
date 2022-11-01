package ds_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Structures Suite")
}
