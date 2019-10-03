
var home=window.location.origin;



function getKMLList() {
  $.getJSON(home+"/api/listKML", function (data) {
    listElm=document.getElementById('kmlList');
    table="<table>"
    table=table+"</tr><th>remove KML file </th> <th> kml file name</th>"
    for(let i=0;i<data.length;i++){
      table=table+"<tr><td>"+data[i]+"</td></tr>"
    }
    table=table+"</table>"
    listElm.innerHTML=table;
    console.log(data);
  });
}

function uploadKML() {
  file=document.getElementById("kmlFileUpload").files[0]
  sendKML(file)
  function getAsText(fileToRead) {
    var reader = new FileReader();
    // Read file into memory as UTF-8
    reader.readAsText(fileToRead);
    // Handle errors load
    reader.onload = kmlhandler;
    reader.onerror = errormessage;
  }


  function kmlhandler(event){
    let kml = event.target.result;
    console.log(kml);
    sendKML(kml);
  }

  // getAsText(file);

  function errormessage(event) {

    popup("<h3>an error has occurred please try again</h3>")
    console.log(event);
  }

  function sendKML(kml) {
    data={};
    data.kml=kml;

    $.ajax({
      type: 'POST',
      cache: false,
      contentType: 'multipart/form-data',
      processData: false,
      url: home+'/api/putKML',
      data: data,

      success: function(data) {
          alert("Data sending was successful");
      }
  });
  }


}
