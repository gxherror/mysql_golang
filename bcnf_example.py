U=input("请输入U\n")
U_list=U.split(",")
L_F=input("请输入函数依赖的左边\n")
L_F_list=L_F.split(",")
R_F=input("请输入函数依赖的右边\n")
R_F_list=R_F.split(",")
t=input("初始值\n")
res=list(t)
a=[1]*len(L_F_list)
count=len(L_F_list)
while count>0:
    for i in range(len(L_F_list)):
        _tmp = L_F_list[i]
        tmp = list(_tmp)
        res1 = list()
        inter = list(set(res).intersection(set(tmp)))   #两个集合的交集
        inter=list(set(inter).difference(set(tmp)))#在前者之中不在后者里面
        if len(inter)==0:  # 判断是否包含
            _tmp = R_F_list[i]
            tmp = list(_tmp)
            res1 = list(set(res).union(set(tmp)))
            if(a[i]==0):
                continue
            a[i]=0
            diff1=list(set(U).difference(set(res)))
            diff2=list(set(res1).difference(set(res)))
            if (len(diff1)==0|len(diff2)==0):
                break
            res = res1
    count=count-1;
print(res)
 
 
 
 
 
 
 
 
 
 
 