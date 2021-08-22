package jsrt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

var ReadLineBufSize = 4096

type JSON map[string]interface{}

func Sort(fin io.Reader, key string, fout io.Writer) error {
	ndjson, err := readAll(fin, key)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(fout)
	defer writer.Flush()

	sort.Slice(ndjson, func(i, j int) bool {
		iv := jsonStringValue(ndjson[i], key)
		jv := jsonStringValue(ndjson[j], key)

		return iv < jv
	})

	for _, j := range ndjson {
		line, err := json.Marshal(j)

		if err != nil {
			return err
		}

		fmt.Fprintln(writer, string(line))
	}

	return nil
}

func readAll(fin io.Reader, key string) ([]JSON, error) {
	reader := bufio.NewReader(fin)
	ndjson := make([]JSON, 2844047)

	for i := 0; true; i++ {
		line, err := readLine(reader)

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}

		var j JSON

		if err := json.Unmarshal([]byte(line), &j); err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}

		if _, ok := j[key]; !ok {
			return nil, fmt.Errorf("line %d: key '%s' not found", i+1, key)

		}

		//ndjson = append(ndjson, j)
		ndjson[i] = j
	}

	return ndjson, nil
}

func readLine(reader *bufio.Reader) ([]byte, error) {
	buf := make([]byte, 0, ReadLineBufSize)
	var err error

	for {
		line, isPrefix, e := reader.ReadLine()
		err = e

		if len(line) > 0 {
			buf = append(buf, line...)
		}

		if !isPrefix || err != nil {
			break
		}
	}

	return buf, err
}

func jsonStringValue(j JSON, key string) string {
	v, ok := j[key]

	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}
