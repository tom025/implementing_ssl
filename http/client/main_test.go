package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("http client", func() {
	var client string

	BeforeSuite(func() {
		var err error
		client, err = gexec.Build("github.com/tom025/http/client")
		Expect(err).ToNot(HaveOccurred())
	})

	It("returns body of request", func() {
		session, err := start(client)
		Expect(err).ToNot(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
		Expect(string(session.Out.Contents())).To(ContainSubstring("foo"))
	})

})

func start(execPath string, args ...string) (*gexec.Session, error) {
	return gexec.Start(
		exec.Command(execPath, args...),
		GinkgoWriter,
		GinkgoWriter,
	)
}
