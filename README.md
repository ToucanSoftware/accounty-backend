![Build Status](https://github.com/ToucanSoftware/accounty-backend/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/ToucanSoftware/accounty-backend)](https://goreportcard.com/report/github.com/ToucanSoftware/accounty-backend) [![GoDoc](https://godoc.org/github.com/ToucanSoftware/accounty-backend?status.svg)](https://godoc.org/github.com/ToucanSoftware/accounty-backend)

# accounty-backend

Account Backend Services

## Add Certificates

```bash
mkdir -p cert
openssl genrsa -out cert/server.key 2048
openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
openssl req -new -sha256 -key cert/server.key -out cert/server.csr
openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
```