package opt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOpt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optional Suite")
}
