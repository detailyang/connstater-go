// Package connstater defines the interface ConnectionStater.
package connstater

import (
	"crypto/tls"
)

// ConnectionStater indicates whether implementations the ConnectionState.
type ConnectionStater interface {
	ConnectionState() tls.ConnectionState
}
