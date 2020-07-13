# ROT13은 알파벳 대문자와 소문자에만 적용할 수 있다. 알파벳이 아닌 글자는 원래 글자 그대로 남아 있어야 한다
# ROT13은 카이사르 암호의 일종으로 영어 알파벳을 13글자씩 밀어서 만든다.


alphabet = 'abcdefghijklmnopqrstuvwxyz'
ALPHABET = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
test = 'Baekjoon Online Judge'
test2 = 'One is 1'

def ROT13(text) :
    
    result = ''
    for s in text:
        if s == ' ':
            result += ''
            
        if s in alphabet:
            i = alphabet.index(s)
            if i > 12:
                result += alphabet[i-13]
            else:
                result += alphabet[i+13]
                
        if s in ALPHABET:
            i = ALPHABET.index(s)
            if i > 12:
                result += ALPHABET[i-13]
            else:
                result += ALPHABET[i+13]
        else :
            result += s
    print(result)
    
ROT13(test)
ROT13(test2)
