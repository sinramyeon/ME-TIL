def bestTimeToPartySmart(schedule):
    times = []
    for c in schedule:
        times.append((c[0], 'start'))
        times.append((c[1], 'end'))
    
    def sortList(tlist):
        # 리스트 안의 가장 작은 값 찾아서 처음으로 보냄
        for ind in range(len(tlist)-1):
            iSm = ind
            for i in range(ind, len(tlist)):
                if tlist[iSm][0] > tlist[i][0]:
                    iSm = i
            tlist[ind], tlist[iSm] = tlist[iSm], tlist[ind]
    
    def chooseTime(times):
        rcount = 0
        maxcount = time = 0
        for t in times:
            if t[1] == 'start':
                rcount = rcount +1
            elif t[1] == 'end':
                rcount = rcount -1
            if rcount > maxcount:
                maxcount = rcount
                time = t[0]
        return maxcount, time
        
    sortList(times)
    maxcount, time = chooseTime(times)
    print("Best time to attend the party is at ", time, ' o clock : ', maxcount, 'celebrities will be in there')
    
