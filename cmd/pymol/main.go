package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("pymol", "-p", "-K", "cif/Avr2.cif")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		// Customize the cartoon
		io.WriteString(stdin, "set cartoon_color, 0x4C62D2, %fwk\n")
		io.WriteString(stdin, "set cartoon_transparency, 0.6, %fwk\n")
		io.WriteString(stdin, "set cartoon_transparency, 0.8, not %fwk\n")
		
		// Set the lighting
		io.WriteString(stdin, "bg white\n")
		io.WriteString(stdin, "set ambient, 0.05\n")
		io.WriteString(stdin, "set direct, 0.2\n")
		io.WriteString(stdin, "set spec_direct, 0\n")
		io.WriteString(stdin, "set shininess, 10.\n")
		io.WriteString(stdin, "set reflect, 0.5\n")
		io.WriteString(stdin, "set spec_count, -1\n")
		io.WriteString(stdin, "set spec_reflect, -1.\n")
		io.WriteString(stdin, "set specular, 1\n")
		io.WriteString(stdin, "set specular_intensity, 0.5\n")

		// Select and modify Q49,C55,H187,N208
		io.WriteString(stdin, "select Q, id 379-387\n")
		io.WriteString(stdin, "color blue, Q\n")
		io.WriteString(stdin, "show sticks, Q\n")

// pm(f'select C, id 419-424')
// pm('color red, C')
// pm('show sticks, C')

// pm(f'select H, id 1419-1428')
// pm('color blue, H')
// pm('show sticks, H')

// pm(f'select N, id 1587-1594')
// pm('color blue, N')
// pm('show sticks, N')

// # Select Chain B, change color
// pm(f'select B, chain B')
// pm('color red, B')

		
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
