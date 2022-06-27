package external_dns_con

import (
	"net"
	"testing"
	"time"

	log_helper "github.com/krafton-hq/golib/log-helper"
	"github.com/krafton-hq/red-fox/server/services/external_dns_service"
)

func TestNewController(t *testing.T) {
	log_helper.Initialize(true, false)

	service := external_dns_service.NewService("srv.sbx-central.io", "srv.sbx-central.io")

	con := NewController(service)
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		t.Fatal(err)
	}
	err = con.Start(ln)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1000 * time.Second)
}
