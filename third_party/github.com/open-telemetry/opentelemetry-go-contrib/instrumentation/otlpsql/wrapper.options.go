// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package otlpsql

import (
	"go.opentelemetry.io/otel/attribute"
)

const defaultInstanceName = "default"

// wrapper holds configuration of our sql tracing middleware.
// By default, all options are set to false intentionally when creating a wrapped
// driver and provide the most sensible default with both performance and
// security in mind.
//go:generate go-option -type=wrapper
type wrapper struct {
	// AllowRoot, if set to true, will allow ocsql to create root spans in
	// absence of existing spans or even context.
	// Default is to not trace ocsql calls if no existing parent span is found
	// in context or when using methods not taking context.
	AllowRoot bool

	// Ping, if set to true, will enable the creation of spans on Ping requests.
	Ping bool

	// RowsNext, if set to true, will enable the creation of spans on RowsNext
	// calls. This can result in many spans.
	RowsNext bool

	// RowsClose, if set to true, will enable the creation of spans on RowsClose
	// calls.
	RowsClose bool

	// RowsAffected, if set to true, will enable the creation of spans on
	// RowsAffected calls.
	RowsAffected bool

	// LastInsertID, if set to true, will enable the creation of spans on
	// LastInsertId calls.
	LastInsertID bool

	// Query, if set to true, will enable recording of sql queries in spans.
	// Only allow this if it is safe to have queries recorded with respect to
	// security.
	Query bool

	// QueryParams, if set to true, will enable recording of parameters used
	// with parametrized queries. Only allow this if it is safe to have
	// parameters recorded with respect to security.
	// This setting is a noop if the Query option is set to false.
	QueryParams bool

	// DefaultAttributes will be set to each span as default.
	DefaultAttributes []attribute.KeyValue

	// InstanceName identifies database.
	InstanceName string

	// DisableErrSkip, if set to true, will suppress driver.ErrSkip errors in spans.
	DisableErrSkip bool
}

// WithAllWrapperOptions enables all available trace options.
func WithAllWrapperOptions() WrapperOption {
	return WrapperOptionFunc(func(o *wrapper) {
		*o = AllWrapperOptions
	})
}

// AllWrapperOptions has all tracing options enabled.
var AllWrapperOptions = wrapper{
	AllowRoot:    true,
	Ping:         true,
	RowsNext:     true,
	RowsClose:    true,
	RowsAffected: true,
	LastInsertID: true,
	Query:        true,
	QueryParams:  true,
}

// WithOptions sets our ocsql tracing middleware options through a single
// WrapperOptions object.
func WithOptions(options wrapper) WrapperOption {
	return WrapperOptionFunc(func(o *wrapper) {
		*o = options
		o.DefaultAttributes = append(
			[]attribute.KeyValue{}, options.DefaultAttributes...,
		)
	})
}
