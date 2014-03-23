package gover

import (
	"testing"
)

var rtests = []struct {
	name        string
	in, version string
}{

	// Adding an import to an existing parenthesized import
	{
		name: "version_update",
		in: `package foo

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// The version number that is being used at the moment.
const Version = "0.1.0"

`,
		version: "0.1.0",
	},
}

var utests = []struct {
	name    string
	in, out string
}{

	// Adding an import to an existing parenthesized import
	{
		name: "version_update",
		in: `package foo

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// The version number that is being used at the moment.
const Version = "0.1.0"

`,
		out: `package foo

// The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// The version number that is being used at the moment.
const Version = "0.2.0"

`,
	},
}

func TestUpdateVersion(t *testing.T) {
	for _, tt := range rtests {
		version, err := ReadVersion(tt.name+".go", []byte(tt.in))
		if err != nil {
			t.Errorf("error on %q: %v", tt.name, err)
			continue
		}
		if version != tt.version {
			t.Errorf("ReadVersion returned %s, expected %s", version, tt.version)
		}
	}
}
