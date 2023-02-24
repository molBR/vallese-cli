package utils

import (
	"encoding/binary"
	"os"
	"strings"

	"github.com/gordonklaus/portaudio"
)

func Record(watchdog *bool) {
	fileName := "audio"
	if !strings.HasSuffix(fileName, ".aiff") {
		fileName += ".aiff"
	}
	f, err := os.Create(fileName)
	Chk(err)

	// form chunk
	_, err = f.WriteString("FORM")
	Chk(err)
	Chk(binary.Write(f, binary.BigEndian, int32(0))) //total bytes
	_, err = f.WriteString("AIFF")
	Chk(err)

	// common chunk
	_, err = f.WriteString("COMM")
	Chk(err)
	Chk(binary.Write(f, binary.BigEndian, int32(18)))                  //size
	Chk(binary.Write(f, binary.BigEndian, int16(1)))                   //channels
	Chk(binary.Write(f, binary.BigEndian, int32(0)))                   //number of samples
	Chk(binary.Write(f, binary.BigEndian, int16(32)))                  //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	Chk(err)

	// sound chunk
	_, err = f.WriteString("SSND")
	Chk(err)
	Chk(binary.Write(f, binary.BigEndian, int32(0))) //size
	Chk(binary.Write(f, binary.BigEndian, int32(0))) //offset
	Chk(binary.Write(f, binary.BigEndian, int32(0))) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		Chk(err)
		Chk(binary.Write(f, binary.BigEndian, int32(totalBytes)))
		_, err = f.Seek(22, 0)
		Chk(err)
		Chk(binary.Write(f, binary.BigEndian, int32(nSamples)))
		_, err = f.Seek(42, 0)
		Chk(err)
		Chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)))
		Chk(f.Close())
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	Chk(err)
	defer stream.Close()

	Chk(stream.Start())
	for *watchdog {
		Chk(stream.Read())
		Chk(binary.Write(f, binary.BigEndian, in))
		nSamples += len(in)
	}
	Chk(stream.Stop())
}
