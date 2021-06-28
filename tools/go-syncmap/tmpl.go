package main

// Arguments to format are:
//	TmplPackageRender
const tmplPackageCode = `
{{ $package_scope := . }}

// Code generated by "{{.GenerateToolArgs}} {{.GenerateToolArgs}}"; DO NOT EDIT.
package {{.PackageName}}

{{- if $package_scope.WithSyncMapMethod }}
import (
	"sync" // Used by sync.Map.

{{- range .MapRenders}}
{{- if .MapImport }}
	"{{.MapImport}}"
{{- end }}
{{- if .KeyTypeImport }}
	"{{.KeyTypeImport}}"
{{- end }}
{{- if .ValueTypeImport }}
	"{{.ValueTypeImport}}"
{{- end }}
{{- end }}
)

// Generate code that will fail if the constants change value.
{{- range .MapRenders}}
func _() {
	// An "cannot convert {{.MapTypeName}} literal (type {{.MapTypeName}}) to type sync.Map" compiler error signifies that the base type have changed.
	// Re-run the go-syncmap command to generate them again.
	_ = (sync.Map)({{.MapTypeName}}{})
}
{{- end }}

{{- range .MapRenders}}
{{- if .ValueTypeNilVal }}
var {{.ValueTypeNilVal}} = func() (val {{.ValueTypeName}}) { return }()
{{- end }}
{{- end }}

{{- range .MapRenders}}
// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *{{.MapTypeName}}) Load(key {{.KeyTypeName}}) ({{.ValueTypeName}}, bool) {
    value, ok := (*sync.Map)(m).Load(key)
    if value == nil {
    	return {{.ValueTypeNilVal}}, ok
    }
    return value.({{.ValueTypeName}}), ok
}

// Store sets the value for a key.
func (m *{{.MapTypeName}}) Store(key {{.KeyTypeName}}, value {{.ValueTypeName}}) {
    (*sync.Map)(m).Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *{{.MapTypeName}}) LoadOrStore(key {{.KeyTypeName}}, value {{.ValueTypeName}}) ({{.ValueTypeName}}, bool) {
    actual, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if actual == nil {
        return {{.ValueTypeNilVal}}, loaded
    }
    return actual.({{.ValueTypeName}}), loaded
}

{{- if $package_scope.WithMethodLoadAndDelete }}
// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *{{.MapTypeName}}) LoadAndDelete(key {{.KeyTypeName}}) (value {{.ValueTypeName}}, loaded bool) {
	actual, loaded := (*sync.Map)(m).LoadAndDelete(key)
	if actual == nil {
        return {{.ValueTypeNilVal}}, loaded
    }
    return actual.({{.ValueTypeName}}), loaded
}
{{- else }}
// LoadAndDelete will not be generated for Go version 1.14 and earlier version.
// https://github.com/golang/go/issues/33762
{{- end }}

// Delete deletes the value for a key.
func (m *{{.MapTypeName}}) Delete(key {{.KeyTypeName}}) {
    (*sync.Map)(m).Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *{{.MapTypeName}}) Range(f func(key {{.KeyTypeName}}, value {{.ValueTypeName}}) bool) {
    (*sync.Map)(m).Range(func(key, value interface{}) bool {
        return f(key.({{.KeyTypeName}}), value.({{.ValueTypeName}}))
    })
}
{{- end }}
{{- else }}
// sync.Map will not be generated for Go version 1.8 and earlier version.
// https://github.com/golang/go/issues/18177
{{- end }}
`
