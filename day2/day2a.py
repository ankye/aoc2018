if __name__ == '__main__':
    print('start')
    with open('input.txt',encoding='utf-8') as file:
        lines = list(file)
        two = 0
        three = 0
        for line in lines:
            print(line)
            dict = {}
            for c in line:
                if c in dict:
                    dict[c] += 1
                else:
                    dict[c] = 1
            count2=0
            count3=0
           
            for k,v in dict.items():
                if v == 2:
                    count2 =1
                elif v == 3:
                    count3 = 1
            if count2 == 1:
                two += 1
            if count3 == 1:
                three +=1
        print("two ",two)
        print("three ",three)
        print("two * three ",two * three)