var table = [];
var header = {};

function handleFiles(files) {
    table = [];
    if (window.FileReader) {
        getAsText(files[0]);
    } else {
        error("Sorry", "try to use a higher version browser")
    }
}

function getAsText(fileToRead) {
    var reader = new FileReader();
    reader.onload = loadHandler;
    reader.onerror = errorHandler;
    reader.readAsText(fileToRead);
}

function loadHandler(event) {
    let sheet = event.target.result;
    table = CSVToArray(sheet, ",");
    var rows = {};

    for (let i = 0; i < table[0].length; i++) {
        rows[table[0][i]] = i
    }
    header = rows;
    let newTable = [];
    let state = false;
    for (let i = 0; i < table.length; i++) {
        if (table[i][header["common_name"]] !== "" && table[i][header["common_name"]] !== null) {
            if (table[i][header["latitude"]] !== "" && table[i][header["latitude"]] !== null) {
                if (table[i][header["longitude"]] !== "" && table[i][header["longitude"]] !== null) {
                    if (i !== 0 && (isNaN(parseFloat(table[i][header["latitude"]])) || isNaN(parseFloat(table[i][header["longitude"]])))) {
                        state = true;
                    } else {
                        newTable.push(table[i]);
                    }
                }
            }
        }
    }
    if (state) {
        error("Error", "Some of the pins has a wrong latitude or longitude<br/>the other pins are shown here")
    }
    table = newTable;
    display(table)

}

function errorHandler(evt) {
    error("Error", "Can't read this file")
}

function enable(x) {
    let b = $("#line" + x.toString());
    table[x][0] = !table[x][0];
    if (table[x][0]) {
        b.text("remove");
    } else {
        b.text("enable");
    }
}


function display(table) {
    let csvHead = $("#csvHead");
    for (let i = 0; i < table[0].length; i++) {
        csvHead.append($("<th>{0}</th>".format(i === 0 ? "add" : table[0][i])));
    }
    let csvBody = $("#csvBody");
    for (let i = 1; i < table.length; i++) {
        let line = table[i];
        let row = $("<tr></tr>");
        for (let e = 0; e < line.length; e++) {
            if (e === 0) {
                if (!table[i][0]) {
                    row.append($(`<td><button id="line{0}" onclick="enable({0})">enable</button></td>`.format(i)));
                } else {
                    row.append($(`<td><button id="line{0}" onclick='enable({0})'>remove</button></td>`.format(i)));
                }
            } else {
                row.append($("<td>{0}</td>".format(line[e])))
            }
        }
        csvBody.append(row);
    }
}


function tabletoPins() {
    let data = [];
    var rows = {};
    for (let i = 0; i < table[0].length; i++) {
        rows[table[0][i]] = i
    }
    let state = false;
    for (let i = 1; i < table.length; i++) {
        if (table[i][0]) {
            let pin = {};
            pin["name"] = table[i][rows["common_name"]];
            pin["tag_type"] = "zoo";
            pin["lat"] = parseFloat(table[i][rows["latitude"]]) + getNumberInNormalDistribution(0, 0.0000005);
            pin["lon"] = parseFloat(table[i][rows["longitude"]]) + getNumberInNormalDistribution(0, 0.0000005);
            pin["color"] = "#FFFFFF";
            pin["description"] = pinGenerateDiscription(table[i]);
            pin["time"] = 2;
            data.push(pin);
        }
    }
    let pins = {};
    pins.data = JSON.stringify(data);
    let n = data.length;
    $.post("../api/addPins?num=" + n.toString(16), pins, function (ret) {
        message("Thank you", "upload success");
    }).fail(function () {
        error("Error", "Please login and retry.")
    });
}

function pinGenerateDiscription(line) {
    let b = $("<p></p>");
    b.append($(`<p><a href="{0}" target="_blank">look at me on inaturalist.nz</a></p>`.format(line[header["url"]])));
    if (line[header["image_url"]] !== "") {
        b.append($(`<img src="{0}" alt="{0}"/>`.format(line[header["image_url"]])));
    }
    b.append($(`<p>{0}</p>`.format(line[header["description"]])));
    let t = $("<table></table>");
    let title = table[0];
    for (let i = 0; i < title.length; i++) {
        let key = title[i];
        let value = line[header[key]];
        if (value) {
            if (value !== "latitude" && value !== "longitude") {
                t.append(`'<tr><td>{0}</td><td>{1}</td></tr>`.format(key, value));
            }
        }
    }
    b.append(t);
    return b.html();
}

// from internet
function CSVToArray(strData, strDelimiter) {
    strDelimiter = (strDelimiter || ",");
    var objPattern = new RegExp(
        (
            "(\\" + strDelimiter + "|\\r?\\n|\\r|^)" +
            "(?:\"([^\"]*(?:\"\"[^\"]*)*)\"|" +
            "([^\"\\" + strDelimiter + "\\r\\n]*))"
        ),
        "gi"
    );
    var arrData = [[]];
    var arrMatches = null;
    while (arrMatches = objPattern.exec(strData)) {
        var strMatchedDelimiter = arrMatches[1];
        if (
            strMatchedDelimiter.length &&
            strMatchedDelimiter !== strDelimiter
        ) {
            arrData.push([]);
        }

        var strMatchedValue;
        if (arrMatches[2]) {
            strMatchedValue = arrMatches[2].replace(
                new RegExp("\"\"", "g"),
                "\""
            );
        } else {
            strMatchedValue = arrMatches[3];
        }
        arrData[arrData.length - 1].push(strMatchedValue);
    }
    return (arrData);
}


// from internet
function getNumberInNormalDistribution(mean, std_dev) {
    return mean + (randomNormalDistribution() * std_dev);
}

// from internet
function randomNormalDistribution() {
    let u = 0.0, v = 0.0, w = 0.0, c = 0.0;
    do {
        u = Math.random() * 2 - 1.0;
        v = Math.random() * 2 - 1.0;
        w = u * u + v * v;
    } while (w === 0.0 || w >= 1.0);
    c = Math.sqrt((-2 * Math.log(w)) / w);
    return u * c;
}