#!/usr/bin/env bash

find_common_letter() {
    num_strings=$#
    declare -A common_occurrences
    for str in "$@"; do
        declare -A string_occurrences=() # to prevent duplicates within the same string
        for ((i = 0; i < ${#str}; i++)); do
            letter=${str:$i:1}
            if [[ -z ${string_occurrences[$letter]} ]]; then
                if [[ -z ${common_occurrences[$letter]} ]]; then
                    common_occurrences[$letter]=1
                else
                    common_occurrences[$letter]=$((common_occurrences[$letter]+1))
                fi
                string_occurrences[$letter]=1;
            fi
            if [[ ${common_occurrences[$letter]} -eq $num_strings ]]; then
                echo "$letter"
                return
            fi
        done
    done
}

compartments() {
    str_len=${#1}
    first_compartment="${1:0:$str_len/2}"
    second_compartment="${1:$str_len/2}"
    find_common_letter "$first_compartment" "$second_compartment"
}

declare -A priorities
p=0
for char in {a..z}; do
    p=$((p + 1))
    priorities[$char]=$p
done
for char in {A..Z}; do
    p=$((p + 1))
    priorities[$char]=$p
done

priority_sum_a=0
priority_sum_b=0
group=()
while IFS= read -r rucksack || [[ -n "$rucksack" ]]; do
    echo "$rucksack"

    # part a
    # common_letters=$(compartments "$rucksack")
    # for i in $common_letters; do
    #     priority_sum=$((priority_sum + ${priorities[$i]}))
    # done
    common_letter=$(compartments "$rucksack")
    priority_sum_a=$((priority_sum_a + ${priorities[$common_letter]}))

    # part b
    group+=("$rucksack")
    if [[ ${#group[@]} = 3 ]]; then
        common_letter=$(find_common_letter "${group[@]}")
        priority_sum_b=$((priority_sum_b + ${priorities[$common_letter]}))
        group=()
    fi
done < "$1"

printf 'Sum A: %d Sum B: %d\n' $priority_sum_a $priority_sum_b