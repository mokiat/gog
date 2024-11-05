package seq_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSeq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sequence Suite")
}
