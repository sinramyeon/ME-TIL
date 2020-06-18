def printReverse(text) :
    text = text.split() ## how about use a join() and reverse() ?
    for t in text :        
        if len(t) == 1:
            print(t, end = ' ')
        else :
            length = len(t)
            while length > 0 : 
                print(t[length-1], end = '')
                length -= 1
            print(' ', end='')
