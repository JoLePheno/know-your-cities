package hook

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var givenZipCode = map[string]bool{
	"75010":  true,
	"67130":  true,
	"1233":   false,
	"123456": false,
	"11111":  true,
	"123A1":  false,
	"":       false,
}

var givenIDs = map[string]bool{
	"1":         true,
	"223":       true,
	"123456789": true,
	"1A11":      false,
	"":          false,
	"-1234":     false,
}

func TestZipCode(t *testing.T) {
	h := &HookAdapter{}
	for zip, expectedResp := range givenZipCode {
		require.EqualValues(t, expectedResp, h.CheckZipCode(zip))
	}
}

func TestID(t *testing.T) {
	h := &HookAdapter{}
	for id, expectedResp := range givenIDs {
		_, err := h.CheckId(id)
		if expectedResp {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}

	maxInt := strconv.FormatUint(math.MaxUint64, 10) + "1"
	_, err := h.CheckId(maxInt)
	require.Error(t, err)
}
