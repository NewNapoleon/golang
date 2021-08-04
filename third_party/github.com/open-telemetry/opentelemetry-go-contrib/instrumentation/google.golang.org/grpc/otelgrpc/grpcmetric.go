// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package otelgrpc

import (
	otelcontrib "go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel/attribute"
)

const (
	// GRPCStatusCodeKey is convention for numeric status code of a gRPC request.
	GRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")
	// GRPCTypeKey is convention for grpc type of a gRPC connection.
	GRPCTypeKey = attribute.Key("rpc.grpc.type")
)

var (
	// InstrumentationName is the name of this instrumentation package.
	InstrumentationName = "github.com/searKing/golang/third_party/github.com/open-telemetry/opentelemetry-go-contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	// InstrumentationVersion is the version of this instrumentation package.
	InstrumentationVersion = otelcontrib.SemVersion()

	// AttrsFilter is a filter before Report
	AttrsFilter = func(attrs ...attribute.KeyValue) []attribute.KeyValue {
		return attrs
	}
)