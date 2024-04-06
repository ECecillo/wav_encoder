package sound

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/ECecillo/wav_encoder/types"
)

func WriteWavSound(writer *bufio.Writer, wavFileHeader *types.WavHeader) {
	// Variable to make a sound
	duration := 2
	maxAmplitude := 32760 // = 16 bits - 1 bits pour le signe, donc 15 bits.
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

func WriteWAVDataSizeToBuffer(writer *os.File, dataSize int32) {
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
