

var home=window.location.origin;

var articles;

function loadeditor(){
$('#summernote2').summernote({
height: 500,   //set editable area's height
codemirror: { // codemirror options
theme: 'monokai'
}
});
}


function uploadFrag() {
  frag=document.getElementById('articleSum').value;


  $.post(home+"/api/addArticleFragment","{\"data\":\""+frag+"\"}",function(data){
    console.log(data);
    popup("<h4> post sussesful</h4>");

  });
}

function uploadArticleData(){
  console.log($('#summernote2').summernote('code'));
  article={}
  article.sum=$('#summernote2').summernote('code');
  console.log(article)

  send={}
  send.data='['+JSON.stringify(article)+']'

  $.post(home+"/api/addArticle?num=1",send,function(data){console.log(data);
  if(data=='0 -1'){
    popup("<h3>there has been a database error</h3>");
  }else{
    popup("<h4> post sussesful</h4>");
  }
  }).fail(function(err) {
    console.log(err);
    popup("<h3>there has been an error</h3>");
  });
  }
