#!/usr/bin/env python3
import operator


def minion_game(S):
    vowels = 'AEIOU'
    player_win = ''
    score = 0
    player = {'Stuart':0, 'Kevin': 0}
    words = {}
    i = 0
    for s in range(len(S)):
        j = i + 1
        while j <= len(S):
            count = 0
            idx = 0
            # words[S[i:j]] = S.count(S[i:j])
            # print(S[i:j])
            while True:
                idx = S.find(S[i:j], idx)
                # print(idx)
                if idx >= 0:
                    count += 1
                    idx += 1
                else:
                    break
            words[S[i:j]] = count

            j += 1
        i += 1

    for i in words:
        if i[0] in vowels:
            player['Kevin'] += words[i]
        else:
            player['Stuart'] += words[i]
    # print(words)
    # print(player)
    if player['Stuart'] > player['Kevin']:
        print('Stuart', player['Stuart'])
    elif player['Stuart'] < player['Kevin']:
        print('Kevin', player['Kevin'])
    else:
        print('Draw')


# S = 'BANANANAAAS'
# S = 'BA'
with open('the-minion-game-test-14.txt') as fh:
    S = fh.read()
# print(S)
minion_game(S)
# print("{} {}".format(player, result))
