// Code generated by "go-option -type ResolveAddr"; DO NOT EDIT.

package resolver

// A ResolveAddrOption sets options.
type ResolveAddrOption interface {
	apply(*ResolveAddr)
}

// EmptyResolveAddrOption does not alter the configuration. It can be embedded
// in another structure to build custom options.
//
// This API is EXPERIMENTAL.
type EmptyResolveAddrOption struct{}

func (EmptyResolveAddrOption) apply(*ResolveAddr) {}

// ResolveAddrOptionFunc wraps a function that modifies ResolveAddr into an
// implementation of the ResolveAddrOption interface.
type ResolveAddrOptionFunc func(*ResolveAddr)

func (f ResolveAddrOptionFunc) apply(do *ResolveAddr) {
	f(do)
}

// sample code for option, default for nothing to change
func _ResolveAddrOptionWithDefault() ResolveAddrOption {
	return ResolveAddrOptionFunc(func(*ResolveAddr) {
		// nothing to change
	})
}
func (o *ResolveAddr) ApplyOptions(options ...ResolveAddrOption) *ResolveAddr {
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt.apply(o)
	}
	return o
}