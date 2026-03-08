package localblast

import "encoding/xml"

type lBlast struct {
	cmd    string
	query  string
	db     string
	outfmt string
	out    string
}

// BlastOutput was generated 2026-03-07 16:18:48 by https://xml-to-go.github.io/ in Ukraine.
type BlastOutput struct {
	XMLName              xml.Name `xml:"BlastOutput"`
	Text                 string   `xml:",chardata"`
	BlastOutputProgram   string   `xml:"BlastOutput_program"`
	BlastOutputVersion   string   `xml:"BlastOutput_version"`
	BlastOutputReference string   `xml:"BlastOutput_reference"`
	BlastOutputDb        string   `xml:"BlastOutput_db"`
	BlastOutputQueryID   string   `xml:"BlastOutput_query-ID"`
	BlastOutputQueryDef  string   `xml:"BlastOutput_query-def"`
	BlastOutputQueryLen  string   `xml:"BlastOutput_query-len"`
	BlastOutputParam     struct {
		Text       string `xml:",chardata"`
		Parameters struct {
			Text                string `xml:",chardata"`
			ParametersMatrix    string `xml:"Parameters_matrix"`
			ParametersExpect    string `xml:"Parameters_expect"`
			ParametersGapOpen   string `xml:"Parameters_gap-open"`
			ParametersGapExtend string `xml:"Parameters_gap-extend"`
			ParametersFilter    string `xml:"Parameters_filter"`
		} `xml:"Parameters"`
	} `xml:"BlastOutput_param"`
	BlastOutputIterations struct {
		Text      string `xml:",chardata"`
		Iteration struct {
			Text              string `xml:",chardata"`
			IterationIterNum  string `xml:"Iteration_iter-num"`
			IterationQueryID  string `xml:"Iteration_query-ID"`
			IterationQueryDef string `xml:"Iteration_query-def"`
			IterationQueryLen string `xml:"Iteration_query-len"`
			IterationHits     struct {
				Text string `xml:",chardata"`
				Hit  []struct {
					Text         string `xml:",chardata"`
					HitNum       string `xml:"Hit_num"`
					HitID        string `xml:"Hit_id"`
					HitDef       string `xml:"Hit_def"`
					HitAccession string `xml:"Hit_accession"`
					HitLen       string `xml:"Hit_len"`
					HitHsps      struct {
						Text string `xml:",chardata"`
						Hsp  []struct {
							Text          string `xml:",chardata"`
							HspNum        string `xml:"Hsp_num"`
							HspBitScore   string `xml:"Hsp_bit-score"`
							HspScore      string `xml:"Hsp_score"`
							HspEvalue     string `xml:"Hsp_evalue"`
							HspQueryFrom  string `xml:"Hsp_query-from"`
							HspQueryTo    string `xml:"Hsp_query-to"`
							HspHitFrom    string `xml:"Hsp_hit-from"`
							HspHitTo      string `xml:"Hsp_hit-to"`
							HspQueryFrame string `xml:"Hsp_query-frame"`
							HspHitFrame   string `xml:"Hsp_hit-frame"`
							HspIdentity   string `xml:"Hsp_identity"`
							HspPositive   string `xml:"Hsp_positive"`
							HspGaps       string `xml:"Hsp_gaps"`
							HspAlignLen   string `xml:"Hsp_align-len"`
							HspQseq       string `xml:"Hsp_qseq"`
							HspHseq       string `xml:"Hsp_hseq"`
							HspMidline    string `xml:"Hsp_midline"`
						} `xml:"Hsp"`
					} `xml:"Hit_hsps"`
				} `xml:"Hit"`
			} `xml:"Iteration_hits"`
			IterationStat struct {
				Text       string `xml:",chardata"`
				Statistics struct {
					Text               string `xml:",chardata"`
					StatisticsDbNum    string `xml:"Statistics_db-num"`
					StatisticsDbLen    string `xml:"Statistics_db-len"`
					StatisticsHspLen   string `xml:"Statistics_hsp-len"`
					StatisticsEffSpace string `xml:"Statistics_eff-space"`
					StatisticsKappa    string `xml:"Statistics_kappa"`
					StatisticsLambda   string `xml:"Statistics_lambda"`
					StatisticsEntropy  string `xml:"Statistics_entropy"`
				} `xml:"Statistics"`
			} `xml:"Iteration_stat"`
		} `xml:"Iteration"`
	} `xml:"BlastOutput_iterations"`
}
