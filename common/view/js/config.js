/*=========================================公共变量信息-开始=========================================*/
let publicUrl = "http://127.0.0.1:8003/api/v1"
let login_url = "/login"
let url_404 = "/404"
let url_500 = "/500"
let expire_date = (new Date().getTime() / 1000) + 86400
/*=========================================基础数据变量信息-开始=========================================*/
let http_url_login = publicUrl + '/admin/login'
let http_url_menu = publicUrl + '/admin/menu'
let http_url_device_list = publicUrl + '/device/list'


let http_url_province = publicUrl + '/region/province'
let http_url_city = publicUrl + '/region/city'
let http_url_area = publicUrl + '/region/area'
let http_url_community = publicUrl + '/device/community/no/page/list'
let http_url_wing_room = publicUrl + '/device/cabinet/wing_room'

/*=========================================公共方法-开始=========================================*/

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

function scrollTop() {
    //document.body.scrollIntoView()
    $("body,html").animate({scrollTop: 0}, 200);
}

function sendApi(url, data, sendType, token) {
    if (token == undefined) {
        token = ""
    }
    let result = {};
    $.ajax({
        type: sendType,
        url: url,
        dataType: 'JSON',
        data: JSON.parse(data),
        headers: {"token": token},
        async: false,
        beforeSend: function () {
        },
        success: function (res) {
            result = res
            jump_login(result)
        },
        complete: function () {
            hide_loading()
        },
        error: function (res) {
            hide_loading()
            result = res
            layer.msg("服务器繁忙！接口返回空数据", {anim: 6});
            return false
        }
    })
    return result
}

function jump_login(ret) {
    let hash = parent.location.hash;
    if (ret == 40004) {
        layer.msg('您还没登录，请先去登录', {shade: 0.5, time: 3000}, function () {
            window.parent.frames.location.href = login_url + hash
        });
    }
    return false
}

//时间戳格式化，v:需要转义的时间戳毫秒，format：需要转为的时间格式，如：yyyy-MM-dd HH:mm:ss
function formatDateCommon(v, format) {
    if (v == 0) {
        return "";
    }
    var dateV = new Date(v);
    var year = dateV.getFullYear();
    var month = dateV.getMonth() + 1;
    month = month < 10 ? '0' + month : month;
    var date = dateV.getDate();
    date = date < 10 ? ("0" + date) : date;
    var hour = dateV.getHours();
    hour = hour < 10 ? ("0" + hour) : hour;
    var minute = dateV.getMinutes();
    minute = minute < 10 ? ("0" + minute) : minute;
    var second = dateV.getSeconds();
    second = second < 10 ? ("0" + second) : second;
    var str1 = year + "-" + month + "-" + date
    var str2 = hour + ":" + minute + ":" + second;
    var str
    if ("yyyy-MM-dd" == format) {
        str = str1;
    } else {
        str = str1 + " " + str2
    }
    return str;
}

function formatDateToTimestamp(v) {
    let t = new Date(v)
    return t.getTime() / 1000
}

/*=========================================公共方法-结束=========================================*/

