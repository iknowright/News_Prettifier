document.getElementById("Button").disabled = true;

var bkg = chrome.extension.getBackgroundPage();

var port = chrome.extension.connect({
    name: "Sample Communication"
});



port.postMessage("Hi BackGround");
chrome.runtime.onMessage.addListener(function (msg) {
    bkg.console("RECEIVED MSG :" + msg + "(popup.js)");
    //if ((msg.from === 'background') && (msg.subject === 'contentReady')) {
        // bkg.console.log("message recieved " + msg);
        bkg.console("CONTENT IS READY (popup.js)");
        document.getElementById("Button").enabled = true;
    //}
});


// var content_ready = false;
// while(!content_ready) {
//     chrome.storage.local.get("content", function(data) {
//         if(typeof data.content == "undefined") {
//             console.log("CONTENT NOT READY");
//             continue;
//         } else {
//             content_ready = true;
//             // port.postMessage("open the tab");
//             break;
//         }
//     });
// } 
var activateButton = false;

setInterval(function() {  
    chrome.storage.local.get('content', function(result) {
        var state = result.content;
        bkg.console.log('RESULT: ' + state);
        if(state === "ready") {
            bkg.console.log("BUTTON SHOULD BE ACTIVATED");
            document.getElementById("Button").disabled = false;
            document.getElementById("Button").enabled = true;
            chrome.storage.local.set({'content': 'unready'}, function() {
                console.log("-------set the key--------");
            });
            // document.getElementById("Button").enabled = true;
        }
    });
},2000);



var gothtml;
$(document).ready(function () {
    $("button").on("click", function() { 
        // document.getElementById("Button").enabled = true;
        port.postMessage("open the tab");
        // $.ajax({
        //     url: "http://google.com",
        //     success: function(data){
        //         gothtml = data;  
        //         port.postMessage(data);
        //     }
        // });
        // port.postMessage("open the tab");
    })
});



