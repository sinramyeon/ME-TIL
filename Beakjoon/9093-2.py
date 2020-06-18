def printReverse2(text) :
   result = []
   text = text.split()
   for t in text :
     result.append(t[::-1])
   return ' '.join(result)
