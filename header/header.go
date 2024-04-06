package header

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/ECecillo/wav_encoder/types"
	"github.com/ECecillo/wav_encoder/utils"
)

func WriteWavHeader(writer *bufio.Writer, wavFileHeader *types.WavHeader) {
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

func VerifyWavHeader(fileName string, wavHeader *types.WavHeader) {
	file, err := os.Open(fileName)
	utils.Check(err)
	defer file.Close()

	var wavFileHeader [44]byte
	_, err = file.Read(wavFileHeader[:])
	utils.Check(err)

	bufferData := bytes.NewBuffer(wavFileHeader[:])

	var riffID [4]byte
	var chunkSize int32
	var format [4]byte
	var subchunk1ID [4]byte
	var subchunk1Size int32
	var audioFormat int16
	var numChannels int16
	var sampleRate int32
	var byteRate int32
	var blockAlign int16
	var bitsPerSample int16
	var subchunk2ID [4]byte
	var subchunk2Size int32

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
	if chunkSize == int32(wavHeader.RIFF_CHUNK.CHUNK_SIZE) {
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
	if subchunk1Size == int32(wavHeader.FORMAT_SUBCHUNK.SUBCHUNK1_SIZE) {
		fmt.Println("SUBCHUNK1_SIZE OK ✅")
	} else {
		panic("SUBCHUNK1_SIZE ❌")
	}
	if audioFormat == int16(wavHeader.FORMAT_SUBCHUNK.AUDIO_FORMAT) {
		fmt.Println("AUDIO_FORMAT OK ✅")
	} else {
		panic("AUDIO_FORMAT ❌")
	}
	if numChannels == int16(wavHeader.FORMAT_SUBCHUNK.NUM_CHANNELS) {
		fmt.Println("NUM_CHANNELS OK ✅")
	} else {
		panic("NUM_CHANNELS ❌")
	}
	if sampleRate == int32(wavHeader.FORMAT_SUBCHUNK.SAMPLE_RATE) {
		fmt.Println("SAMPLE_RATE OK ✅")
	} else {
		panic("SAMPLE_RATE ❌")
	}
	if byteRate == int32(wavHeader.FORMAT_SUBCHUNK.BYTE_RATE) {
		fmt.Println("BYTE_RATE OK ✅")
	} else {
		panic("BYTE_RATE ❌")
	}
	if blockAlign == int16(wavHeader.FORMAT_SUBCHUNK.BLOCK_ALIGN) {
		fmt.Println("BLOCK_ALIGN OK ✅")
	} else {
		panic("BLOCK_ALIGN ❌")
	}
	if bitsPerSample == int16(wavHeader.FORMAT_SUBCHUNK.BITS_PER_SAMPLE) {
		fmt.Println("BITS_PER_SAMPLE OK ✅")
	} else {
		panic("BITS_PER_SAMPLE ❌")
	}
	if string(subchunk2ID[:]) == wavHeader.DATA_SUBCHUNK.SUBCHUNK2_ID {
		fmt.Println("SUBCHUNK2_ID OK ✅")
	} else {
		panic("SUBCHUNK2_ID ❌")
	}
	if subchunk2Size == int32(wavHeader.DATA_SUBCHUNK.SUBCHUNK2_SIZE) {
		fmt.Println(subchunk2Size)
		fmt.Println("SUBCHUNK2_SIZE OK ✅")
	} else {
		panic("SUBCHUNK2_SIZE ❌")
	}
}
