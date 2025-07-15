//go:build tools
// +build tools

package graphql

import (
	_ "github.com/99designs/gqlgen/cmd"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
