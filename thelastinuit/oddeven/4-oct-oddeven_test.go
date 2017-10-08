package oddeven_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/oddeven"
)
var _ = Describe("Test Example", func() {
  It("should test that the solution returns the correct value", func() {
	Expect(EvenOrOdd(1)).To(Equal("Odd"))
	Expect(EvenOrOdd(2)).To(Equal("Even"))
  })
})
