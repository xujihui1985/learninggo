package stream

import (
	"bytes"
	"io"
)

var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
}

func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	input := bytes.NewBuffer(data)

	size := len(find)

	buf := make([]byte, size)
	end := size - 1

	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			output.Write(buf[:end])
			return
		}

		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}
			continue
		}
		output.WriteByte(buf[0])
		copy(buf, buf[1:])
	}
}

func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	input := bytes.NewReader(data)

	size := len(find)

	idx := 0

	for {

		b, err := input.ReadByte()
		if err != nil {
			break
		}

		if b == find[idx] {

			idx++

			if idx == size {
				output.Write(repl)
				idx = 0
			}
			continue
		}

		if idx != 0 {
			output.Write(find[:idx])

			input.UnreadByte()

			idx = 0
			continue
		}

		output.WriteByte(b)
		idx = 0
	}
}

func algThree(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	input := bytes.NewBuffer(data)

	size := len(find)

	buf := make([]byte, size)
	end := size - 1

	if n, err := input.Read(buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {

		var err error
		buf[end:][0], err = input.ReadByte()
		if err != nil {
			output.Write(buf[:end])
			return
		}
		// if _, err := io.ReadFull(input, buf[end:]); err != nil {
		// }

		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			if n, err := input.Read(buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}
			continue
		}
		output.WriteByte(buf[0])
		copy(buf, buf[1:])
	}
}
