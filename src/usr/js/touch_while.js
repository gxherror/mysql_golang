window.onload=function(){
    /*
     * 模块化编程
     * 1.必须由外部的封闭函数被调用一次
     * 2.封闭函数至少返回一个内部函数形成闭包
     */
    
    var clock=null;
    function Module(){
        var speed=10;
        //获取id
        function $(id){
        return document.getElementById(id);
        }
        //兼容浏览器
        function addEvent(element,type,handle){
            if(element.addEventListener){
                element.addEventListener(type,handle,false);
            }else if(element.attachEvent){
                element.attachEvent("on"+type,handle);
            }else{
                element['on'+type]=handle;
            }
        }
        //初始化init
        function init(){
            for(var i=0;i<4;i++){
                Module.createrow();
            }
        }
        //判断用户是否点击到了黑块
        function judge(ev){
            if(ev.target.className.indexOf('black')==-1){
                pass;
            }else{
                ev.target.className='cell';
                ev.target.parentNode.pass=1;//定义属性pass，表明此行row的黑块已经被点击
                Module.score();
            }
        }
        //游戏结束
        function fail(){
            clearInterval(clock);
            confirm('你的最终得分为'+parseInt(Module.$('score').innerHTML));
        }
        //创建div，参数className是其类名
        function creatediv(className){
            var div=document.createElement('div');
            div.className=className;
            return div;
        }
        //创建一个<div class="row">并且有四个子节点<div class="cell">
        function createrow(){
            var con=Module.$('con'),
                row=Module.creatediv('row'),//创建div className=row
                arr=Module.createcell(),//定义div cell的类名，其中一个为cell black
                i;
            con.appendChild(row);//添加row为con的子节点
            for(i=0;i<4;i++){
                row.appendChild(creatediv(arr[i]));//添加row的子节点
            }
            if(con.firstChild==null){
                con.appendChild(row);
            }else{
                con.insertBefore(row,con.firstChild);
            }
        }
        //创建一个类名的数组，其中一个为cell black，其余为cell
        function createcell(){
            var temp=['cell','cell','cell','cell'],
                i=Math.floor(Math.random()*4);//随机生成黑块位置
                temp[i]='cell black';
                return temp;
        }
        //黑块向下移动
        function move(){
            //Window.getComputedStyle() 方法给出应用活动样式表后的元素的所有CSS属性的值，并解析这些值可能包含的任何基本计算。
            var con=Module.$('con'),
                top = parseInt(window.getComputedStyle(con, null)['top']);
            if(speed+top>0){
                top=0;
            }else{
                top+=speed;
            }
            con.style.top=top+'px';
            if(top==0){
                Module.createrow();
                con.style.top='-200px';
                Module.derow();
            }else if(top==(-200+speed)){
                var rows=con.childNodes;
                if((rows.length==10)&&(rows[rows.length-1].pass!==1)){
                    Module.fail();
                }
            }
        }
        //加速函数
        function speedup(){
            speed+=2;
        }
        //删除div#con的子节点中最后那个<div class='row'>
        function derow(){
            var con=Module.$("con");
            if(con.childNodes.length==11){
                con.removeChild(con.lastChild);
            }
        }
        //记分
        function score(){
            var newscore=parseInt($("score").innerHTML)+10;
            $('score').innerHTML=newscore;
            if(newscore%50==0){
                Module.speedup();
            }
        }
        return {
            $:$,
            addEvent:addEvent,
            init:init,
            judge:judge,
            fail:fail,
            creatediv:creatediv,
            createrow:createrow,
            createcell:createcell,
            move:move,
            speedup:speedup,
            derow:derow,
            score:score
        };
    }
    
    var Module=Module();
    //添加点击事件
    Module.addEvent(Module.$('main'),'click',function(ev){
        Module.judge(ev);
    });
    //定时器每30s调用一次move();
    clock=setInterval(Module.move,30);
    Module.init();
}