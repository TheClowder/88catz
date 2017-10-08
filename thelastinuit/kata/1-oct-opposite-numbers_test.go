package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("opposite function", func() {
  It("should invert number", func() {
	Expect(Opposite(1)).To(Equal(-1))
  })
})
