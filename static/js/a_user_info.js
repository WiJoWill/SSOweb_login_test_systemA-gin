/*
如果你想要让这个并非在页面加载后就自动完成的函数而需要用户点击的话
请激活这段代码，并且注释掉a.html中js部分window.onload开始的代码到js结束
并且激活btn1 button 在a.html里
$(document).ready(function () {
    $("#btn1").click(function() {
        $.ajax({
            url: "/",//请求地址
            dataType: "json",//数据格式
            type: "post",//请求方式
            async: false,//是否异步请求
            success: function (data) { //如何发送成功
                var html = JSON.stringify(data.UserInfo);
                $("#info").html(html);
            },
        })
    });
})
*/