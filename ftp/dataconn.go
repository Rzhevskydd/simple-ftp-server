package ftp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (c *Conn) dataConnect() (net.Conn, error) {
	addr := c.dataPort.toAddress()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
	host 		   string
	port 		   int
}

func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d",
		&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2)
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

func dataPortFromEprtHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	delim := string(hostPort[0:1])
	parts := strings.Split(hostPort, delim)
	addressFamily, err := strconv.Atoi(parts[1])
	host := parts[2]
	port, err := strconv.Atoi(parts[3])

	if addressFamily != 1 && addressFamily != 2 {
		return nil, err
	}

	dp.host = host
	dp.port = port

	return &dp, nil
}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}

	if d.p1 != 0 && d.p2 != 0 {
		port := d.p1<<8 + d.p2
		return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
	}

	return net.JoinHostPort(d.host, strconv.Itoa(d.port))
}
