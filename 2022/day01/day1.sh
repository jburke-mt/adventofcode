#!/usr/bin/env bash

max_sum=0
cur_sum=0
cal_sums=()

append_cur_sum() {
    if [[ $cur_sum -gt $max_sum ]]; then
        max_sum=$cur_sum
    fi
    cal_sums+=("$cur_sum")
    cur_sum=0
}

while IFS= read -r elf_line || [[ -n "$elf_line" ]]; do
    if [[ -z $elf_line ]]; then
        append_cur_sum
    else
        ((cur_sum=cur_sum + elf_line))
    fi
done < "$1"
append_cur_sum # handle no newline at end of file
echo $max_sum # part a
IFS=$'\n' read -r -d '' -a sorted_cals < <(printf "%s\n" "${cal_sums[@]}" | sort -n -r && printf '\0')
top_3_sums=$((sorted_cals[0] + sorted_cals[1] + sorted_cals[2]))
echo $top_3_sums # part b