
var home=window.location.origin;



function getKMLList() {
  $.getJSON(home+"/api/listKML", function (data) {
    console.log(data);
    listElm=document.getElementById('kmlList');
    table="<table>"
    table=table+"</tr><th>remove KML file </th> <th> kml file name</th>"
    for(let i=0;i<data.length;i++){
      table=table+"<tr><td><button onclick=\'removeKML(\""+data[i]+"\")\'>remove</button></td><td>"+data[i]+"</td></tr>"
    }
    table=table+"</table>"
    listElm.innerHTML=table;
  });
}



function removeKML(name) {
  console.log("removeing: "+name);
  $.post(home+"/api/deleteKML&name="+name,name,function(){
    console.log("removed");
    popup("<h4>remove sussesful</h4>");
  });
}
