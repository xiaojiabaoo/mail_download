//获取请求地址中携带的参数
function getRequestParam() {
    let obj = {};
    let arr = window.location.search.slice(1).split("&");
    for (let i = 0, len = arr.length; i < len; i++) {
        let nv = arr[i].split("=");
        obj[decodeURIComponent(nv[0]).toLowerCase()] = decodeURIComponent(nv[1]);
    }
    return obj;
}

//全局加载提示
function show_loading() {
    this.loadingMsg = layer.msg('加载中，请稍后...',
        {
            /*icon: 16,*/
            id: 'layer_loading' + new Date().getTime(),
            zIndex: layer.zIndex,
            offset: '100px',
            shade: [0.1],
            scrollbar: false,
            time: 0,
            success: function (layero) {
                layer.setTop(layero);
            }
        });
}

function hide_loading() {
    layer.close(this.loadingMsg);
}

function sendGetApi(url,sendType,token) {
    let result = {};
    $.ajax({
        url: url, // 请求的URL
        method: 'GET', // 请求方法
        dataType: sendType, // 预期的返回数据类型
        beforeSend: function () {
            hide_loading()
        },
        success: function(response) {
            hide_loading()
            result = response
            return false
        },
        error: function(xhr, status, error) {
            hide_loading()
            result = error
            layer.msg("服务器繁忙！接口返回空数据", {anim: 6});
            return false
        }
    });
    return result
}

function sendPostApi(url, data, sendType, token) {
    if (token == undefined) {
        token = ""
    }
    let result = {};
    $.ajax({
        type: sendType,
        url: url,
        dataType: 'JSON',
        data: data,
        headers: {"token": token},
        contentType: 'application/json; charset=utf-8', // 设置请求内容的类型为JSON
        async: true,
        beforeSend: function () {

        },
        success: function (res) {
            result = res
        },
        complete: function () {

        },
        error: function (res) {

            result = res
            layer.msg("服务器繁忙！接口返回空数据", {anim: 6});
            return false
        }
    })
    return result
}

