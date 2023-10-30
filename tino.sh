#!/usr/bin/bash

# trans timestamp into Standard Beijing time
((time= $(date +%s) + 8 * 3600))

((second = $time % 60))
((minute = ($time / 60) % 60))
((hour = ($time / 3600) % 24))

lines=$(tput lines)
cols=$(tput cols)


for ((i=1; i<($lines / 2); i++)); do
    echo ""
done


((leftPad= ($cols -28) / 2))
for ((i=0; i<leftPad; i++)); do
    printf " "
done
# echo -e "\033[1;32m$hour - $minute - $second\033[0m @ $lines * $cols"
echo -e "\033[1;32m$(date)\033[0m"


((actualLen = $cols - ($cols % 6)))
((leftPad = ($cols % 6) / 2))
for ((i=0; i<leftPad; i++)); do
    printf " "
done


((pos = (( time / 60 ) % 1440 + (l440 / actualLen / 2)) * actualLen / 1440 ))
((offNine = 9 * actualLen / 24))
((offSix = 18 * actualLen / 24))

color="\033[1;33m"
if ((time % 2 == 0)); then
    color="\033[1;37m"
fi
printf "\033[1;34m"
for ((i=0; i<cols-leftPad*2; i++)); do
    if [ $pos -eq $i ]; then
        if [ $pos -lt $offNine ]; then
            printf "${color}>\033[1;34m"
        elif [ $pos -gt $offSix ]; then 
            printf "${color}>\033[1;34m"
        else   
            printf "${color}>\033[1;32m"
        fi
    elif [ $i -eq $offNine ]; then
        printf "\033[1;32m{"
    elif [ $i -eq $offSix ]; then
        printf "}\033[1;34m"
    elif [ $i -lt $pos ]; then
        printf "="
    else 
        printf "_"
    fi

done

# flush
# to set cursor to new line
echo ""