import re
if __name__ == '__main__':
    with open('input.txt',encoding='utf-8') as file:
        lines = list(file)
        max = 1000
        g = [[0] * max for i in range(max)]
        maxX = max
        maxY = max
        for line in lines:
            result = re.match( r'(.*)@\s(\d+),(\d+):\s(\d+)x(\d+)', line)
            x = int(result.group(2))
            y = int(result.group(3))
            xp = int(result.group(4))
            yp = int(result.group(5))
            mx = x + xp
            my = y + yp
            if mx > maxX:
                maxX = mx
            if my >= maxY:
                maxY = my
            i = x
            j = y
            print(x,y,mx,my)
            for i in range(x,mx):
                for j in range(y,my):
                    g[i][j] += 1
                    
        count = 0
        for i in range(0,max):
            for j in range(0,max):
                if g[i][j] >=2 :
                    count +=1
        for i in range(0,max):
            print(g[i])
           
        print(maxX,maxY,count)