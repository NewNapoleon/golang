// Copyright 2019 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmux_test

import (
	"context"
	"github.com/searKing/golang/go/net/cmux"
	"github.com/searKing/golang/go/testing/leakcheck"
	"io"
	"net"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

var (
	benchHTTP1Payload = make([]byte, 4096)
	benchHTTP2Payload = make([]byte, 4096)
)

func init() {
	copy(benchHTTP1Payload, []byte("GET http://www.w3.org/ HTTP/1.1"))
	copy(benchHTTP2Payload, http2.ClientPreface)
}

type mockConn struct {
	net.Conn
	r io.Reader
}

func (c *mockConn) Read(b []byte) (n int, err error) {
	return c.r.Read(b)
}

func (c *mockConn) SetReadDeadline(time.Time) error {
	return nil
}

func discard(ctx context.Context, l net.Listener) {
	for {
		select {
		case <-ctx.Done():
			return
		}
		c, err := l.Accept()
		if err != nil {
			return
		}
		c.Close()
	}
}

func run(ctx context.Context, m cmux.CMux, matchers ...cmux.Matcher) {
	for _, matcher := range matchers {
		l := m.Match(matcher)
		go func() {
			_ = m.Serve(l)
			defer l.Close()
		}()

		go discard(ctx, l)
	}
}

func BenchmarkCMuxConnHTTP1(b *testing.B) {
	defer leakcheck.Check(b)
	m := cmux.New(context.Background())
	defer m.Close()
	lis := testListener(b)
	defer lis.Close()
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			run(ctx, m, cmux.HTTP1Fast())
		}
	})
}

func BenchmarkCMuxConnHTTP2(b *testing.B) {
	defer leakcheck.Check(b)
	m := cmux.New(context.Background())
	defer m.Close()
	lis := testListener(b)
	defer lis.Close()
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			run(ctx, m, cmux.HTTP2())
		}
	})
}
func BenchmarkCMuxConnHTTP1n2(b *testing.B) {
	defer leakcheck.Check(b)
	m := cmux.New(context.Background())
	defer m.Close()
	lis := testListener(b)
	defer lis.Close()
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			run(ctx, m, cmux.HTTP1Fast())
			run(ctx, m, cmux.HTTP2())
		}
	})
}

func BenchmarkCMuxConnHTTP2n1(b *testing.B) {
	defer leakcheck.Check(b)
	m := cmux.New(context.Background())
	defer m.Close()
	lis := testListener(b)
	defer lis.Close()
	ctx, cancelFn := context.WithCancel(context.TODO())
	defer cancelFn()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			run(ctx, m, cmux.HTTP2())
			run(ctx, m, cmux.HTTP1Fast())
		}
	})
}