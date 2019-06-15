// content.js
var url = window.location.href;
var hostname = window.location.hostname;
var title = "";
var author = "";
var content = "";

/* CNN */
if (hostname == "edition.cnn.com") {
	title = $("h1.pg-headline").text();
	author = $("span.metadata__byline__author").text();
	
	$("section.zn-body-text > div.l-container").children().each(function () {
		if ($(this)[0].className == "el__leafmedia el__leafmedia--sourced-paragraph") {
			content = content.concat("<p>", $(this).text(), "</p>");
		} else if ($(this)[0].className == "zn-body__paragraph speakable") {
			content = content.concat("<p>", $(this).text(), "</p>");
		}
	});
	$("div.zn-body__read-all").children().each(function () {
		if ($(this)[0].className == "zn-body__paragraph") {
			if ($(this).children("h3").get(0) != null) {
				content = content.concat("<br/><h4>", $(this).text(), "</h4><br/>");
			} else {
				content = content.concat("<p>", $(this).text(), "</p>");
			}
		} else if ($(this)[0].className == "el__embedded el__embedded--fullwidth") {
			content = content.concat("<br/><div style=\"width:500px; word-wrap:break-word; font-size:80%;\">");
			content = content.concat("<img src=\"", $(this).children("div").children("div").children("img").attr("src"), "\"");
			content = content.concat(" width=500px style=\"padding-bottom:0.5em;\" /><br/>");
			content = content.concat($(this).children("div").children("div").children("div.media__caption").text(), "</div><br/>");
		}
	});
	if ($("p.zn-body__footer").text() != "") {
		content = content.concat("<p><i>", $("p.zn-body__footer").text(), "</i></p>");
	}
}

/* New York Times */
if (hostname == "www.nytimes.com") {
	title = $("h1.e1h9rw200").text();
	author = $("p.css-16vrk19").text();

	if (document.querySelector("p.css-1ifw933") != null) {
		content = content.concat("<h3>", $("p.css-1ifw933").text(), "</h3>");
	}
	$("div.css-53u6y8").children().each(function () {
		if ($(this)[0].className == "css-18icg9x evys1bk0") {
			if ($(this).children("a.css-1g7m0tk").get(0) == null) {
				content = content.concat("<p>", $(this).text(), "</p>");
			}
		} else if ($(this)[0].className == "css-ani50b eoo0vm40") {
			content = content.concat("<br/><h4>", $(this).text(), "</h4><br/>");
		}
	});
	if (document.querySelector("div.css-1yif149") != null) {
		content = content.concat("<p><i>", $("div.css-1yif149").text(), "</i></p>");
	}
}

/* Sky News */
if (hostname == "news.sky.com") {
	title = $("h1.sdc-site-component-header--h1").text().trim();
	author = "";
	content = content.concat("<br/><h4>", $("h2.sdc-site-component-header--h2").text(), "</h4><br/>");
	$("div.sdc-article-body > p").each(function () {
		content = content.concat("<p>", $(this).text(), "</p>");
	});
}

/* ABC News */
if (hostname == "abcnews.go.com") {
	title = $("header.article-header > h1").text();
	
	$("ul.authors > li").each(function () {
		var str = $(this).text();
		if (str == "By ") {
			author = str;
		} else if (str.substring(0,2) == "By") {
			str = str.substring(2,str.length).toUpperCase();
			var str1 = "";
			author = str1.concat("By ", str);
		} else if (str.substring(0,3) == "and") {
			str = str.substring(3,str.length).toUpperCase();
			str1 = "";
			str1 = str1.concat(" and ", str);
			author = author.concat(str1);
		} else {
			author = author.concat(str.toUpperCase());
		}
	});
	
	$("div.article-copy").children().each(function () {
		if ($(this)[0].tagName == "P") {
			if ($(this).children("em").get(0) != null) {
				content = content.concat("<p><em>", $(this).text(), "</em></p>");
			} else {
				if ($(this).children("strong").get(0) == null) {
					var str = $(this).text();
					if (str[0] == "\n") {
						str = str.substring(1, str.length);
					}
					if (str[str.length-1] == "\n") {
						str = str.substring(0, str.length-1);
					}
					content = content.concat("<p>", str, "</p>");
				}
			}
		} else if ($(this)[0].tagName == "FIGURE") {
			content = content.concat("<br/><div style=\"width:500px; word-wrap:break-word; font-size:80%;\">");
			content = content.concat("<img src=\"", $(this).children("div").children("picture").children("img").attr("src"), "\"");
			content = content.concat(" width=500px style=\"padding-bottom:0.5em;\" /><br/>");
			content = content.concat($(this).children("figcaption").children("span").text(), "</div><br/>");
		}
	});
}

/* Fox News */
if (hostname == "www.foxnews.com") {
	title = $("h1.headline").text();
	
	author = "By ";
	$("div.author-byline > span > span").each(function () {
		author = author.concat($(this).text());
	});
	
	$("div.article-body").children().each(function () {
		if ($(this)[0].tagName == "P") {
			if ($(this).children("i").get(0) != null || $(this).children("em").get(0) != null) {
				content = content.concat("<p><i>", $(this).text(), "</i></p>");
			} else if ($(this).children("strong").get(0) == null && $(this).children("a").children("strong").get(0) == null) {
				content = content.concat("<p>", $(this).text(), "</p>");
			}
		} else if ($(this)[0].className == "image-ct inline") {
			content = content.concat("<br/><div style=\"width:500px; word-wrap:break-word; font-size:80%;\">");
			content = content.concat("<img src=\"", $(this).children("div").children("picture").children("img").attr("src"), "\"");
			content = content.concat(" width=500px style=\"padding-bottom:0.5em;\" /><br/>");
			content = content.concat($(this).children("div.caption").text(), "</div><br/>");

		}
	});
}

/* China Post */
if (hostname == "chinapost.nownews.com") {
	title = $("h1.entry-title").text();

	var paragraph = "";
	$("div.td-post-content > p").each(function () {
		paragraph = $(this).text()
		content = content.concat("<p>", paragraph, "</p>");
	});
	if (paragraph.substring(0,2) == "By") {
		content = content.substring(0,content.length-(paragraph.length+7));
		author = paragraph;
	}
}

/* Taipei Times */
if (hostname == "www.taipeitimes.com") {
	title = $("h1.title").text();
	if (document.querySelector("div.reporter") != null) {
		author = $("div.reporter").text();
	}
	$("div.text > p").each(function () {
		content = content.concat("<p>", $(this).text(), "</p>");
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

chrome.runtime.sendMessage(a_json);