var bkg = chrome.extension.getBackgroundPage();

var port = chrome.extension.connect({
    name: "Sample Communication"
});
port.postMessage("Hi BackGround");
port.onMessage.addListener(function (msg) {
    bkg.console.log("message recieved " + msg);
});

var gothtml;
$(document).ready(function () {
    $(".btn-pretty").on("click", function() {
        // $.ajax({
        //     url: "http://google.com",
        //     success: function(data){
        //         gothtml = data;  
        //         port.postMessage(data);
        //     }
        // });
        port.postMessage("open the tab");
    })
});