/*
Sniperkit-Bot
- Status: analyzed
*/

// generated by amalgomate; DO NOT EDIT
package amalgomatedcheck

import (
	"fmt"
	"sort"

	importalias "github.com/sniperkit/snk.fork.palantir-godel-okgo-asset-importalias/generated_src/internal/github.com/palantir/go-importalias"
)

var programs = map[string]func(){"importalias": func() {
	importalias.AmalgomatedMain()
},
}

func Instance() Amalgomated {
	return &amalgomated{}
}

type Amalgomated interface {
	Run(cmd string)
	Cmds() []string
}

type amalgomated struct{}

func (a *amalgomated) Run(cmd string) {
	if _, ok := programs[cmd]; !ok {
		panic(fmt.Sprintf("Unknown command: \"%v\". Valid values: %v", cmd, a.Cmds()))
	}
	programs[cmd]()
}

func (a *amalgomated) Cmds() []string {
	var cmds []string
	for key := range programs {
		cmds = append(cmds, key)
	}
	sort.Strings(cmds)
	return cmds
}
