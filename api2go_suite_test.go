package api2go

import (
	"io"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApi2go(t *testing.T) {
	RegisterFailHandler(Fail)
	log.SetOutput(io.Discard)
	RunSpecs(t, "Api2go Suite")
}
