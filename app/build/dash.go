// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package build

import (
	"net/http"
	"strings"

	"appengine"
)

func handleFunc(path string, h http.HandlerFunc) {
	for _, d := range dashboards {
		http.HandleFunc(d.Prefix+path, h)
	}
}

// Dashboard describes a unique build dashboard.
type Dashboard struct {
	Name      string     // This dashboard's name (eg, "Go")
	Namespace string     // This dashboard's namespace (eg, "" (default), "Git")
	Prefix    string     // The path prefix (no trailing /)
	Packages  []*Package // The project's packages to build
}

// dashboardForRequest returns the appropriate dashboard for a given URL path.
func dashboardForRequest(r *http.Request) *Dashboard {
	for _, d := range dashboards[1:] {
		if d.Prefix == "" {
			panic("prefix can be empty only for the first dashboard")
		}
		if strings.HasPrefix(r.URL.Path, d.Prefix) {
			return d
		}
	}
	if dashboards[0].Prefix != "" {
		panic("prefix for the first dashboard should be empty")
	}
	return dashboards[0]
}

// Context returns a namespaced context for this dashboard, or panics if it
// fails to create a new context.
func (d *Dashboard) Context(c appengine.Context) appengine.Context {
	if d.Namespace == "" {
		return c
	}
	n, err := appengine.Namespace(c, d.Namespace)
	if err != nil {
		panic(err)
	}
	return n
}

// The currently known dashboards.
// The first one should have an empty prefix and
// the other ones a non empty prefix.
var dashboards = []*Dashboard{goDash, gccgoDash}

// goDash is the dashboard for the main go repository.
var goDash = &Dashboard{
	Name:      "Go",
	Namespace: "Git",
	Prefix:    "",
	Packages:  goPackages,
}

// goPackages is a list of all of the packages built by the main go repository.
var goPackages = []*Package{
	{
		Kind: "go",
		Name: "Go",
	},
	{
		Kind: "subrepo",
		Name: "arch",
		Path: "golang.org/x/arch",
	},
	{
		Kind: "subrepo",
		Name: "blog",
		Path: "golang.org/x/blog",
	},
	{
		Kind: "subrepo",
		Name: "crypto",
		Path: "golang.org/x/crypto",
	},
	{
		Kind: "subrepo",
		Name: "exp",
		Path: "golang.org/x/exp",
	},
	{
		Kind: "subrepo",
		Name: "image",
		Path: "golang.org/x/image",
	},
	{
		Kind: "subrepo",
		Name: "mobile",
		Path: "golang.org/x/mobile",
	},
	{
		Kind: "subrepo",
		Name: "net",
		Path: "golang.org/x/net",
	},
	{
		Kind: "subrepo",
		Name: "oauth2",
		Path: "golang.org/x/oauth2",
	},
	{
		Kind: "subrepo",
		Name: "review",
		Path: "golang.org/x/review",
	},
	{
		Kind: "subrepo",
		Name: "sync",
		Path: "golang.org/x/sync",
	},
	{
		Kind: "subrepo",
		Name: "sys",
		Path: "golang.org/x/sys",
	},
	{
		Kind: "subrepo",
		Name: "talks",
		Path: "golang.org/x/talks",
	},
	{
		Kind: "subrepo",
		Name: "term",
		Path: "golang.org/x/term",
	},
	{
		Kind: "subrepo",
		Name: "text",
		Path: "golang.org/x/text",
	},
	{
		Kind: "subrepo",
		Name: "time",
		Path: "golang.org/x/time",
	},
	{
		Kind: "subrepo",
		Name: "tools",
		Path: "golang.org/x/tools",
	},
	{
		Kind: "subrepo",
		Name: "tour",
		Path: "golang.org/x/tour",
	},
}

// gccgoDash is the dashboard for gccgo.
var gccgoDash = &Dashboard{
	Name:      "Gccgo",
	Namespace: "Gccgo",
	Prefix:    "/gccgo",
	Packages: []*Package{
		{
			Kind: "gccgo",
			Name: "Gccgo",
		},
	},
}
