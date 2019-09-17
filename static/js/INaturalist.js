var table=[]

function handleFiles(files) {
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
            var data = allTextLines[i].split(';');
                var tarr = [];
                for (var j=0; j<data.length; j++) {
                    tarr.push(data[j]);
                }
                lines.push(tarr);
        }
      console.log(lines);
      display(lines)
    }

    function errorHandler(evt) {
      if(evt.target.error.name == "NotReadableError") {
          alert("Canno't read file !");
      }
    }


    function display(lines){
      let htmlTable='<table>';
      table=[];
      for (let i=0; i<lines.length; i++){
        let line=lines[i][0].split(',');
        table.push(line)
        htmlTable=htmlTable+'<tr>'
        for (let e=0; e<line.length; e++){
          htmlTable=htmlTable+'<td>'+line[e]+'</td>'
        }
        htmlTable=htmlTable+'</tr>'
      }

      document.getElementById('INaturalistdata').innerHTML=htmlTable;
    }
