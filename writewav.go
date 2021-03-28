package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type waveHeader struct {
	ChunkId       uint32
	ChunkSize     uint32
	Format        uint32
	SubchunkId    uint32
	SubchunkSize  uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Data          uint32
	Size          uint32
}

func generateHeader(sampleRate int, channels int, bitsPerSample int, databytes int) waveHeader {
	h := waveHeader{
		ChunkId:       binary.LittleEndian.Uint32([]byte("RIFF")),        // 'RIFF'
		ChunkSize:     uint32(databytes + 36),                            // file size-8 (in bytes)
		Format:        binary.LittleEndian.Uint32([]byte("WAVE")),        // 'WAVE'
		SubchunkId:    binary.LittleEndian.Uint32([]byte("fmt ")),        // 'fmt '
		SubchunkSize:  16,                                                // chunk length (16 bytes)
		AudioFormat:   1,                                                 // 1 is PCM, rest not important
		NumChannels:   uint16(channels),                                  // channels
		SampleRate:    uint32(sampleRate),                                // sampling frequency
		ByteRate:      uint32(sampleRate * channels * bitsPerSample / 8), // nBlockAlign * rate
		BlockAlign:    uint16(channels * bitsPerSample / 8),              // channels * bitsPerSample / 8
		BitsPerSample: uint16(bitsPerSample),                             // size of each sample (8, 16, 32)
		Data:          binary.LittleEndian.Uint32([]byte("data")),        // 'data'
		Size:          uint32(databytes),                                 // sound data size in bytes
	}

	return h
}

func printHeader(h waveHeader) {
	fmt.Printf("Wave Header Contents:\n")
	fmt.Printf(" ChunkId: %d\n", h.ChunkId)
	fmt.Printf(" ChunkSize: %d\n", h.ChunkSize)
	fmt.Printf(" Format: %d\n", h.Format)
	fmt.Printf(" SubchunkId: %d\n", h.SubchunkId)
	fmt.Printf(" SubchunkSize: %d\n", h.SubchunkSize)
	fmt.Printf(" AudioFormat: %d\n", h.AudioFormat)
	fmt.Printf(" NumChannels: %d\n", h.NumChannels)
	fmt.Printf(" SampleRate: %d\n", h.SampleRate)
	fmt.Printf(" ByteRate: %d\n", h.ByteRate)
	fmt.Printf(" BlockAlign: %d\n", h.BlockAlign)
	fmt.Printf(" BitsPerSample: %d\n", h.BitsPerSample)
	fmt.Printf(" Data: %d\n", h.Data)
	fmt.Printf(" Size: %d\n", h.Size)
}

func main() {
	var sampleRate int = 44100
	var blockFrames int = 256

	args := os.Args[1:]
	fileName := args[0]
	duration, _ := strconv.ParseFloat(args[1], 16)  // Unsafe: ignoring errors
	frequency, _ := strconv.ParseFloat(args[2], 16) // Unsafe: ignoring errors

	fmt.Printf("Writing Wave File:\n")
	fmt.Printf(" Duration: %f seconds\n", duration)
	fmt.Printf(" Sine Frequency: %f hertz\n", frequency)

	audioBuffer := make([]int16, blockFrames)

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var totalSamples int = int(duration * float64(sampleRate))

	databytes := totalSamples * 2 // 2 bytes per int16 sample

	header := generateHeader(sampleRate, 1, 16, databytes)

	printHeader(header)

	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		log.Fatal(err)
	}

	var phaseIndex int = 0
	var twopi float64 = 8 * math.Atan(1.0)

	for i := 0; i < totalSamples; i += blockFrames {
		for j := 0; j < blockFrames; j++ {
			audioBuffer[j] = int16(16000 * math.Sin(float64(phaseIndex)*twopi*frequency/float64(sampleRate)))
			phaseIndex++
		}

		err = binary.Write(file, binary.LittleEndian, audioBuffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}
