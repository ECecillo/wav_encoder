# Context

Hard coding a wav encoder in Go for case studies and learning Go language.

This project is based on the following [video](https://www.youtube.com/watch?v=rHqkeLxAsTc&t=439s&ab_channel=Ferrabacus)

I strongly recommend watching the video to understand the concepts and the code.

Most of the stuff inside this README is based on the video and the [website](http://soundfile.sapp.org/doc/WaveFormat/)

## Definitions and Reminders

- 1 bytes = 8 bits

Little Endian : Least significant byte is stored first.
Big Endian : Most significant byte is stored first.

## WAV Header description

### RIFF Descriptor

- CHUNK_ID : 4 bytes, **"RIFF" | "ASCII"**, big-endian.
- CHUNK_SIZE : 4 bytes, File size - 8 (to ignore CHUNK_ID and CHUNK_SIZE), little-endian.
- FORMAT : 4 bytes, "WAVE", big-endian.

### Format Sub-chunk

- SUBCHUNK1_ID : 4 bytes, "fmt ", big-endian.
- SUBCHUNK1_SIZE : 4 bytes, 16, little-endian.
- AUDIO_FORMAT : 2 bytes, PCM = 1, little-endian.
- NUM_CHANNELS : 2 bytes, Mono = 1, Stereo = 2, ..., little-endian.
- SAMPLE_RATE : 4 bytes, 44100, little-endian.
- BYTE*RATE : 4 bytes, SAMPLE_RATE * NUM*CHANNELS * BITS_PER_SAMPLE / 8, little-endian.
- BLOCK_ALIGN : 2 bytes, NUM_CHANNELS \* BITS_PER_SAMPLE / 8, little-endian.
- BITS_PER_SAMPLE : 2 bytes, 16, little-endian.

### Subchunk 2 - Data sub chunk

- SUBCHUNK2_ID : 4 bytes, data, big-endian.
- SUBCHUNK2_SIZE : 4 bytes, NUM_SAMPLES * NUM*CHANNELS \* BITS_PER_SAMPLE / 8, little-endian.
  32
- DATA : n bytes, NUM_SAMPLES * NUM*CHANNELS \* BITS_PER_SAMPLE / 8 bytes, little-endian.

## Links and References

- [Wave Encoding Format](http://soundfile.sapp.org/doc/WaveFormat/)
- [Audio From Scratch](https://www.youtube.com/watch?v=rHqkeLxAsTc&t=439s&ab_channel=Ferrabacus)
