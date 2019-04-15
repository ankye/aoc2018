def sameLen(lines,i,j):
    linea = lines[i]
    lineb = lines[j]
    count = 0
    s = ''
    for step in range(0,len(linea)):
        if linea[step] == lineb[step]:
            count +=1
            s += linea[step]
    return count,s

if __name__ == '__main__':
    with open('input2.txt',encoding='utf-8') as file:
        lines = list(file)
        max = len(lines)
        mcount = 0
        s = '' 
        for i in range(0,max):
            for j in range(i+1,max):
                count,sub = sameLen(lines,i,j)
                if count > mcount:
                    mcount = count
                    s = sub
        print(mcount,s)

