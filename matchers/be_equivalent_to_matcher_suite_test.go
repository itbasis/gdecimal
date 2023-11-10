package matchers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
)

func TestDecimalBeEquivalentToMatcher(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "DecimalBeEquivalentToMatcher Suite")
}
