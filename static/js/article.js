

var home=window.location.origin;



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
    popup("<h4> post sussesful</h4>");
  });
  }
