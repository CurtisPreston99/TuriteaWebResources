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
    $.get('../api/fragment?id=' + home, function (r) {
        console.log(r);
        let f = JSON.parse(r);
        console.log(f);
        let content = f["content"];
        let res = f["res"];

        for (let i = 0; i < res.length; i++) {
            if (res[i]["type"] === "f") {
                eachFragment($("<div></div>"), res[i]["id"]);
            }
        }
        console.log(content);
        homeDiv.append($("<p></p>").append(content));
        homeDiv.append($("<span hidden></span>").text(r));
    });
}

function eachFragment(node, id) {
    $.get('../api/fragment?id=' + id, function (r) {
        let f = JSON.stringify(r);
        let content = f["content"];
        let res = f["res"];
        node.append($("<p></p>").append(content));
        node.append($("<span hidden></span>").text(r));
        home.append(node);
    });
}
var articleEditor;
var information = null;
var article = null;
function loadArticleNote() {
    articleEditor =$("#articleSummerNote");
    articleEditor.summernote({
        height: 500,   //set editable area's height
        codemirror: { // codemirror options
            theme: 'monokai'
        },
        placeholder: "feel free to edit"});
    article = JSON.parse(localStorage.getItem("editArticle"));
    if (article) {
        $("#articleTitle").text("edit article");
        console.log(article.content);
        articleEditor.summernote("code", article.content);
        $("#articleSum").val(article.sum);
        information = JSON.parse(article.home);
        console.log(information);
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
    let delImage = [];
    if (information !== null) {
        let fid = information["id"];
        console.log(fid);
        let res = information["res"];
        // $.get("../api/delete?type=0&id="+fid);
        console.log("../api/delete?type=0&id="+fid);
        for (let i = 0; i < res.length; i++) {
            if (res[i]["type"] === "m") {
                delImage.push(res[i]["m"]["uid"].toString(16));
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
            img.attr("src", "../api/getImage?id=%x");
        }
    }
    console.log(delImage);
    combination.images = JSON.stringify(datas);
    combination.imageNum = datas.length;
    combination.articles = JSON.stringify([{"sum":$("#articleSum").val()},]);
    combination.content = help.html();
    //console.log(help);
    // for (let i = 0; i < help.length; i++) {
    //     combination.content += $(help[i]).html();
    // }
    if (information !== null) {
        let id = article["id"];
        combination.aid = id;
        console.log(combination);

        for (let i = 0; i < delImage.length; i++) {
            let imgId = delImage[i];
            if (combination.content.indexOf("../api/getImage?id=" + imgId) < 0) {
                // $.get("../api/delete?type=1&id=" + imgId)
                // $.get("../api/delete?type=3&id=" + imgId)
                console.log("../api/delete?type=3&id=" + imgId);
                console.log("../api/delete?type=1&id=" + imgId);
            }
        }
        $.post("../api/updateArticle", combination, function (r) {
            localStorage.setItem("editArticle", null);
            information = null;
            // console.log("success");
            window.location.href = "../article/" + id;
        }).fail(function (r) {
            console.log(r);
        });
    } else {
        console.log(combination);
        $.post("../api/addArticleWithImage", combination, function (r) {
            localStorage.setItem("editArticle", null);
            // console.log("success");
            window.location.href = "../article/" + r;
        }).fail(function (r) {
            console.log(r);
        })
    }


}
