import sys

def getName(num,upper):
    num = num -1
    l = 'abcdefghijklmnopqrstuvwxyz'
    one = int(num/len(l))
    two = num % len(l)
    
    s= l[one] + l[two]
    # if one == 0:
    #     s = l[two]
    if upper:
        return s.upper()
    else:
        return s
def pk(g,x,y,n,step):
    if g[x][y] == 0:
        g[x][y] = ['',0,9999]
    pn,pid,l = g[x][y]
    if pid != n:
        if l > step+1:
            g[x][y] = [getName(n,False),n,step+1]
            return True
        elif l == step+1:
            g[x][y] = ['-',0,l]
            return True
        else:
            return False
    else:
        if l > step+1:
            g[x][y] = [getName(n,False),n,step+1]
            return True
        return False
def walkN(g,step,x,y,pid):
    if x < 0 or x >= len(g):
        return False
    if y < 0 or y >= len(g):
        return False
    return pk(g,x,y,pid,step)
def walkNode(g,step,x,y,pid):
    
    if walkN(g,step,x,y-1,pid):
        walkNode(g,step+1,x,y-1,pid)
    if walkN(g,step,x,y+1,pid):
        walkNode(g,step+1,x,y+1,pid)
    if walkN(g,step,x-1,y,pid):
        walkNode(g,step+1,x-1,y,pid)
    if walkN(g,step,x+1,y,pid):
        walkNode(g,step+1,x+1,y,pid)


    # for i in range(-1,2):
    #     for j in range(-1,2):
    #         if i == 0 and j == 0:
    #             continue
    #         else:
    #             if walkN(g,step,x-i,y-j,pid):
    #                 walkNode(g,step+1,x-i,y-j,pid)

if __name__ == "__main__":
    sys.setrecursionlimit(10000000) #例如这里设置为一百万

    with open('input.txt',encoding='utf-8') as file:
        lines = list(file)
        max = 400
        g = [[0]*max for i in range(max)]
        num = 1
        nodes = []
        for line in lines:
            arr = line.replace('\n','').split(',')
            y = int(arr[0])
            x = int(arr[1])
            
            name = (getName(num,True))
            
            g[x][y] = [name,num,0]
            nodes.append([x,y,num])
            num += 1
        close = {}
        counter = {}
        for node in nodes:
            x,y,pid = node
            step = 0
            close[pid] = [x,y,pid]
            counter[pid] = 0
            walkNode(g,step,x,y,pid)

        for i in range(max):
            for j in range(max):
                pid = g[i][j][1]
                if pid > 0:
                    counter[pid] += 1

        for i in range(max):
            c = g[0][i][1]
            if c in close:
                close[c] = None
            c = g[max-1][i][1]
            if c in close:
                close[c] = None
            c = g[i][0][1]
            if c in close:
                close[c] = None
            c = g[i][max-1][1]
            if c in close:
                close[c] = None   
 
        maxNum = 0
        maxK = None
        for k,v in close.items():
            if v is not None:
               # print(getName(v[2],True))
                if maxNum < counter[k]:
                    maxNum = counter[k]
                    maxK = getName(k,True)
        # for v in g:
        #     print(v)
        for v in g:
            s = []
            for v1 in v:
                s.append(v1[0])
        print(maxK,maxNum)
