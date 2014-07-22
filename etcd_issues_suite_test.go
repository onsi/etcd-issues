package etcd_issues_test

import (
	"github.com/cloudfoundry/storeadapter/storerunner/etcdstorerunner"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"

	"testing"
)

var etcdRunner *etcdstorerunner.ETCDClusterRunner

func TestEtcdIssues(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EtcdIssues Suite")
}

var _ = BeforeSuite(func() {
	etcdPort := 5000 + (config.GinkgoConfig.ParallelNode-1)*10
	etcdRunner = etcdstorerunner.NewETCDClusterRunner(etcdPort, 1)
})

var _ = BeforeEach(func() {
	etcdRunner.Start()
})

var _ = AfterEach(func() {
	etcdRunner.Stop()
})

var _ = AfterSuite(func() {
	etcdRunner.Stop()
})
