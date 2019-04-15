def same(a,b):
    if a.isupper() and b.islower() and b.upper() == a:
        return True
    elif b.isupper() and a.islower() and a.upper() == b:
        return True
    else:
       # print(a,b)
        return False
def check(v):
    maxLen = len(v)
    nv = []
    i=0
    while i< maxLen:
        if i+1 == maxLen:
            nv.append(v[i])
            break
        if not same(v[i],v[i+1]):
            nv.append(v[i])
            i+=1
        else:
            i+=2
    return nv
def check2(v,a):
    maxLen = len(v)
    nv = []
    i=0
    while i< maxLen:
        if v[i] == a or v[i] == a.upper():
            i += 1
            continue
        if i+1 == maxLen:
            nv.append(v[i])
            break
        if not same(v[i],v[i+1]):
            nv.append(v[i])
            i+=1
        else:
            i+=2
    return nv
def getLen2(v,a):
    result = v
    lastLen = 0
    while True:
        #print(''.join(result))
        result = check2(result,a)
        if lastLen == len(result):
            break
        lastLen = len(result)
    return lastLen
def getLen(v):
    result = v
    lastLen = 0
    while True:
        #print(''.join(result))
        result = check(result)
        if lastLen == len(result):
            break
        lastLen = len(result)
    return lastLen

if __name__ == "__main__":
    with open('input.txt',encoding='utf-8') as file:
        v = file.read()
        l = 'abcdefghijklmnopqrstuvwxyz'
        minX = 9999
        for i in range(len(l)):
            c = l[i]
            length = getLen2(v,c)
            print(c,length)
            if length < minX:
                minX = length
        print(minX)