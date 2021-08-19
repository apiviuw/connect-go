package rerpc

import (
	"fmt"
	"runtime"
)

// Version is the semantic version of the reRPC module.
const Version = "0.0.1"

// MaxHeaderBytes is 8KiB, gRPC's recommended maximum header size. To enforce
// this limit, set MaxHeaderBytes on your http.Server.
const MaxHeaderBytes = 1024 * 8

// ReRPC's supported HTTP Content-Types. Servers decide whether to use the gRPC
// or Twirp protocol based on the request's Content-Type. See the protocol
// documentation at https://rerpc.github.io for more information.
const (
	TypeDefaultGRPC = "application/grpc"
	TypeProtoGRPC   = "application/grpc+proto"
	TypeProtoTwirp  = "application/protobuf"
	TypeJSON        = "application/json"
)

// ReRPC's supported compression methods.
const (
	CompressionIdentity = "identity"
	CompressionGzip     = "gzip"
)

// StreamType describes whether the client, server, neither, or both is
// streaming.
type StreamType uint8

const (
	StreamTypeClient        StreamType = 0b01
	StreamTypeServer                   = 0b10
	StreamTypeUnary                    = StreamTypeClient & StreamTypeServer
	StreamTypeBidirectional            = StreamTypeClient | StreamTypeServer
)

// These constants are used in compile-time handshakes with reRPC's generated
// code.
const (
	SupportsCodeGenV0 = iota
)

var userAgent = fmt.Sprintf("grpc-go-rerpc/%s (%s)", Version, runtime.Version())
