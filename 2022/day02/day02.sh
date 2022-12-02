#!/usr/bin/env bash

declare -A shape_scores=([A]=1 [X]=1 [B]=2 [Y]=2 [C]=3 [Z]=3)

# for my sanity
declare -A shapes=([A]=rock [X]=rock [B]=paper [Y]=paper [C]=scissors [Z]=scissors)
declare -A results=([rock]="paper,B rock,A scissors,C" [paper]="scissors,C paper,B rock,A" [scissors]="rock,A scissors,C paper,B")

get_round_result() { # get_round_result opponent_shape (A,B,C) player_shape (X,Y,Z)
    opponent_shape=${shapes[$1]}
    player_shape=${shapes[$2]}
    result_index=0
    IFS=' ' read -r -a loss_draw_win <<< "${results[$player_shape]}"
    while [[ ${loss_draw_win[$result_index]%,*} != "$opponent_shape" ]]; do
        result_index=$((result_index + 1))
    done
    echo $((result_index * 3))
}

get_needed_shape() { # get_needed_shape opponent_shape (A,B,C) desired_result (X,Y,Z) = (loss,draw,win)
    opponent_shape=${shapes[$1]}
    results_len="${#results[@]}"
    desired_result_index=$((results_len - shape_scores[$2]))
    IFS=' ' read -r -a loss_draw_win <<< "${results[$opponent_shape]}"
    desired_result_shape=${loss_draw_win[$desired_result_index]#*,}
    echo "${shape_scores[$desired_result_shape]}"
}

total_score_part_a=0
total_score_part_b=0
while IFS=' ' read -r opponent player || [[ -n $opponent ]]; do
    # part a
    game_result_a=$(get_round_result "$opponent" "$player")
    shape_score_a=${shape_scores["$player"]}
    total_score_part_a=$((total_score_part_a + game_result_a + shape_score_a))

    # part b
    desired_result=$player
    shape_score_b=$(get_needed_shape "$opponent" "$desired_result")
    game_result_b=$(((shape_scores[$desired_result] - 1) * 3))
    total_score_part_b=$((total_score_part_b + game_result_b + shape_score_b))
done < "$1"

printf 'Total score Part A = %s Part B = %s\n' "$total_score_part_a" "$total_score_part_b"