package integration_test

import (
	. "github.com/loveandpeople/lp.go/api"
	. "github.com/loveandpeople/lp.go/api/integration/samples"
	. "github.com/loveandpeople/lp.go/consts"
	. "github.com/loveandpeople/lp.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"strings"
)

var _ = Describe("GetTrytes()", func() {

	api, err := ComposeAPI(HTTPClientSettings{}, nil)
	if err != nil {
		panic(err)
	}

	Context("call", func() {
		It("resolves to correct response", func() {
			trytes, err := api.GetTrytes(DefaultHashes()...)
			Expect(err).ToNot(HaveOccurred())
			Expect(trytes).To(Equal([]Trytes{
				strings.Repeat("9", TransactionTrytesSize),
				strings.Repeat("9", TransactionTrytesSize),
			}))
		})
	})

	Context("invalid input", func() {
		It("returns an error for invalid transaction hashes", func() {
			_, err := api.GetTrytes("")
			Expect(errors.Cause(err)).To(Equal(ErrInvalidTransactionHash))
		})
	})

})
