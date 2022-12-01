#!/usr/bin/env bash

elf_index=0
max_sum=0
cur_sum=0
while IFS= read -r elf_line; do
    if [[ -z $elf_line ]]; then
        if [[ $cur_sum -gt $max_sum ]]; then
            max_sum=$cur_sum
        fi
        cur_sum=0
        ((elf_index=elf_index + 1)) # keeping track of this incase we need to do more with this input in future exercises
    else
        ((cur_sum=cur_sum + elf_line))
    fi
done < "$1"
echo $max_sum