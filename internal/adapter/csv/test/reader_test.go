package csv

import (
	basic_csv "encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/adapter/csv"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/stretchr/testify/require"
)

func TestReaderControllerGetFileName(t *testing.T) {
	fileName := "data.csv"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	c := csv.NewReader(f)
	require.EqualValues(t, fileName, c.GetFileName())
}

func TestReaderAdapterReadFirstLine(t *testing.T) {
	fileName := "data.csv"
	//first reader
	f, err := os.Open(fileName)
	if err != nil {

		log.Fatal(err)
	}
	defer f.Close()

	//second reader
	f2, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	c := csv.NewReader(f)
	reader1, err := c.Read()
	require.NoError(t, err)

	r := basic_csv.NewReader(f2)
	reader2, err := r.Read()
	require.NoError(t, err)
	require.EqualValues(t, reader1, reader2)
}

func TestReaderAdapterReadAll(t *testing.T) {
	fileName := "data.csv"
	//first reader
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//second reader
	f2, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	c := csv.NewReader(f)
	countFirstReader := 0
	for {
		_, err := c.Read()
		if err != nil {
			require.Error(t, err, port.ErrEOF)
			break
		}
		countFirstReader++
	}

	r := basic_csv.NewReader(f2)
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
