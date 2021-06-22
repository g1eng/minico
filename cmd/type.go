package cmd

import "github.com/g1eng/minico/functions"

type Minico struct {
	user      string
	container string
	display   string
	locale    string
	ime       string
	networks  []string
	volumes   []string
	provider  functions.Provider
}
