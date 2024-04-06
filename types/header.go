package types

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

type WavHeader struct {
	RIFF_CHUNK      RiffChunk
	FORMAT_SUBCHUNK FormatSubChunk
	DATA_SUBCHUNK   DataSubChunk
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
