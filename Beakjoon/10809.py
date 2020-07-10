word = 'baekjoon'
alphabetstring = 'abcdefghijklmnopqrstuvwxyz'
alphabet = [[alpha, -1] for alpha in alphabetstring]

for alpha in alphabet:
    for character in word:
        if alpha[0] == character:
            print(character, word.index(character)) # index 만 
            alpha[1] = word.index(character)
            

for alpha in alphabet:
    print(alpha[1], end = ' ')
