package main

import (
	"fmt"
	//"io"
	"log"
	"os/exec"

	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	cmd := exec.Command("pymol", "-p", "-K", "cif/9172_0.cif")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		pymol.CustomizeCartoon(stdin)
		pymol.SetLighting(stdin)

		// Select and modify Q49,C55,H187,N208
		pymol.SelectByID(stdin, "Q", "blue", 379, 387, true)
		pymol.SelectByID(stdin, "C", "red", 419, 424, true)
		pymol.SelectByID(stdin, "H", "blue", 1419, 1428, true)
		pymol.SelectByID(stdin, "N", "blue", 1587, 1594, true)

		// # Select Chain B, change color
		pymol.SelectByChain(stdin, "B", "red", "B", false)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
