var articles;

function openArticle(id){
  let article;
  if(article==null){

    $.ajax({
    url : './api/lastArticle',
    type : "get",
    async: false,
    success : function(data) {
        articles=JSON.parse(data);

    },
    error: function() {
       connectionError();
    }
  });
  }
  for(let i=0;i<articles.length;i++){
    console.log(articles[i].id);
    console.log(id);

    if(articles[i].id===parseInt(id)){
      console.log(id);
      article=articles[i];
      displayArticle(article);
      break;
    }
  }
}


function displayArticle(article) {

  console.log(article);

  document.getElementById('articleTable').style.display='none'
  document.getElementById('articleLarge').style.display=''

  document.getElementById('articleLarge').innerHTML=article.sum;
}

function displayArticleList() {


    document.getElementById('articleTable').style.display=''
    document.getElementById('articleLarge').style.display='none'
  $.ajax({
  url : './api/lastArticle',
  type : "get",
  async: false,
  success : function(data) {
      articles=JSON.parse(data);

  },
  error: function() {
     connectionError();
  }
});


articleTable="<table class='articleList'>"
for(let i=0;i<articles.length;i++){
  console.log(articles[i]);
  articleTable=articleTable+"<tr>"
  articleTable=articleTable+"<td><div onclick=openArticle('"+articles[i].id+"')>"
  articleTable=articleTable+articles[i].sum;

  articleTable=articleTable+"</div></td>"

  articleTable=articleTable+"</tr>"
}
articleTable=articleTable+"</table>"
document.getElementById('articleTable').innerHTML=articleTable;

}

window.onload=function(){
  var url = new URL(document.URL);
  console.log(url);
  var aid=url.searchParams.get("aid");

  if(aid==null){
  displayArticleList();
}else{
  openArticle(aid);
}
}
