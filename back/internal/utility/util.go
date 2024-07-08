package utility

import (
	"errors"
	"net"
)

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func SplitProxy(addr string) (string, int, error) {
	if addr == "" {
		return "", 0, errors.New("Прокси не указан")
	}

	var (
		err        error
		host, port string
		portnum    int
	)

	if host, port, err = net.SplitHostPort(addr); err != nil {
		return "", 0, err
	}
	if portnum, err = net.LookupPort("tcp", port); err != nil {
		return "", 0, err
	}

	return host, portnum, nil
}

func FindFreePort() (int, error) {
	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
