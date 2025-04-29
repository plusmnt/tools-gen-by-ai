# Go IP Lookup Tool

This is a simple command-line tool written in Go that takes an IPv4 address as input, checks its validity, determines if it's a private IP, and retrieves public IP information using the `ipinfo.io` API.

---

## Features

- Validates if the input is a proper IPv4 address.
- Detects if the IP belongs to a private subnet.
- Queries public IP geolocation and organization data using `ipinfo.io`.
- Friendly error handling and usage hints.

---

## Requirements

- [Go](https://golang.org/dl/) 1.13 or newer
- Internet connection (for public IP lookups)

---

## Installation

1. Clone or download this repository.
2. Navigate to the directory containing the code.
3. Run the program using `go run`:

```bash
go run iplookup.go <IPv4_ADDRESS>
```

Alternatively, you can build the binary:

```bash
go build iplookup.go
./iplookup <IPv4_ADDRESS>
```

---

## Usage

### Syntax

```bash
go run iplookup.go <IPv4_ADDRESS>
```

### Examples

#### Valid Public IP

```bash
go run iplookup.go 8.8.8.8
```

**Output:**

```
IP: 8.8.8.8
City: Mountain View
Region: California
Country: US
Organization: AS15169 Google LLC
Hostname: dns.google
```

#### Private IP Address

```bash
go run iplookup.go 192.168.0.1
```

**Output:**

```
192.168.0.1 is a private IP address.
```

#### Invalid IP Address

```bash
go run iplookup.go abc.def.ghi.jkl
```

**Output:**

```
Invalid IPv4 address: abc.def.ghi.jkl
Usage: go run iplookup.go <IPv4_ADDRESS>
Example: go run iplookup.go 8.8.8.8
```

---

## Notes

- The tool uses [ipinfo.io](https://ipinfo.io/) for IP lookup. No API key is required for basic usage.
- Private IPs are not sent to the lookup service.
- Only IPv4 addresses are supported in this version.

---

## License

MIT License

---

Would you like me to generate a downloadable `README.md` file for you?
