package sylphy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type Sylphy struct{}

func (s *Sylphy) Start(r io.Reader) error {
	// Has to be bufio.Reader to be able to read till the end of the line.
	buffReader := bufio.NewReader(r)

	decoder := json.NewDecoder(buffReader)

	for {
		entry := make(map[string]any)

		if err := decoder.Decode(&entry); err == io.EOF {
			// All readers are exhausted at this point.
			break
		} else if err != nil {
			// json.Decoder internally buffers 4kB read from the reader
			// so to process input need to aggregate remaining buffered data
			// with the primary reader.
			mr := io.MultiReader(decoder.Buffered(), buffReader)
			buffReader = bufio.NewReader(mr)

			line, err := buffReader.ReadString('\n')
			if err != nil {
				return err
			}

			fmt.Print(line)

			// Recreate decoder after processing the line that caused
			//   the previous one to fail. It doesn't mean that the next line
			//   will be parsed successfully, but in that case it will
			//   be handled in the same maner.
			decoder = json.NewDecoder(buffReader)

			continue
		}

		fmt.Println(entry)
	}

	return nil
}
