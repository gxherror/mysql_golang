

def closure_FD(R,F):
    F=F.copy()
    while(1):
        F_copy0=F.copy()
        for tuple in F_copy0:
            fset_k,fset_v=tuple
            diffs=R.difference(fset_k.union(fset_v))
            if diffs==None: bool=1
            for diff in diffs:
                set_k=set(fset_k)
                set_v=set(fset_v)
                set_k.add(diff)
                set_v.add(diff)
                F.add((frozenset(set_k),frozenset(set_v))) 
        F_copy1=F.copy()       
        for tuple in F_copy1:
            fset_k,fset_v=tuple
            for tuple in F_copy1:
                tset_k,tset_v=tuple
                if fset_v==tset_k:
                    F.add((fset_k,tset_v))       
            pass
        if F_copy0==F:break
    return F

def is_BCNF(R,F_plus,K):
    '''
    return True is R is BCNF
    or return (set_k,set_v)[FD] which violate the rule
    and set_k intersection set_V = empty
    '''
    for tuple in F_plus:#AB->ABCD
        set_k,set_v=tuple
        diff_setk=set_k.difference(set_v)
        if set_k.issuperset(set_v) or diff_setk in K :pass
        else:
            diff_setv=set_v.difference(set_k)
            return (diff_setk,diff_setv)
    return True

def para_input(R,F):      
    R_str=input("R:\nA,B\n:")
    F_str=input("F:\nA->B,C->D,etc\n:")
    #K_str=input("K:\nAB,CD,etc\n:")
    R.update(set(R_str.split(',')))
    F_list=F_str.split(',')
    #K_list=K_str.split(',')
    for f in F_list:
        tmp=f.split('->')
        F.add((frozenset(set(tmp[0])),frozenset(set(tmp[1]))))
    #for k in K_list:
    #    K.add(frozenset(k))
    return

def BCNF_decomposition(R,F_plus,K):
    result=set()
    result.add(frozenset(R))
    while(1):
        result_copy=result.copy()
        for R_sub in result_copy:
            F_sub=set()
            for tuple in F_plus:
                set_k,set_v=tuple
                if set_k.issubset(R_sub) and set_v.issubset(R_sub):
                    F_sub.add(tuple)
            K_sub=ge_superkey(R_sub,F_sub)
            tmp=is_BCNF(R_sub,F_sub,K_sub)
            if tmp==True :
                pass
            else:
                diff_setk,diff_setv=tmp
                result.discard(R_sub)
                result.add(R_sub.difference(diff_setv))
                result.add(diff_setk.union(diff_setv))
        if result_copy==result:break
    return result

def ge_superkey(R,F):
    K=set()
    for tuple in F:
        set_k,set_v=tuple
        result=closure_AS(set_k,F)
        if result==R:
            K.add(frozenset(set_k))
            #! update a set as union  
            #! add a set as a element
    return K   

def ge_candidatekey(R,F):
    K=ge_superkey(R,F)
    #TODO:to binary search
    K_copy=K.copy()
    for k in K_copy:
        for tmp in K_copy:
            if k.issuperset(tmp) and k!=tmp :
                K.discard(k)
    return K
            
def closure_AS(AS,F):
    
    result=AS
    while(1):
        result_old=result
        for tuple in F:
            set_k,set_v=tuple
            if set_k.issubset(result):#!
                result = result.union(set_v)
        if result_old==result:
            break
    return result

def canonical_cover():
    pass
       
R,F=set(),set()
para_input(R,F)
F_plus=closure_FD(R,F)
K=ge_candidatekey(R,F_plus)
print("initial:",'\n',R,'\n',F,'\n',K,'\n')
result=BCNF_decomposition(R,F_plus,K)
#tmp=is_BCNF(R,F_plus,K)
#print(tmp)
print(result)
#list_F=list(F)
#list_F.sort()
#for tuple in list_F:
#    set_k,set_v=tuple
#    print("".join(set_k),"->","".join(set_v))
