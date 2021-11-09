#输入，偏移，输出

#输入
password='tre pfl wzezjy cvmvc3?'
#偏移，共有25种情况，设置循环语句，循环25次，在每次循环中都输出结果
for i in range(26):
    for s in password:
        if s==' ': 
            print(' ',end='')
        else :
            if ord(s)+i>122:
               print(chr(ord(s)+i-26),end='')
            elif 97<=ord(s)+i<=122 :
                print(chr(ord(s)+i),end='')
    print('\n')

