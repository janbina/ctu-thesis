#!/bin/bash

# Get up to four windows from the current desktop
# and organize them on the screen

D=$(xdotool get_desktop)
I=0
ORIGIN=("nw" "ne" "sw" "se")

xdotool search --desktop $D --class "" | while read ID; do
    if [[ $I -gt 3 ]]; then exit 0; fi
    swmctl moveresize -id $ID \
                      -o ${ORIGIN[$I]} \
                      -xr .05 -yr .05 \
                      -wr .425 -hr .425
    ((I++))
done