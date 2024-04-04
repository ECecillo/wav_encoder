package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type RiffChunk struct {
	CHUNK_ID   string
	CHUNK_SIZE int32
	FORMAT     string
}

type FormatSubChunk struct {
	SUBCHUNK1_ID    string
	SUBCHUNK1_SIZE  int32
	AUDIO_FORMAT    int16
	NUM_CHANNELS    int16
	SAMPLE_RATE     int32
	BYTE_RATE       int32
	BLOCK_ALIGN     int16
	BITS_PER_SAMPLE int16
}

type DataSubChunk struct {
	SUBCHUNK2_ID   string
	SUBCHUNK2_SIZE int32
}

func (rif *RiffChunk) fill_defaults() {
	if rif.CHUNK_ID == "" {
		rif.CHUNK_ID = "RIFF"
	}
	if rif.FORMAT == "" {
		rif.FORMAT = "WAVE"
	}
}

func (fmt *FormatSubChunk) fill_defaults() {
	if fmt.SUBCHUNK1_ID == "" {
		fmt.SUBCHUNK1_ID = "fmt "
	}
	fmt.SUBCHUNK1_SIZE = 16
	fmt.AUDIO_FORMAT = 1
	if fmt.NUM_CHANNELS == 0 {
		fmt.NUM_CHANNELS = 2
	}
	fmt.BITS_PER_SAMPLE = 16
	fmt.BYTE_RATE = fmt.SAMPLE_RATE * int32(fmt.NUM_CHANNELS) * int32(fmt.BITS_PER_SAMPLE) / 8
	fmt.BLOCK_ALIGN = fmt.NUM_CHANNELS * int16(fmt.BITS_PER_SAMPLE) / 8
}

func VerifyWavHeader(fileName string, wavHeader *WavHeader) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	var wavFileHeader [44]byte
	_, err = file.Read(wavFileHeader[:])
	check(err)

	bufferData := bytes.NewBuffer(wavFileHeader[:])

	var riffID [4]byte
	var chunkSize uint32
	var format [4]byte
	var subchunk1ID [4]byte
	var subchunk1Size uint32
	var audioFormat uint16
	var numChannels uint16
	var sampleRate uint32
	var byteRate uint32
	var blockAlign uint16
	var bitsPerSample uint16
	var subchunk2ID [4]byte
	var subchunk2Size uint32

	binary.Read(bufferData, binary.BigEndian, &riffID)
	binary.Read(bufferData, binary.LittleEndian, &chunkSize)
	binary.Read(bufferData, binary.BigEndian, &format)
	binary.Read(bufferData, binary.BigEndian, &subchunk1ID)
	binary.Read(bufferData, binary.LittleEndian, &subchunk1Size)
	binary.Read(bufferData, binary.LittleEndian, &audioFormat)
	binary.Read(bufferData, binary.LittleEndian, &numChannels)
	binary.Read(bufferData, binary.LittleEndian, &sampleRate)
	binary.Read(bufferData, binary.LittleEndian, &byteRate)
	binary.Read(bufferData, binary.LittleEndian, &blockAlign)
	binary.Read(bufferData, binary.LittleEndian, &bitsPerSample)
	binary.Read(bufferData, binary.BigEndian, &subchunk2ID)
	binary.Read(bufferData, binary.LittleEndian, &subchunk2Size)

	if string(riffID[:]) == wavHeader.RIFF_CHUNK.CHUNK_ID {
		fmt.Println("CHUNK_ID OK ✅")
	} else {
		panic("CHUNK_ID ❌")
	}
	if chunkSize == uint32(wavHeader.RIFF_CHUNK.CHUNK_SIZE) {
		fmt.Println(chunkSize)
		fmt.Println("CHUNK_SIZE OK ✅")
	} else {
		panic("CHUNK_SIZE ❌")
	}
	if string(format[:]) == wavHeader.RIFF_CHUNK.FORMAT {
		fmt.Println("FORMAT OK ✅")
	} else {
		fmt.Println(string(format[:]), wavHeader.RIFF_CHUNK.FORMAT)
		panic("FORMAT ❌")
	}
	if string(subchunk1ID[:]) == wavHeader.FORMAT_SUBCHUNK.SUBCHUNK1_ID {
		fmt.Println("SUBCHUNK1_ID OK ✅")
	} else {
		panic("SUBCHUNK1_ID ❌")
	}
	if subchunk1Size == uint32(wavHeader.FORMAT_SUBCHUNK.SUBCHUNK1_SIZE) {
		fmt.Println("SUBCHUNK1_SIZE OK ✅")
	} else {
		panic("SUBCHUNK1_SIZE ❌")
	}
	if audioFormat == uint16(wavHeader.FORMAT_SUBCHUNK.AUDIO_FORMAT) {
		fmt.Println("AUDIO_FORMAT OK ✅")
	} else {
		panic("AUDIO_FORMAT ❌")
	}
	if numChannels == uint16(wavHeader.FORMAT_SUBCHUNK.NUM_CHANNELS) {
		fmt.Println("NUM_CHANNELS OK ✅")
	} else {
		panic("NUM_CHANNELS ❌")
	}
	if sampleRate == uint32(wavHeader.FORMAT_SUBCHUNK.SAMPLE_RATE) {
		fmt.Println("SAMPLE_RATE OK ✅")
	} else {
		panic("SAMPLE_RATE ❌")
	}
	if byteRate == uint32(wavHeader.FORMAT_SUBCHUNK.BYTE_RATE) {
		fmt.Println("BYTE_RATE OK ✅")
	} else {
		panic("BYTE_RATE ❌")
	}
	if blockAlign == uint16(wavHeader.FORMAT_SUBCHUNK.BLOCK_ALIGN) {
		fmt.Println("BLOCK_ALIGN OK ✅")
	} else {
		panic("BLOCK_ALIGN ❌")
	}
	if bitsPerSample == uint16(wavHeader.FORMAT_SUBCHUNK.BITS_PER_SAMPLE) {
		fmt.Println("BITS_PER_SAMPLE OK ✅")
	} else {
		panic("BITS_PER_SAMPLE ❌")
	}
	if string(subchunk2ID[:]) == wavHeader.DATA_SUBCHUNK.SUBCHUNK2_ID {
		fmt.Println("SUBCHUNK2_ID OK ✅")
	} else {
		panic("SUBCHUNK2_ID ❌")
	}
	if subchunk2Size == uint32(wavHeader.DATA_SUBCHUNK.SUBCHUNK2_SIZE) {
		fmt.Println(subchunk2Size)
		fmt.Println("SUBCHUNK2_SIZE OK ✅")
	} else {
		panic("SUBCHUNK2_SIZE ❌")
	}
}

