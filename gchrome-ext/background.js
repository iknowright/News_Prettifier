// When the extension is installed or upgraded ...
chrome.runtime.onInstalled.addListener(function () {
    // Replace all rules ...
    chrome.declarativeContent.onPageChanged.removeRules(undefined, function () {
        // With a new rule ...
        chrome.declarativeContent.onPageChanged.addRules([
            {
                // That fires when a page's URL contains a 'g' ...
                conditions: [
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlContains: 'g' },
                    })
                ],
                // And shows the extension's page action.
                actions: [new chrome.declarativeContent.ShowPageAction()]
            }
        ]);
    });
});

chrome.extension.onConnect.addListener(function(port) {
    console.log("Connected .....");
    port.onMessage.addListener(function(msg) {
        console.log("message recieved " + msg);
        chrome.tabs.query({ active: true, currentWindow: true },(tabs) => {
            var currentTab = tabs[0];
            port.postMessage(currentTab.url);
        });
        if(msg === "open the tab") {
            chrome.tabs.create({"url": "index.html"});
        }
    });
})

a_json = {
    "title": "test title",
    "author": "the author",
    "content": "the content",
    "origin": " no where",
}
console.log("here");

$.ajax({
    url: 'https://0f952153.ngrok.io/article',
    type: 'post',
    dataType: 'json',
    contentType: 'application/json',
    success: function (data) {
        console.log(data);
    },
    data: JSON.stringify(a_json),
});
