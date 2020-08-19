#!/usr/bin/env python3

with open('input.txt') as f:
    _ = f.readline()
    r_songs = f.readline()
    f_songs = f.readline()

r_songs = r_songs.split()
f_songs = f_songs.split()


merge_list = list(zip(r_songs, f_songs))
merge_list = [item for tempList in merge_list for item in tempList]

merge_str = " ".join(merge_list)

print(merge_str)