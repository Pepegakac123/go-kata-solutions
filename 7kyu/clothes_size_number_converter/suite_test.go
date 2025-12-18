package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClothesSizeNumberConverter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ClothesSizeNumberConverter Suite")
}
