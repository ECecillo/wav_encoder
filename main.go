package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ECecillo/wav_encoder/header"
	"github.com/ECecillo/wav_encoder/sound"
	"github.com/ECecillo/wav_encoder/types"
)

func main() {
	var sampleRate int32 = 44100
	var numChannels int16 = 2
	var bitsPerSample int16 = 16

	wavFileHeader := types.WavHeader{
		RIFF_CHUNK: types.RiffChunk{
			CHUNK_ID:   "RIFF",
			CHUNK_SIZE: 36,
			FORMAT:     "WAVE",
		},
		FORMAT_SUBCHUNK: types.FormatSubChunk{
			SUBCHUNK1_ID:    "fmt ",
			SUBCHUNK1_SIZE:  16,
			AUDIO_FORMAT:    1,
			NUM_CHANNELS:    numChannels,
			SAMPLE_RATE:     sampleRate,
			BYTE_RATE:       sampleRate * int32(numChannels) * (int32(bitsPerSample) / 8),
			BLOCK_ALIGN:     numChannels * (bitsPerSample / 8),
			BITS_PER_SAMPLE: bitsPerSample,
		},
		DATA_SUBCHUNK: types.DataSubChunk{
			SUBCHUNK2_ID:   "data",
			SUBCHUNK2_SIZE: 0,
		},
	}

	wav, err := os.Create("test.wav")
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(wav)
	header.WriteWavHeader(writer, &wavFileHeader)
	startPos, _ := wav.Seek(0, io.SeekCurrent)
	sound.WriteWavSound(writer, &wavFileHeader)
	endPos, _ := wav.Seek(0, io.SeekCurrent)

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataSize := endPos - startPos
	wavFileHeader.DATA_SUBCHUNK.SUBCHUNK2_SIZE = int32(dataSize)
	wavFileHeader.RIFF_CHUNK.CHUNK_SIZE = int32(endPos) - 8
	sound.WriteWAVDataSizeToBuffer(wav, int32(dataSize))

	fmt.Println("Ecriture terminé 📝 !")
	fmt.Println("WAV Header 💿", wavFileHeader)

	header.VerifyWavHeader("test.wav", &wavFileHeader)
}
