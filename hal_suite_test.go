package hal_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HAL Suite")
}
