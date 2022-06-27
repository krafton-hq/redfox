package external_dns_con

import (
	"bufio"
	"encoding/gob"
	"net"

	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/services/external_dns_service"
	"go.uber.org/zap"
)

type Controller struct {
	service *external_dns_service.Service
}

func NewController(service *external_dns_service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Start(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			zap.S().Infow("External-Dns Tcp Accept Error", "error", err)
			return err
		}
		go func(inConn net.Conn) {
			defer inConn.Close()
			zap.S().Infow("External-Dns Tcp Connection Accepted", "client", inConn.RemoteAddr().String())

			err = c.handle(inConn)
			if err != nil {
				zap.S().Infow("External-Dns Tcp Server Run Error", "error", err, "client", inConn.RemoteAddr().String())
				return
			}
		}(conn)
	}
}

func (c *Controller) handle(conn net.Conn) error {
	records, err := c.service.Records()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(conn)
	enc := gob.NewEncoder(w)
	if err = enc.Encode(records); err != nil {
		return errors.WrapError(err, "External-Dns Tcp error while writing and encoding records")
	}
	if err = w.Flush(); err != nil {
		return errors.WrapError(err, "External-Dns Tcp error while flush buffer")
	}
	return nil
}
