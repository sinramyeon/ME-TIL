def pleaseConformOnepass(caps):
    caps = caps + [caps[0]]
    for i in range(1, len(caps)):
        if caps[i] != caps[0]: #맨앞의 것과 다르면
            print('Peple in positions ', i, end='')
        else:
            print(' through ', i-1, 'flip your caps!')
