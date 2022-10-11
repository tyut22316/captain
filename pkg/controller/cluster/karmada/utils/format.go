package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	// Query the IP address of the current host accessing the Internet
	getInternetIPUrl = "https://myexternalip.com/raw"
	// A split symbol that receives multiple values from a command flag
	separator = ","
)

// StringToNetIP String To NetIP
func StringToNetIP(addr string) net.IP {
	if ip := net.ParseIP(addr); ip != nil {
		return ip
	}
	return net.ParseIP("127.0.0.1")
}

// FlagsIP Receive master external IP from command flags
func FlagsIP(ip string) []net.IP {
	var ips []net.IP

	arr := strings.Split(ip, separator)
	for _, v := range arr {
		ips = append(ips, StringToNetIP(v))
	}
	return ips
}

// FlagsDNS Receive master external DNS from command flags
func FlagsDNS(dns string) []string {
	return strings.Split(dns, separator)
}

// InternetIP Current host Internet IP.
func InternetIP() (net.IP, error) {
	resp, err := http.Get(getInternetIPUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return StringToNetIP(string(content)), nil
}

// MapToString  labels to string
func MapToString(labels map[string]string) string {
	v := new(bytes.Buffer)
	for key, value := range labels {
		fmt.Fprintf(v, "%s=%s,", key, value)
	}
	return strings.TrimRight(v.String(), ",")
}

// IsExist Determine whether the path exists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}