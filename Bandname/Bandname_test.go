package Bandname_test

import (
	. "codeVar/Bandname"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should return the correct values", func() {
		Expect(BandNameGenerator(string("knife"))).Should(BeEquivalentTo("The Knife"))
		Expect(BandNameGenerator(string("tart"))).Should(BeEquivalentTo("Tartart"))
		Expect(BandNameGenerator(string("sandles"))).Should(BeEquivalentTo("Sandlesandles"))
		Expect(BandNameGenerator(string("bed"))).Should(BeEquivalentTo("The Bed"))
	})
})
