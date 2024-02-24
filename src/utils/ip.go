package utils

import "net"

func ips(isPrivate func(net.IP) bool) ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := []string{""}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.To4()
			case *net.IPAddr:
				ip = v.IP.To4()
			}

			if ip == nil {
				continue
			}

			if isPrivate(ip) {
				ips = append(ips, ip.String())
			}
		}
	}

	return ips, nil
}

func PrivateIPs() ([]string, error) {
	privateIPs, err := ips(func(ip net.IP) bool {
		return ip.IsPrivate()
	})

	if err != nil {
		return nil, err
	}

	for i, ip := range privateIPs {
		if ip == "" {
			privateIPs = append(privateIPs[:i], privateIPs[i+1:]...)
		}
	}

	return append([]string{"localhost", "127.0.0.1"}, privateIPs...), nil
}

func PublicIPs() ([]string, error) {
	return ips(func(ip net.IP) bool {
		return !ip.IsPrivate() && !ip.IsLoopback()
	})
}
