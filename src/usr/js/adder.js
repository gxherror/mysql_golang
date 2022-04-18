
function loadXMLDoc()
{
var xmlhttp;
document.getElementById("myDiv").innerHTML="修改后";
if (window.XMLHttpRequest)
    {// IE7+, Firefox, Chrome, Opera, Safari 代码
    xmlhttp=new XMLHttpRequest();
    }
else
    {// IE6, IE5 代码
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }

xmlhttp.onreadystatechange=function()
{
    if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
    document.getElementById("result").innerHTML=xmlhttp.responseText;
    }
}

//call when readySate change
xmlhttp.open("GET","/",true);
xmlhttp.send();
}


function addNumber() 
{
    var xmlhttp=new XMLHttpRequest();
    var url = "/adder/operate?Num1=" + document.getElementById("num1").value + "&Num2=" + document.getElementById("num2").value;
    document.getElementById("result").innerHTML=url;
    xmlhttp.open("GET", url=url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function ()
    {
        if (xmlhttp.readyState==4 && xmlhttp.status==200)
        {
          document.getElementById("result").innerHTML=xmlhttp.responseText;
        }
    }
    
}

function myFunction() {
    var x = document.getElementById("fname");
    x.value = x.value.toUpperCase();
    var xmlhttp;
    if (window.XMLHttpRequest)
    {// IE7+, Firefox, Chrome, Opera, Safari 代码
    xmlhttp=new XMLHttpRequest();
    }
    else
    {// IE6, IE5 代码
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlHttp.open("GET", "/", true);
    xmlhttp.send();
  }
