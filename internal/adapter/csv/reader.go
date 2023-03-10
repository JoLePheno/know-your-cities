package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"

	"github.com/JoLePheno/know-your-cities/internal/port"
)

var _ port.Reader = (*CSVAdapter)(nil)

type CSVAdapter struct {
	reader   *csv.Reader
	fileName string
}

func NewReader(r io.Reader, fileName string) *CSVAdapter {
	return &CSVAdapter{
		reader:   csv.NewReader(r),
		fileName: fileName,
	}
}

func (c *CSVAdapter) Read() ([]string, error) {
	data, err := c.reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, port.ErrEOF
		}
		return nil, err
	}

	return data, nil
}

func (c *CSVAdapter) LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			if len(buf) != 0 {
				count++
			}
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func (c *CSVAdapter) GetFileName() string {
	return c.fileName
}