type WavHeader struct {
	RIFF_CHUNK      RiffChunk
	FORMAT_SUBCHUNK FormatSubChunk
	DATA_SUBCHUNK   DataSubChunk
}

func WriteWavHeader(writer *bufio.Writer, wavFileHeader *WavHeader) {
	writer.Write([]byte(wavFileHeader.RIFF_CHUNK.CHUNK_ID))
	binary.Write(writer, binary.LittleEndian, wavFileHeader.RIFF_CHUNK.CHUNK_SIZE)
	writer.Write([]byte(wavFileHeader.RIFF_CHUNK.FORMAT))

	writer.Write([]byte(wavFileHeader.FORMAT_SUBCHUNK.SUBCHUNK1_ID))
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.SUBCHUNK1_SIZE)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.AUDIO_FORMAT)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.NUM_CHANNELS)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.SAMPLE_RATE)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.BYTE_RATE)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.BLOCK_ALIGN)
	binary.Write(writer, binary.LittleEndian, wavFileHeader.FORMAT_SUBCHUNK.BITS_PER_SAMPLE)

	writer.Write([]byte(wavFileHeader.DATA_SUBCHUNK.SUBCHUNK2_ID))
	binary.Write(writer, binary.LittleEndian, wavFileHeader.DATA_SUBCHUNK.SUBCHUNK2_SIZE)
}

func WriteWavSound(writer *bufio.Writer, wavFileHeader *WavHeader) {
	// Variable to make a sound
	var duration int = 2
	var maxAmplitude int = 32760 // = 16 bits - 1 bits pour le signe, donc 15 bits.
	var frequency float32 = 440

	nsamps := int(wavFileHeader.FORMAT_SUBCHUNK.SAMPLE_RATE) * duration
	angle := 2 * math.Pi * float64(frequency)

	for i := 0; i < nsamps; i++ {
		amplitude := float64(i) / float64(int(wavFileHeader.FORMAT_SUBCHUNK.SAMPLE_RATE)*maxAmplitude)
		value := math.Sin(angle * float64(i) / float64(wavFileHeader.FORMAT_SUBCHUNK.SAMPLE_RATE))

		channel1 := float32(amplitude*value) / 2                     // Goes high
		channel2 := float32(float64(maxAmplitude) - amplitude*value) // Low pitch

		binary.Write(writer, binary.LittleEndian, channel1)
		binary.Write(writer, binary.LittleEndian, channel2)
	}
}

func writeWAVDataSizeToBuffer(writer *os.File, dataSize int32) {
	_, err := writer.Seek(40, 0) // Position de SUBCHUNK2_SIZE dans le fichier
	if err != nil {
		fmt.Println(err)
		return
	}
	binary.Write(writer, binary.LittleEndian, dataSize)

	_, err = writer.Seek(4, 0) // Position de CHUNK_SIZE dans le fichier
	if err != nil {
		fmt.Println(err)
		return
	}
	binary.Write(writer, binary.LittleEndian, dataSize-8)
}

func main() {
	var sampleRate int32 = 44100
	var numChannels int16 = 2
	var bitsPerSample int16 = 16

	wavFileHeader := WavHeader{
		RIFF_CHUNK: RiffChunk{
			CHUNK_ID:   "RIFF",
			CHUNK_SIZE: 36,
			FORMAT:     "WAVE",
		},
		FORMAT_SUBCHUNK: FormatSubChunk{
			SUBCHUNK1_ID:    "fmt ",
			SUBCHUNK1_SIZE:  16,
			AUDIO_FORMAT:    1,
			NUM_CHANNELS:    numChannels,
			SAMPLE_RATE:     sampleRate,
			BYTE_RATE:       sampleRate * int32(numChannels) * (int32(bitsPerSample) / 8),
			BLOCK_ALIGN:     numChannels * (bitsPerSample / 8),
			BITS_PER_SAMPLE: bitsPerSample,
		},
		DATA_SUBCHUNK: DataSubChunk{
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
	WriteWavHeader(writer, &wavFileHeader)
	startPos, _ := wav.Seek(0, io.SeekCurrent)
	WriteWavSound(writer, &wavFileHeader)
	endPos, _ := wav.Seek(0, io.SeekCurrent)

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	dataSize := endPos - startPos
	wavFileHeader.DATA_SUBCHUNK.SUBCHUNK2_SIZE = int32(dataSize)
	wavFileHeader.RIFF_CHUNK.CHUNK_SIZE = int32(endPos) - 8
	writeWAVDataSizeToBuffer(wav, dataSize)

	fmt.Println("Ecriture terminé 📝 !")
	fmt.Println("WAV Header 💿", wavFileHeader)

	VerifyWavHeader("test.wav", &wavFileHeader)
}
