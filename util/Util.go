package util

import (
	"errors"
	"net"
)

func ConvertIPToDecimal(ipAddress string) (int64, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0, errors.New("Geçersiz IP adresi")
	}

	ipBytes := ip.To4()

	decimalIP := int64(ipBytes[0])<<24 +
		int64(ipBytes[1])<<16 +
		int64(ipBytes[2])<<8 +
		int64(ipBytes[3])

	return decimalIP, nil
}

// burası birden fazla ipadresi dondürüyor şimdilik ben [0] alıcam
func ResolveDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ipAddresses []string
	for _, ip := range ips {
		ipAddresses = append(ipAddresses, ip.String())
	}

	return ipAddresses, nil
}
