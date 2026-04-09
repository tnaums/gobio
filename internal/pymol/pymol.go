package pymol

import (
	"fmt"
	"io"
	"strings"
)

func CustomizeCartoon(r io.Writer) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("set cartoon_color, 0x4C62D2, %fwk\n"))
	builder.WriteString(fmt.Sprintf("set cartoon_transparency, 0.6, %fwk\n"))
	builder.WriteString(fmt.Sprintf("set cartoon_transparency, 0.8, not %fwk\n"))
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

func SelectByID(r io.Writer, name string, color string, idstart int, idend int, showsticks bool) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("select %s, id %d-%d\n", name, idstart, idend))
	builder.WriteString(fmt.Sprintf("color %s, %s\n", color, name))
	if showsticks {
		builder.WriteString(fmt.Sprintf("show sticks, %s\n", name))
	}
	io.WriteString(r, builder.String())
	return
}

func SelectByChain(r io.Writer, name string, color string, chain string, showsticks bool) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("select %s, chain %s\n", name, chain))
	builder.WriteString(fmt.Sprintf("color %s, %s\n", color, name))
	if showsticks {
		builder.WriteString(fmt.Sprintf("show sticks, %s", name))
	}
	io.WriteString(r, builder.String())
	return
}
