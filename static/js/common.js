// 缁存姢web绔痗ookie key鍊兼灇涓惧垪琛�
var CookieKeyEnum= {
    USER_IP: "user_ip",   // 鍚庡彴绠＄悊鑿滃崟
    USER_CITY: "user_city"    // 鍓嶅彴header鑿滃崟
}


// 缃《
// 鑾峰彇缃《瀵硅薄
var obj = document.getElementById('top-event');
var scrollTop = null;

// 缃《瀵硅薄鐐瑰嚮浜嬩欢
obj.onclick = function() {
    var timer = setInterval(function() {
        window.scrollBy(0, -100);
        if (scrollTop == 0)
            clearInterval(timer);
    }, 2);
}

// 绐楀彛婊氬姩妫€娴�
window.onscroll = function() {
    scrollTop = document.documentElement.scrollTop || window.pageYOffset || document.body.scrollTop;
    obj.style.display = (scrollTop >= 300) ? "block" : "none";
}

layui.use('element', function(){
    var $ = layui.jquery
        ,element = layui.element; //Tab鐨勫垏鎹㈠姛鑳斤紝鍒囨崲浜嬩欢鐩戝惉绛夛紝闇€瑕佷緷璧杄lement妯″潡
});

// 鎻愬彇鍦板潃鏍忓弬鏁�
function GetUrlParamByName(key) {
    // 鑾峰彇鍙傛暟
    var url = window.location.search;
    // 姝ｅ垯绛涢€夊湴鍧€鏍�
    var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
    // 鍖归厤鐩爣鍙傛暟
    var result = url.substr(1).match(reg);
    //杩斿洖鍙傛暟鍊�
    return result ? decodeURIComponent(result[2]) : "";
}

$(window).scroll(function () {
    var scrollHeight = $(document).height();
    var scrollTop = $(window).scrollTop();
    var $windowHeight = $(window).innerHeight();
    scrollTop > 50 ? $("#scrollUp").fadeIn(200).css("display","block") : $("#scrollUp").fadeOut(200);
    $bottomTools.css("bottom", scrollHeight - scrollTop > $windowHeight ? 40 : $windowHeight + scrollTop + 40 - scrollHeight);
});

$('#scrollUp').click(function (e) {
    alert(1);
    e.preventDefault();
    $('html,body').animate({ scrollTop:0});
});

// 鑾峰彇銆恘,m銆戝尯闂寸殑闅忔満姝ｆ暣鏁�
function getRandom(n, m){
    var random = Math.floor(Math.random()*(m-n+1)+n);
    return random;
}



//------------- my web cookie deal method define start ------------- //
function setCookieSetDate(key, value, n) {
    var oDate = new Date();
    oDate.setDate(oDate.getDate() + 1); // 榛樿web绔痗ookie鐢熷懡鍛ㄦ湡涓�1灏忔椂
    document.cookie = key + '=' + value + ';expires=' + oDate;
}
function setCookieDafaultOneDate(key, value) {
    setCookieSetDate(key, value, 1);
}
function removeCookie(key) {
    setCookieSetDate(key, '', -1);//杩欓噷鍙渶瑕佹妸Cookie淇濊川鏈熼€€鍥炰竴澶╀究鍙互鍒犻櫎
}
function getCookie(key) {
    var cookieString = document.cookie == "" ? window.parent.document.cookie : document.cookie;
    var cookieArr = cookieString.split(';');
    for(var i = 0; i < cookieArr.length; i++) {
        var arr = cookieArr[i].split('=');
        if(arr[0] == key) {
            return arr[1];
        }
    }
    return false;
}
//------------- my web cookie deal method define end ------------- //


// ---------------  缃戠珯鏀剧疆鍒版闈�  ----------------//
function toDesktop(sUrl, sName) {
    try {
        var WshShell = new ActiveXObject("WScript.Shell");
        var oUrlLink = WshShell.CreateShortcut(WshShell
                .SpecialFolders("Desktop")
            + "\\" + sName + ".url");
        oUrlLink.TargetPath = sUrl;
        oUrlLink.Save();
        layer.msg("鎴愬姛鍒涘缓妗岄潰蹇嵎鏂瑰紡! 鎰熻阿鎮ㄧ殑鏀寔");
    } catch (e) {
        layer.msg("褰撳墠IE瀹夊叏绾у埆涓嶅厑璁告搷浣滄垨鎮ㄧ殑娴忚鍣ㄤ笉鏀寔姝ゅ姛鑳斤紒<br/>鍙互鍏堟敹钘忥紝浠ュ厤鍚庨潰鎵句笉鍒癧杩欓噷闇插嚭灏村艾鐨勮〃鎯匽",{time:8*1000});
    }
}