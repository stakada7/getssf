package getssf

import (
	"fmt"
	"runtime"
)

func PrintVersion() {
	fmt.Printf(`Getssf %s
Compiler: %s %s
Copyright (C) TribalMediaHouse, Inc.
`,
		Version,
		runtime.Compiler,
		runtime.Version())
}
