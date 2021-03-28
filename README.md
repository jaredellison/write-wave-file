# write-wave-file

*Command line utility to generate wave files containing sine waves*

Adapted from "The Audio Programming Book" by Boulanger and Lazzarini, appendix C "Self-Describing Soundfile Formats" with translation from C to Go.

## Usage

```
#                    filename  seconds frequency
$ go run writewav.go hello.wav 10 220
Writing Wave File:
 Duration: 10.000000 seconds
 Sine Frequency: 220.000000 hertz
Wave Header Contents:
 ChunkId: 1179011410
 ChunkSize: 882036
 Format: 1163280727
 SubchunkId: 544501094
 SubchunkSize: 16
 AudioFormat: 1
 NumChannels: 1
 SampleRate: 44100
 ByteRate: 88200
 BlockAlign: 2
 BitsPerSample: 16
 Data: 1635017060
 Size: 882000
```