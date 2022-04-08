'''
relational db design
BCNF,3NF 
'''

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
    input:
    R:a relation schema
    F: F+
    K:super key
    
    return True is R is BCNF
    or return (set_k,set_v)[FD] which violate the rule
    and set_k intersection set_V = empty
    '''
    maxlen=0
    target_setk=set()
    target_setv=set()
    for tuple in F_plus:#AB->ABCD
        set_k,set_v=tuple
        diff_setk=set_k.difference(set_v)
        if( set_k.issuperset(set_v) )or( diff_setk in K )or (set_k in K) or(not diff_setk) :pass
        #  diff_setk=={}  false
        #  not diff_setk  true
        else:
            diff_setv=set_v.difference(set_k)
            if (len(diff_setk)+len(diff_setv))>maxlen:
                maxlen=(len(diff_setk)+len(diff_setv))
                target_setk=diff_setk
                target_setv=diff_setv
    if maxlen==0:
        return True
    else:
        return(target_setk,target_setv)

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
    '''
    input:
    R:ralational schema
    F:F_plus
    
    return:
    K:candidate key
    '''
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

def extraneous_attibutes(F_origin):
    F=F_origin.copy()
    beta=set()
    for tuple in F:
        fset_k,fset_v=tuple
        beta=beta.union(fset_v)

    for tuple in F:
        fset_k,fset_v=tuple
        tfset_k=fset_k
        tfset_v=fset_v
        for A in fset_k:
            F_tmp=F.copy()
            tfset_k=frozenset(tfset_k)
            AS=tfset_k.difference(A)#AS as gamma
            result=closure_AS(AS,F_tmp)
            if result.issuperset(beta):
                F_tmp.discard((tfset_k,fset_v))
                F_tmp.add((AS,fset_v))
                F=F_tmp
                tfset_k=set(tfset_k)
                tfset_k=tfset_k.difference(A)
                pass
        for A in fset_v:
            F_tmp=F.copy()
            tfset_v=frozenset(tfset_v)
            F_tmp.discard((fset_k,tfset_v))
            AS=tfset_v.difference(A)
            F_tmp.add((fset_k,AS))
            result=closure_AS(fset_k,F_tmp)
            if A in result:
                F=F_tmp
                tfset_v=set(tfset_v)
                tfset_v=tfset_v.difference(A)
                pass
    return F
                
def canonical_cover(F_origin):
    '''
    input:
    F_origin: just F
    
    return:
    F: canonical cover F
    
    '''
    F0=F_origin.copy()
    while(1):
        F=F0.copy()
        for tuple in F0:
            fset_k,fset_v=tuple
            for tmp_tuple in F0:
                tfset_k,tfset_v=tmp_tuple
                if fset_k==tfset_k:
                    if fset_v==tfset_v:pass
                    else:
                        tmp=fset_v.union(tfset_v)
                        F.discard((fset_k,fset_v))
                        F.discard((tfset_k,tfset_v))
                        F.add((fset_k,tmp))
                        
        F=extraneous_attibutes(F)
        
        if  F==F0:break
        else:F0=F              
    return F
    pass

def is_3NF(R,F,K):
    '''
    input:
    R:a relation schema
    F:just F
    K:candidate key 
    
    return True is R is 3NF
    or return (set_k,set_v)[FD] which violate the rule
    and set_k intersection set_V = empty
    '''
    maxlen=0
    target_setk=set()
    target_setv=set()
    K_super=ge_superkey(R,F)
    K_A=set()
    for k in K:
        K_A=K_A.union(k)
    for tuple in F:#AB->ABCD
        set_k,set_v=tuple
        diff_setk=set_k.difference(set_v)
        #  diff_setk=={}  false
        #  not diff_setk  true
        if( set_k.issuperset(set_v) )or( diff_setk in K_super )or (set_k in K_super) or(not diff_setk) :pass
        else:
            diff_setv=set_v.difference(set_k)
            for A in diff_setv:
                if A not in K_A:
                    if (len(diff_setk)+len(diff_setv))>maxlen:
                        maxlen=(len(diff_setk)+len(diff_setv))
                        target_setk=diff_setk
                        target_setv=diff_setv
    if maxlen==0:
        return True
    else:
        return(target_setk,target_setv)

def TNF_decomposition(R,F_c,K_c):
    '''
    input:
    F_c: canonical cover key
    K_c: candidate key
    
    return:
    result
    '''
    result=set()
    bool=0
    for tuple in F_c:
        set_k,set_v=tuple
        tmp=set_k.union(set_v)
        result.add(tmp)
        for k in K_c:
            if k.issubset(tmp):bool=1
    if not bool:
        minlen=9999
        for k in K_c:
            if len(k)<minlen:
                minlen=len(k)
                tmp=k
        result.add(tmp)
        
    result_copy=result.copy()
    for r in result_copy:
        for tmp in result_copy:
            if r.issubset(tmp) and r!=tmp :
                result.discard(r)
    return result
    pass

R,F=set(),set()
para_input(R,F)
F_plus=closure_FD(R,F)
F_c=canonical_cover(F)
K_c=ge_candidatekey(R,F_plus)
result=TNF_decomposition(R,F_c,K_c)
#F_plus=closure_FD(R,F)
#K=ge_candidatekey(R,F_plus)
#print("initial:",'\n',R,'\n',F,'\n',K,'\n')
#result=BCNF_decomposition(R,F_plus,K)
#tmp=is_BCNF(R,F_plus,K)
#print(tmp)
print(result)
#list_F=list(F)
#list_F.sort()
#for tuple in list_F:
#    set_k,set_v=tuple
#    print("".join(set_k),"->","".join(set_v))
