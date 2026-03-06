package signalp

import (
	"bufio"
	//	"fmt"
	"io"
	"strings"
	"strconv"
)

// Holds output of SignalP 6 analysis which predicts
// probability of protein secretion
type SignalP struct {
	NnCutPos  int
	NnVote    int
	HmmCutPos int
	HmmProb   float64
}

// Holds signalp info for a proteome
type SignalPMap map[int]SignalP


// Parses a *_SigP.tab file from JGI Mycocosm fungal proteome and
// writes secreted protein info into a SignalPMap
func NewSignalPMap(r io.Reader) (SignalPMap, error) {
	scanner := bufio.NewScanner(r)
	retMap := SignalPMap{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		words := strings.Fields(text)
		myKey, _ := strconv.Atoi(words[0])
		nnCutPos, _ := strconv.Atoi(words[1])
		nnVote, _ := strconv.Atoi(words[2])
		hmmCutPos, _ := strconv.Atoi(words[3])
		hmmProb, _ := strconv.ParseFloat(words[4], 64)
		
		retMap[myKey] = SignalP{nnCutPos, nnVote, hmmCutPos, hmmProb}
	}
	return retMap, nil
}
