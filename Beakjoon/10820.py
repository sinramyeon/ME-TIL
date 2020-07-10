text = 'This is String'
Alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
alphabet = 'abcdefghijklmnopqrstuvwxyz'
number = '01234566789'


def findText(text) :
    
    alpNum = 0
    capNum = 0
    intNum = 0
    spaceNum = 0
    for x in text:
        if x in alphabet:
            alpNum +=1
        if x in Alphabet:
            capNum +=1
        if x in number:
            intNum +=1
        if x == ' ':
            spaceNum +=1

    print(alpNum, capNum, intNum, spaceNum)
