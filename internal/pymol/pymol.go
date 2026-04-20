// The pymol package supports control of the pymol protein structure
// viewer from go. In main:
//	cmd := exec.Command("pymol", "-p", "-K", cif)
//	stdin, err := cmd.StdinPipe()
//
// stdin communicates with pymol, and is as an argument of type io.Writer
// in functions in the pymol package.
//
// This package was inspired by the python package 'pymolPy3': 
//     https://github.com/carbonscott/pymolPy3/tree/main
package pymol

import (
	"fmt"
	"io"
	"strings"
)

func CustomizeCartoon(r io.Writer) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("set cartoon_transparency, 0.2\n"))
	builder.WriteString(fmt.Sprintf("set cartoon_highlight_color, grey85\n"))
	io.WriteString(r, builder.String())
	return
}

func SetLighting(r io.Writer) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("bg white\n"))
	builder.WriteString(fmt.Sprintf("set ambient, 0.05\n"))
	builder.WriteString(fmt.Sprintf("set direct, 0.2\n"))
	builder.WriteString(fmt.Sprintf("set spec_direct, 0\n"))
	builder.WriteString(fmt.Sprintf("set shininess, 10.\n"))
	builder.WriteString(fmt.Sprintf("set reflect, 0.5\n"))
	builder.WriteString(fmt.Sprintf("set spec_count, -1\n"))
	builder.WriteString(fmt.Sprintf("set spec_reflect, -1.\n"))
	builder.WriteString(fmt.Sprintf("set specular, 1\n"))
	builder.WriteString(fmt.Sprintf("set specular_intensity, 0.5\n"))
	io.WriteString(r, builder.String())
	return
}

// Makes a pymol selection based on start and end atom id, sets the selection color,
// and optionally shows sticks.
func SelectByID(r io.Writer, name string, color string, idstart int, idend int, showsticks bool) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("select %s, id %d-%d\n", name, idstart, idend))
	builder.WriteString(fmt.Sprintf("color %s, %s\n", color, name))

	if showsticks {
		builder.WriteString(fmt.Sprintf("show sticks, %s\n", name))
		builder.WriteString(fmt.Sprintf("util.cnc %s\n", name))		
	}
	io.WriteString(r, builder.String())
	return
}

// Makes a pymol selection based on chain, sets the selection color, and
// optionally shows sticks.
func SelectByChain(r io.Writer, name string, color string, chain string, showsticks bool) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("select %s, chain %s\n", name, chain))
	builder.WriteString(fmt.Sprintf("color %s, %s\n", color, name))
	if showsticks {
		builder.WriteString(fmt.Sprintf("show sticks, %s\n", name))
	}
	io.WriteString(r, builder.String())
	return
}
