var last = 0;
var min = 0xffffffffffffffffffffffffffffffff;

function loadArticleList() {
    var list = $("#articleList");
    let data = {begin:last, num:10};
    last += 10;
    $.get("../api/lastArticle", data, function (r) {
        let articles = JSON.parse(r);
        articles.sort(function (a, b) {
            return b.id - a.id;
        });
        if (articles.length === 0) {
            let n = $("#next");
            n.attr("disabled", "true");
            $(n[0]).text("no More Article");
            return;
        }
        for (let i = 0; i < articles.length; i ++) {
            if (articles[i].id < min) {
                min = articles[i].id;
                list.append($(`<li>
                    <a href="../article/{0}">
                    <p>{1}</p>
                    </a>
                </li>`.format(articles[i].id, articles[i].sum)));
            } else {
                message("Sorry", "No more article");
                $("#next").hide();
            }
        }
    });
}
