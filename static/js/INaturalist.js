var table = [];
var header = {};


function handleFiles(files) {

    table = [];
    // Check for the various File API support.
    if (window.FileReader) {
        // FileReader are supported.
        getAsText(files[0]);
    } else {
        error("Sorry", "try to use a higher version browser")
    }
}

function getAsText(fileToRead) {
    var reader = new FileReader();
    // Handle errors load
    reader.onload = loadHandler;
    reader.onerror = errorHandler;
    // Read file into memory as UTF-8
    reader.readAsText(fileToRead);
}

function loadHandler(event) {
    let sheet = event.target.result;
    console.log(sheet);
    table = CSVToArray(sheet, ",");
    console.log(table);
    var rows = {};

    for (let i = 0; i < table[0].length; i++) {
        rows[table[0][i]] = i
    }
    header = rows;
    let newTable = [];
    for (let i = 0; i < table.length; i++) {
        if (table[i][header["common_name"]] !== "" && table[i][header["common_name"]] !== null) {
            if (table[i][header["latitude"]] !== "" && table[i][header["latitude"]] !== null) {
                if (table[i][header["longitude"]] !== "" && table[i][header["longitude"]] !== null) {
                    newTable.push(table[i]);
                }
            }
        }
    }
    table = newTable;
    display(table)

}

function errorHandler(evt) {
    error("Error", "Can't read this file")
}

function enable(x) {
    table[x][0] = !table[x][0];

    display(table);
}


function display(table) {
    let htmlTable = '<table>';
    for (let i = 0; i < table.length; i++) {
        line = table[i];
        htmlTable = htmlTable + '<tr>';
        for (let e = 0; e < line.length; e++ ) {
            if (e === 0) {
                if (i === 0) {
                    htmlTable = htmlTable + '<th>' + 'add' + '</th>'
                } else {
                    if (!table[i][e]) {
                        htmlTable = htmlTable + '<td><button onclick="enable(' + i + ')">enable</button></td>'

                    } else {
                        htmlTable = htmlTable + '<td><button onclick="enable(' + i + ')">remove</button></td>'
                    }
                }

            } else {
                if (e > 0) {
                    if (i === 0) {
                        htmlTable = htmlTable + '<th>' + line[e] + '</th>'

                    } else {
                        htmlTable = htmlTable + '<td>' + line[e] + '</td>'
                    }
                }
            }
        }
        htmlTable = htmlTable + '</tr>'
    }
    document.getElementById('INaturalistdata').innerHTML = htmlTable;
}


function tabletoPins() {
    console.log(table);
    let data = [];
    var rows = {};

    for (let i = 0; i < table[0].length; i++) {
        rows[table[0][i]] = i
    }

    // console.log(rows);

    for (let i = 1; i < table.length; i++) {
        if (table[i][0]) {
            let pin = {};
            pin["name"] = table[i][rows["common_name"]];
            pin["tag_type"] = "zoo";
            pin["lat"] = parseFloat(table[i][rows["latitude"]]);

            pin["lon"] = parseFloat(table[i][rows["longitude"]]);
            pin["color"] = "#FFFFFF";
            pin["description"] = pinGenerateDiscription(table[i]);
            pin["time"] = 2;
            data.push(pin);
        }
    }
    console.log(data);
    let pins = {};
    pins.data = JSON.stringify(data);
    let n = data.length;
    console.log(JSON.stringify(pins));
    $.post("../api/addPins?num=" + n.toString(16), pins, function (ret) {
        message("Thank you", "upload success");
        console.log("posted");
        console.log(ret);
    }).fail(function () {
        error("Error", "Please and retry.")
    });
}

function pinGenerateDiscription(line) {
    let html = "<p>";
    html = html + "<a href=" + line[header["url"]] + ' target="_blank">look at me on inaturalist.nz</a>';
    html = html + "</p><p>";

    if (line[header["image_url"]] !== "") {
        html = html + "<img src=" + line[header["image_url"]] + ">"
    }

    html = html + "</p><p>";
    html = html + line[header["description"]];
    html = html + "</p><p>";

    httable = "<table>";
    for (let key in header) {
        httable = httable += '<tr><td>' + key + '</td><td>' + line[key] + '</td></tr>'
    }

    httable = httable + "</table>";
    html = html + httable;


    return html
}

// from internet
function CSVToArray(strData, strDelimiter) {
    // Check to see if the delimiter is defined. If not,
    // then default to comma.
    strDelimiter = (strDelimiter || ",");

    // Create a regular expression to parse the CSV values.
    var objPattern = new RegExp(
        (
            // Delimiters.
            "(\\" + strDelimiter + "|\\r?\\n|\\r|^)" +

            // Quoted fields.
            "(?:\"([^\"]*(?:\"\"[^\"]*)*)\"|" +

            // Standard fields.
            "([^\"\\" + strDelimiter + "\\r\\n]*))"
        ),
        "gi"
    );


    // Create an array to hold our data. Give the array
    // a default empty first row.
    var arrData = [[]];

    // Create an array to hold our individual pattern
    // matching groups.
    var arrMatches = null;


    // Keep looping over the regular expression matches
    // until we can no longer find a match.
    while (arrMatches = objPattern.exec(strData)) {

        // Get the delimiter that was found.
        var strMatchedDelimiter = arrMatches[1];

        // Check to see if the given delimiter has a length
        // (is not the start of string) and if it matches
        // field delimiter. If information does not, then we know
        // that this delimiter is a row delimiter.
        if (
            strMatchedDelimiter.length &&
            strMatchedDelimiter !== strDelimiter
        ) {

            // Since we have reached a new row of data,
            // add an empty row to our data array.
            arrData.push([]);

        }

        var strMatchedValue;

        // Now that we have our delimiter out of the way,
        // let's check to see which kind of value we
        // captured (quoted or unquoted).
        if (arrMatches[2]) {

            // We found a quoted value. When we capture
            // this value, unescape any double quotes.
            strMatchedValue = arrMatches[2].replace(
                new RegExp("\"\"", "g"),
                "\""
            );

        } else {

            // We found a non-quoted value.
            strMatchedValue = arrMatches[3];

        }


        // Now that we have our value string, let's add
        // it to the data array.
        arrData[arrData.length - 1].push(strMatchedValue);
    }

    // Return the parsed data.
    return (arrData);
}
