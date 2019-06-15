article_id = "";
chrome.storage.local.set({'content': 'unready'}, function() {
    console.log("-------set the key--------");
  });
// When the extension is installed or upgraded ...
chrome.runtime.onInstalled.addListener(function () {
    // Replace all rules ...
    chrome.declarativeContent.onPageChanged.removeRules(undefined, function () {
        // With a new rule ...
        chrome.declarativeContent.onPageChanged.addRules([
            {
                conditions: [
                    // new chrome.declarativeContent.PageStateMatcher({
                    //     pageUrl: { pathContains: 'g'}
                    // })
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.cnn\.com/[0-9]+.+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.nytimes\.com/[0-9]+.+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.sky\.com/story.+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.abcnews\.go\.com/(US|International|Politics|Health|Entertainment|Sports|Business|Technology|Lifestyle)/.+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.foxnews\.com/(world|opinion|politics|science|us|entertainment|lifestyle|tech|health).+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.chinapost\.nownews\.com/[0-9]+-[0-9]+' }
                    }),
                    new chrome.declarativeContent.PageStateMatcher({
                        pageUrl: { urlMatches: '.+\.taipeitimes\.com/News.+' }
                    }),
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
            chrome.tabs.create({"url": "https://news-prettifier.herokuapp.com/index/" + article_id});
        }
    });
})

chrome.runtime.onMessage.addListener(function(msg) {
    console.log(msg);
    // if ((msg.from === 'content') && (msg.subject === 'contentReady')) {
    //     // Enable the page-action for the requesting tab.
    //     console.log("CONTENT MESSAGE FROM BACKGROUND, SENDING MSG TO POPUP");
    //     chrome.runtime.sendMessage("CONTENT_READY");
    // }
    // else {
        request = $.ajax({
            url: "https://news-prettifier.herokuapp.com/article",
            type: 'POST',
            dataType: 'json',
            contentType: 'application/json',
            success: function (data) {
                console.log("Content: " + JSON.stringify(data));
                article_id = data.article_id;
                console.log("Article ID: " + article_id);
            },
            data: JSON.stringify(msg)
        });
        request.done(function (response, textStatus, jqXHR){
            // Log a message to the console
            console.log("Hooray, it worked!");
            chrome.storage.local.set({'content': 'ready'}, function() {
                console.log("-------set the key--------");
            });
            chrome.storage.local.set({'repeated': 'true'}, function() {
                console.log("-------set the key--------");
            });
        });
        request.fail(function (jqXHR, textStatus, errorThrown){
            // Log the error to the console
            console.error(
                "The following error occurred: "+
                textStatus, errorThrown
            );
            console.warn(jqXHR.responseText);
        });
    // }
});

// a_json = {
//             "title": "test title",
//             "author": "the author",
//             "content": "the content",
//             "origin": " no where",
//         }

//data = rcv_msg...

// $.ajax({
//     url: "https://c69a1486.ngrok.io/article",
//     type: 'post',
//     dataType: 'json',
//     contentType: 'application/json',
//     success: function (data) {
//         console.log(data);
//     },
//     data: JSON.stringify(a_json)
// });

// $.ajax({
//     type: "POST",
//     url: "127.0.0.1:8000/article",
//     data: {
//         "title": "test title",
//         "author": "the author",
//         "content": "the content",
//         "origin": " no where",
//         "username": ""
//     },
//     success: function(data){
//         console.log(data);
//     }
// });