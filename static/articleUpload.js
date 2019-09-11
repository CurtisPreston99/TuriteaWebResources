var home=window.location.origin;


window.onload = function(){
$('#summernote').summernote({
height: $(window).height()/2   //set editable area's height
});
}


function uploadData(){
  article={}
  article.sum=$('#summernote').summernote('code');
  console.log(article)

  send={}
  send.data='['+JSON.stringify(article)+']'

  $.post(home+"/api/addArticle?num=1",send,function(){console.log("posted");});
}
