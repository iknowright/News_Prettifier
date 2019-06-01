// content.js

var firstHref = $("h1").text();

console.log(firstHref);
console.log(firstHref);
console.log(firstHref);
console.log(firstHref);


var url = window.location.href;
var hostname = window.location.hostname;
var title = "";
var author = "";
var content = "";

if (hostname == "edition.cnn.com") {
  title = $("h1.pg-headline").text();
  author = $("span.metadata__byline__author").text();
  content = content.concat("<p>", $("p.zn-body__paragraph:not(.zn-body__footer)").text(), "</p>");
  $("div.zn-body__paragraph").each(function () {
    if ($(this).children("h3").get(0) != null) {
      content = content.concat("<h3>", $(this).text(), "</h3>");
    } else {
      content = content.concat("<p>", $(this).text(), "</p>");
    }
  });
}

console.log(hostname);
console.log(url);
console.log(title);
console.log(author);
console.log(content);

a_json = {
  "title": title,
  "author": author,
  "content": content,
  "origin": url
}

console.log("CONTENT IS READY (content.js)");

// chrome.storage.local.set({'content': 'ready'}, function() {
//   console.log("-------set the key--------");
// });

// chrome.runtime.sendMessage({
//   from: 'content',
//   subject: 'contentReady',
// }, function(response){
//   console.log(response.from);
// });


// a_json = {
//     "title": "test title",
//     "author": "the author",
//     "content": "the content",
//     "origin": " no where"
// }

// chrome.runtime.sendMessage({greeting: "hello"}, function(response) {
chrome.runtime.sendMessage(a_json);

//send msg + stringify