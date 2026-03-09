package localblast

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/tnaums/gobio/internal/protein"
)

func LocalBlast(query protein.Protein) BlastOutput {

	f, err := os.Create("blastdb/query.fa")
	if err != nil {
		fmt.Printf("unable to create file: %s", err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s", query))

	params := lBlast{
		cmd:    "blastp",
		query:  "blastdb/query.fa",
		db:     "blastdb/Fusgr2.aa.fasta",
		outfmt: "5",
		out:    "-",
	}

	cmd := exec.Command(params.cmd, "-query", params.query, "-db", params.db, "-outfmt", params.outfmt, "-out", params.out)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	var blastoutput BlastOutput
	if err := xml.Unmarshal(out, &blastoutput); err != nil {
		log.Fatalf("XML parsing failed: %v", err)
	}

	return blastoutput

}
