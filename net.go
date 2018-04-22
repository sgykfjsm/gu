package gu

import "net"

func FindLocalIPv4Address() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var ipAddress string
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch i := addr.(type) {
			case *net.IPNet:
				ip = i.IP
			case *net.IPAddr:
				ip = i.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ipV4 := ip.To4()
			if ipV4 == nil {
				continue
			}
			ipAddress = ipV4.String()
			break
		}
	}

	return ipAddress, nil
}
