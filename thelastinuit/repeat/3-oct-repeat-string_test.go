package repeat_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/repeat"
)
var _ = Describe("Example Tests", func() {
  It("should repeat correctly", func() {
	Expect(RepeatStr(4, "a")).To(Equal("aaaa"))
	Expect(RepeatStr(3, "hello ")).To(Equal("hello hello hello "))
	Expect(RepeatStr(2, "abc")).To(Equal("abcabc"))
  })
})
