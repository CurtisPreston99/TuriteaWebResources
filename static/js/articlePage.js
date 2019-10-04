

window.onload=function(){
  $.getJSON('./api/lastArticle', {}, function(data){
    // console.log(data);
    articleTable="<table>"
    for(let i=0;i<data.length;i++){
      console.log(data[i]);
      articleTable=articleTable+"<tr>"
      articleTable=articleTable+"<td>"
      articleTable=articleTable+"<td>"
      articleTable=articleTable+data[i].sum;
      articleTable=articleTable+"</td>"

      articleTable=articleTable+"</tr>"
    }
    articleTable=articleTable+"</table>"
    document.getElementById('articleTable').innerHTML=articleTable;
    // console.log(articleTable);
  });
}
