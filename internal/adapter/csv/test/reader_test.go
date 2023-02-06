package csv

import (
	basic_csv "encoding/csv"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/adapter/csv"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/stretchr/testify/require"
)

var in string = `id;code_postal
1;75000
foo;bar
2;75001
4;baz
123456;67000
98;57870
3;69125
45;69620
34;69380
26;69420
89;69490
87;69430
27;69690
67;69840
7;69500
654;69770
987;69620
98765;69410
89;69840
76543;69680`

func TestReaderControllerGetFileName(t *testing.T) {
	fileName := "data.csv"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	c := csv.NewReader(f, f.Name())
	require.EqualValues(t, fileName, c.GetFileName())
}

func TestReaderAdapterReadFirstLine(t *testing.T) {
	c := csv.NewReader(strings.NewReader(in), "")
	reader1, err := c.Read()
	require.NoError(t, err)

	r := basic_csv.NewReader(strings.NewReader(in))
	reader2, err := r.Read()
	require.NoError(t, err)
	require.EqualValues(t, reader1, reader2)
}

func TestReaderAdapterReadAll(t *testing.T) {
	c := csv.NewReader(strings.NewReader(in), "")
	countFirstReader := 0
	for {
		_, err := c.Read()
		if err != nil {
			require.Error(t, err, port.ErrEOF)
			break
		}
		countFirstReader++
	}

	r := basic_csv.NewReader(strings.NewReader(in))
	countSecondReader := 0
	for {
		_, err := r.Read()
		if err != nil {
			require.Error(t, err, port.ErrEOF)
			break
		}
		countSecondReader++
	}

	require.EqualValues(t, countFirstReader, countSecondReader)
}
