package localblast

import (
	"bufio"
	"fmt"
	"os"
)

func PrintBlastp(b BlastOutput) {
	fmt.Println(b.BlastOutputQueryDef)
	fmt.Println(b.BlastOutputQueryLen)

	for _, hit := range b.BlastOutputIterations.Iteration.IterationHits.Hit {
		fmt.Println()
		fmt.Printf("Hit %s\n", hit.HitNum)
		fmt.Println("Description:")
		//		fmt.Printf("    ID:%s", hit.HitID)
		fmt.Printf("   Def:%s", hit.HitDef)
		//		fmt.Printf("  Accession:%s", hit.HitAccession)
		fmt.Printf("  Length:%s\n", hit.HitLen)
		for _, hsp := range hit.HitHsps.Hsp {
			fmt.Printf("Hsp: %s\n", hsp.HspNum)
			fmt.Printf("    Bitscore=%s; ", hsp.HspBitScore)
			fmt.Printf("Score=%s; ", hsp.HspScore)
			fmt.Printf("Evalue=%s; ", hsp.HspEvalue)
			fmt.Printf("Identity=%s; ", hsp.HspIdentity)
			fmt.Printf("Posititve=%s; ", hsp.HspPositive)
			fmt.Printf("AlignLen=%s; ", hsp.HspAlignLen)
			fmt.Printf("Gaps:%s\n", hsp.HspGaps)
			fmt.Printf("    Query:%s -> %s\n", hsp.HspQueryFrom, hsp.HspQueryTo)
			fmt.Printf("      Hit:%s -> %s\n", hsp.HspHitFrom, hsp.HspHitTo)

			printAlignment(hsp.HspQseq, hsp.HspHseq, hsp.HspMidline)
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("press Enter to continue: ")
			for scanner.Scan() {
				scanner.Text()
				break
			}
			fmt.Print("\033[H\033[2J")  // clears screen and prints at top
			                            // probably only works on linux
		}
	}

	fmt.Println("----------------------------------------")
}

func printAlignment(q, h, m string) {
	query := ""
	hit := ""
	midline := ""
	for idx, x := range q {
		if idx == 0 {
			query += string(x)
			hit += string(h[0])
			midline += string(m[0])
			continue
		}
		if idx%90 == 0 {
			fmt.Println()
			fmt.Printf("    %s\n", query)
			fmt.Printf("    %s\n", midline)
			fmt.Printf("    %s\n", hit)

			query = string(x)
			hit = string(h[idx])
			midline = string(m[idx])
			continue
		}
		query += string(x)
		hit += string(h[idx])
		midline += string(m[idx])
	}
	fmt.Println()
	fmt.Printf("    %s\n", query)
	fmt.Printf("    %s\n", midline)
	fmt.Printf("    %s\n", hit)

}
