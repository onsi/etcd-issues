package etcd_issues_test

import (
	"time"

	"github.com/coreos/go-etcd/etcd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ETCD Issues", func() {
	var client *etcd.Client

	BeforeEach(func() {
		client = etcd.NewClient(etcdRunner.NodeURLS())
	})

	Describe("Watch-related issues", func() {
		It("should send back the correct error when a watch is stopped", func() {
			stop := make(chan bool, 0)
			responseChan := make(chan *etcd.Response, 0)
			errChan := make(chan error, 0)

			go func() {
				response, err := client.Watch("/", 0, true, nil, stop)
				responseChan <- response
				errChan <- err
			}()

			//give the goroutine a chance to run...
			time.Sleep(100 * time.Millisecond)

			close(stop)

			response := <-responseChan
			err := <-errChan

			Expect(response).To(BeNil())
			Expect(err).To(Equal(etcd.ErrWatchStoppedByUser))
		})
	})
})
