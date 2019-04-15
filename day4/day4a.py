import re
import time

def getArr(dict,gid):

    if gid in dict:
        arr = dict[gid]
    else:
        sDict = [-1]*62
        sDict[60] = 0
        sDict[61] = 0
        dict[gid] = sDict
    return dict[gid]
def updateStatus(dict,gid,start,end,status):
    # if gid == 421:
    #     print(gid,start,end,status)
    arr = getArr(dict,gid)
    if status == 1:
        #print(arr[60],start,end)
        arr[60] += end - start
        #print(arr[60])
    for i in range(start,end):
        if arr[i] == -1 and status == 1:
            arr[i] = 1
        elif status == 1:
            arr[i] += 1
        # else:
        #     arr[i] = 0
        

if __name__ == "__main__":
    with open('input.txt',encoding='utf-8') as file:
        lines = list(file)
        gmap = {}
        currentG = None
        lines3 = []
        for line in lines:
           
            result = re.match( r'\[(.*)\s(.*)\]\s(Guard #\d+)*(.*)', line)
            d = int(str(result.group(1)).replace('-',''))
            h = int(str(result.group(2)).replace(':',''))
            if h> 60:
                d +=1
                h=0
            status = 0
            s = str(result.group(4))
            if s == 'falls asleep':
                status = 0
            elif s == 'wakes up':
                status = 1
            g = result.group(3)
            big = 1
            if g is not None:
                g = int(str(g).replace('Guard #',''))
                big= 0
            else:
                g = 0
            lines3.append([d*10000+big * 1000+h,g,d,h,status,line])
        maxM = 60
        gid = 0
        start = 0
        lastStatus = -1
        lines3.sort()
        proc = False
        for v in lines3:
           # print(v)
            
            if v[1] > 0:
                if gid > 0 and proc:
                    if start < maxM and lastStatus >= 0:
                        if lastStatus == 0:
                            print('helloworld')
                            updateStatus(gmap,gid,start,maxM,1)
                        else:
                            updateStatus(gmap,gid,start,maxM,0)
                gid = v[1]
                getArr(gmap,gid)[61] += 1
                start = 0
                lastStatus = -1
                proc = False
            else:
                
                updateStatus(gmap,gid,start,v[3],v[4])
                lastStatus = v[4]
                start = v[3]
                proc = True
        if gid > 0:
            if start < maxM and lastStatus >= 0:
                if lastStatus == 0:
                    print('helloworld')
                    updateStatus(gmap,gid,start,maxM,1)
                else:
                    updateStatus(gmap,gid,start,maxM,0)

        # result = 0
        # max = 0
        # maxk = 0
        # for k,v in gmap.items():
        #     if max < v[60]:
        #         print(k,v[60],v[61])
        #         max = v[60]
        #         maxk = k

        # line = gmap[maxk]
        # max = 0
        # for i in range(len(line)-2):
        #     if max < line[i]:
        #         max = line[i]
        #         print(i)
       

        for k,v in gmap.items():
            for i in range(len(v)-2):
                if v[i] == v[61]:
                    print(k,i,v)
            

