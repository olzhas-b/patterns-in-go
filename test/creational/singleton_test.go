package creational

import (
	"patterns-in-go/creational-patterns/singleton"
	"testing"
	"time"
)

func TestSingleTon(t *testing.T) {
	var (
		initialDNS  = "postgres"
		oracle      = "oracle"
		expectedDNS = "postgres"
	)

	for i := 0; i < 300000; i++ {
		go singleton.GetDatabaseConn(initialDNS)
	}
	time.Sleep(time.Second)

	got := singleton.GetDatabaseConn(oracle)
	if got.DNS != expectedDNS {
		t.Fatalf("expected %s but got %s", expectedDNS, got.DNS)
	}
}
