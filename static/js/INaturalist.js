var table=[]
function handleFiles(files) {

  table=[]
      // Check for the various File API support.
      if (window.FileReader) {
          // FileReader are supported.
          getAsText(files[0]);
      } else {
          alert('FileReader are not supported in this browser.');
      }
    }

    function getAsText(fileToRead) {
      var reader = new FileReader();
      // Read file into memory as UTF-8
      reader.readAsText(fileToRead);
      // Handle errors load
      reader.onload = loadHandler;
      reader.onerror = errorHandler;
    }

    function loadHandler(event) {
      var csv = event.target.result;
      processData(csv);
    }

    function processData(csv) {
        var allTextLines = csv.split(/\r\n|\n/);
        var lines = [];
        for (var i=0; i<allTextLines.length; i++) {
            var data = allTextLines[i].split(',');
                var tarr = [];
                for (var j=0; j<data.length; j++) {
                    tarr.push(data[j]);
                }
                lines.push([true].concat(tarr));
        }
      console.log(lines);

      table=lines;
      display(table)

    }

    function errorHandler(evt) {
      if(evt.target.error.name == "NotReadableError") {
          alert("Canno't read file !");
      }
    }

    function enable(x){
      table[x][0]=!table[x][0];

      display(table);
    }


    function display(table){
      let htmlTable='<table>';
      for (let i=0; i<table.length; i++){
        line=table[i];
        htmlTable=htmlTable+'<tr>'
        for (let e=0; e<line.length; e++){
          if(e==0){
            if(i==0){
              htmlTable=htmlTable+'<th>'+'add'+'</th>'
            }else{
              if(!table[i][e]){
                htmlTable=htmlTable+'<td><button onclick="enable('+i+')">enable</button></td>'

              }else{
            htmlTable=htmlTable+'<td><button onclick="enable('+i+')">remove</button></td>'
            }
          }

        }else{
          if(e>0){
          if(i==0){
            htmlTable=htmlTable+'<th>'+line[e]+'</th>'

          }else{
          htmlTable=htmlTable+'<td>'+line[e]+'</td>'
        }
      }
    }}
        htmlTable=htmlTable+'</tr>'
      }
      document.getElementById('INaturalistdata').innerHTML=htmlTable;
    }




    function tabletoPins(){
      console.log(table);
      let data=[]
      var rows={}

      for(let i=0;i<table[0].length;i++){
        rows[table[0][i]]=i
      }

      // console.log(rows);

      for (let i=1; i<table.length; i++){
        if(table[i][0]){
          pin={}
          pin["name"]=table[i][rows["common_name"]]
          pin["tag_type"]="zoo";
          pin["lat"]=table[i][rows["latitude"]]

          pin["lon"]=table[i][rows["longitude"]]
          pin["color"]="#FFFFFF"
          pin["description"]=pinGenerateDiscription(table[i],rows)
          pin["time"]=2;
          data.push(pin);
        }
      }
      console.log(data);
      let pins="{data: \""+JSON.stringify(data)+"\"}"
      let n=data.length;
      console.log(pins);
      $.post("../api/addPins?num=\""+n.toString(16)+"\"",pins,function(){
        console.log("posted");
      });
    }

    function pinGenerateDiscription(line,rowMap) {
      let html="<p>test</p>"
      // html=html+"<a href="+line[rowMap["url"]]+' target="_blank">look at me on inaturalist.nz</a>'
      // html=html+"</p><p>"
      // if(line[rowMap["image_url"]]!==""){
      //   html=html+"<img src="+line[rowMap["url"]]+">"
      // }
      //
      // if(line[rowMap["image_url"]]!==""){
      //   html=html+"<img src="+line[rowMap["url"]]+">"
      // }
      //
      // html=html+"</p><p>"
      // html=html+line[rowMap["description"]]
      // html=html+"</p><p>"
      //
      // httable="<table>"
      // for(let key in rowMap){
      //   httable=httable+='<tr><td>'+key+'</td><td>'+line[key]+'</td></tr>'
      // }
      //
      // httable=httable+"</table>"
      // html=html+httable


      return html
    }
