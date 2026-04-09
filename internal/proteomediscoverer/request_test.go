package proteomediscoverer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestGetAccession(t *testing.T) {
	r, err := GetAccession("FALSE,High,Master Protein,XP_018742465.1,hypothetical protein FVEG_00370 [Fusarium verticillioides 7600],0,30.677,30,13,236,13,416,43.2,8.56,775.73,13,0,5125861470,High,1,")
	require.NoError(t, err)
	assert.Equal(t, "XP_018742465.1", r.Accession)
}
