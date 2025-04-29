package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

// Struct to parse the JSON response from ipinfo.io
type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Hostname string `json:"hostname"`
}

// isValidIPv4 checks if the given string is a valid IPv4 address
func isValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// isPrivateIPv4 checks if an IP is from a private subnet
func isPrivateIPv4(ip string) bool {
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, block := range privateBlocks {
		_, cidr, _ := net.ParseCIDR(block)
		if cidr.Contains(parsedIP) {
			return true
		}
	}
	return false
}

// getIPInfo fetches IP geolocation info from ipinfo.io
func getIPInfo(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get IP info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var info IPInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &info, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run iplookup.go <IPv4_ADDRESS>")
		fmt.Println("Example: go run iplookup.go 8.8.8.8")
		return
	}

	ip := os.Args[1]

	if !isValidIPv4(ip) {
		fmt.Printf("Invalid IPv4 address: %s\n", ip)
		fmt.Println("Usage: go run iplookup.go <IPv4_ADDRESS>")
		fmt.Println("Example: go run iplookup.go 8.8.8.8")
		return
	}

	if isPrivateIPv4(ip) {
		fmt.Printf("%s is a private IP address.\n", ip)
		return
	}

	info, err := getIPInfo(ip)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nOrganization: %s\nHostname: %s\n",
		info.IP, info.City, info.Region, info.Country, info.Org, info.Hostname)
}
