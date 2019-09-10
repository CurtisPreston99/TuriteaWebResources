var home=window.location.origin;


window.onload = function(){
$('#summernote').summernote({
height: $(window).height()/2   //set editable area's height
});
}


function uploadData(){
  article={}
  article.sum=$('#summernote').summernote('code');
  article.name=document.getElementById('name').value;
  console.log(article)
  $.post(home+"/api/addArticle?num=1",article,function(){console.log("posted");});
}
