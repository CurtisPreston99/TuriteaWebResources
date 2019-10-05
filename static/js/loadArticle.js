var home = $("#main");

function loadMainFragment() {
    let c = document.cookie;
    let homeDiv = $("#home");
    let first = c.indexOf("home");
    if (first < 0) {
        let fragment = $("#at404");
        fragment.show();
        homeDiv.hide();
        return
    }
    let end = c.substring(first, c.length).indexOf(";");
    let home = parseInt(c.substring(first + 5, end===-1?c.length:end));
    if (!home) {
        let fragment = $("#at404");
        fragment.show();
        homeDiv.hide();
        return
    }
    console.log(home);
    $.get('../api/fragment?information=' + home, function (r) {
        console.log(r);
        let f = JSON.parse(r);
        console.log(f);
        let content = f["content"];
        let res = f["res"];

        for (let i = 0; i < res.length; i++) {
            if (res[i]["type"] === "f") {
                eachFragment($("<div></div>"), res[i]["information"]);
            }
        }
        console.log(content);
        homeDiv.append($(content));
        homeDiv.append($("<span hidden></span>").text(r));
    });
}

function eachFragment(node, id) {
    $.get('../api/fragment?information=' + id, function (r) {
        let f = JSON.stringify(r);
        let content = f["content"];
        let res = f["res"];
        for (let i = 0; i < res.length; i++) {
            if (res[i]["type"] === "f") {
                eachFragment(node.clone(true), res[i]["information"]);
            }
        }
        node.append($(content));
        home.append(node);
    });
}
var articleEditor;
var information = null;
function loadArticleNote() {
    articleEditor =$("#articleSummerNote");
    articleEditor.summernote({
        height: 500,   //set editable area's height
        codemirror: { // codemirror options
            theme: 'monokai'
        },
        placeholder: "feel free to edit"});
    let article = JSON.parse(localStorage.getItem("editArticle"));
    if (article) {
        $("#articleTitle").text("edit article");
        articleEditor.summernote("code", article.content);
        $("#articleSum").val(article.sum);
        information = article.id;
        // console.log(information);
    }
}

function editArticle(self) {
    // console.log(self)
    let div = $(self);
    let a = {};
    a.sum = $("#summary").text();
    a.content = div.children("p").html();
    a.home = div.children("span").text();
    a.id = $("#articleId").text();
    // console.log(a)
    localStorage.setItem("editArticle", JSON.stringify(a));
    window.location.href = "../html/settings.html#tabs-3"
}

function submitArticle() {
    if (information !== null) {
        let f = JSON.parse(information);
        let fid = f["fid"];
        let res = f["res"];
        // $.get("../api/delete?type=0&id="+fid);
        console.log("../api/delete?type=0&id="+fid);
        for (let i = 0; i < res.length; i++) {
            if (res[i]["type"] === "m") {
                // $.get("../api/delete?type=1&id=" + res[i]["m"]["uid"].toString(16))
                console.log("../api/delete?type=1&id=" + res[i]["m"]["uid"].toString(16))
            }
        }

    }
    let combination = {};
    let code = articleEditor.summernote('code');
    // code = code.replace("<img", "<myImage");
    let help = $("<p></p>").append($.parseHTML(code));
    //console.log(help.html());
    let images = help.find("img");
    //console.log(images);
    let datas = [];
    for (let i = 0; i < images.length; i++) {
        let image = images[i];
        let filName = image.dataset.filename;
        let img = $(image);
        let src = img.attr("src");
        if (src.startsWith("data:")) {
            datas.push({"image":src.split(",")[1], "title":filName});
            img.attr("src", "../api/getImage?information=%x");
        }
    }
    //console.log(help.html());
    combination.images = JSON.stringify(datas);
    combination.imageNum = datas.length;
    combination.articles = JSON.stringify([{"sum":$("#articleSum").val()},]);
    combination.content = help.html();
    //console.log(help);
    // for (let i = 0; i < help.length; i++) {
    //     combination.content += $(help[i]).html();
    // }
    console.log(combination);
    if (information !== null) {
        $.post("../api/updateArticle", combination, function (r) {
            localStorage.setItem("editArticle", null);
            // console.log("success");
            window.location.href = "../article/" + r;
        }).fail(function (r) {
            console.log(r);
        })
    } else {
        $.post("../api/addArticleWithImage", combination, function (r) {
            localStorage.setItem("editArticle", null);
            // console.log("success");
            window.location.href = "../article/" + r;
        }).fail(function (r) {
            console.log(r);
        })
    }

    // information = null;
}
