# 괄호 풀기(괄호가영어로뭐임?)

def testParenthesis(text) :
    count = 0
    for n in text :
        if n=='(':
            count += 1
        else :
            count -= 1
    if count == 0:
        return True
    else:
        return False
        
