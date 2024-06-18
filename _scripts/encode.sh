#!/bin/bash

src=$1
dst=${1%%.*}
ffmpeg -y -i $src -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" tmp.mp4

if [ "$2" ]; then
    ffmpeg -y -i tmp.mp4 -stream_loop -1 -i $2  -map 0 -map 1:a -c:v copy -shortest $dst.mp4
    rm tmp.mp4
else
    mv tmp.mp4 $dst.mp4
fi

#rm $1
