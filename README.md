# go-audio

AKA "I've spent 15 years of my life playing music but I don't know how it's implemented"

## Sine

* Make sure you have ffmpeg installed
* Rrun the raw binary file with `ffplay -f f32le -ar 44100 -showmode 1 out.bin` (whereas `44100` might change if you up the sample rate)
