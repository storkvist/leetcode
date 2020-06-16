package main

import (
	"regexp"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	IP = strings.ToLower(IP)
	ipv6Parts := strings.Split(IP, ":")
	if len(ipv6Parts) == 8 {
		return validIPv6(ipv6Parts)
	}
	ipv4Parts := strings.Split(IP, ".")
	if len(ipv4Parts) == 4 {
		return validIPv4(ipv4Parts)
	}
	return "Neither"
}

func validIPv4(parts []string) string {
	ok := true
	for _, part := range parts {
		if len(part) == 0 {
			return "Neither"
		}
		ipart, _ := strconv.Atoi(part)
		ok = (ipart >= 0) && (ipart <= 255) && (strconv.FormatInt(int64(ipart), 10) == part)
		if !ok {
			return "Neither"
		}
	}
	return "IPv4"
}

func validIPv6(parts []string) string {
	validPart := regexp.MustCompile(`^[a-f0-9]{1,4}$`)
	for _, part := range parts {
		if !validPart.MatchString(part) {
			return "Neither"
		}
	}
	return "IPv6"
}
