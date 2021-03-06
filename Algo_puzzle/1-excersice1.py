
def excercise1(caps):
    start = forward = backward = 0
    intervals = []
    for i in range(1, len(caps)):
        if caps[start] != caps[i]:
            intervals.append((start, i-1, caps[start]))
            if caps[start] == 'F':
                forward += 1
            else:
                backward += 1
            start = i
    # for문 끝나고 마지막 꾸간
    intervals.append((start, i-1, caps[start]))
    if caps[start] == 'F':
        foward +=1
    else:
        backward +=1
    # 양이 더 적은 것을 뒤집는다.
    if forward < backward:
        flip ='F'
    else:
        flip ='B'
    for t in intervals:
        if t[2] == flip:
            if t[0] != t[1]:
                print('People in positions ', t[0], 'through ', t[1], 'flip your caps')
            else:
                print('Person at positions ', t[0], 'flip your cap')
                
