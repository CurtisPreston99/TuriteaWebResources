var pins;

//fills the icon selection screen
function addIconOptions(element) {
    Options = ['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo']

    var html = "";

    for (var s in Options) {
        let line = '<option value=' + Options[s] + '>' + Options[s] + '</option>';
        html += line;
    }
    $("#" + element).html(html);
    //document.getElementById(element).innerText = html;
}

// //hides/shows map and text boxes
// function hidemap() {
//
//     if (mapshown) {
//
//         //update the map marker
//         updateMarker(getCords().lon, getCords().lat);
//
//         document.getElementById('mapDiv').style.display = '';
//         console.log("show map");
//         document.getElementById('longLat').style.display = 'none';
//         document.getElementById('inputSwitch').innerText = 'manually enter cords';
//         mapshown = false;
//     } else {
//
//         //update text boxes
//         document.getElementById('longText').value = getCords().lon;
//         document.getElementById('latText').value = getCords().lat;
//         document.getElementById('mapDiv').style.display = 'none';
//
//         console.log("hide map");
//         document.getElementById('longLat').style.display = '';
//         document.getElementById('inputSwitch').innerText = 'use map to enter cords';
//         mapshown = true;
//     }
// }
//
// //returns the selected cords
// function getCords() {
//     let cords = {};
//     if (mapshown) {
//         cords["lat"] = document.getElementById('latText').value;
//         cords["lon"] = document.getElementById('longText').value;
//     } else {
//         if (theMarker) {
//             console.log("marker");
//
//             cords["lat"] = theMarker.getLatLng().lat;
//             cords["lon"] = theMarker.getLatLng().lng;
//         } else {
//             cords["lat"] = 0;
//             cords["lon"] = 0;
//         }
//
//     }
//
//     console.log(cords);
//     return cords;
// }

//compiles pin data one object
function getallData() {
    let pin = {};
    let n = $('#summerNotePin');
    console.log(n);
    if (n.length !== 0) {
        pin["name"] = document.getElementById("nameAd").value;
        pin["tag_type"] = document.getElementById("iconSelectAd").value;
        // console.log(lon);
        pin["lon"] = parseFloat($("#longTextAd")[0].value);
        pin["lat"] = parseFloat($("#latTextAd")[0].value);
        pin["color"] = document.getElementById("colorSelectorsAd").value;
        pin["uid"] = parseInt($("#pinIdAd")[0].value);
        // // fixme 这里应该是将这些作为一个新的article，摘取其中的有效信息作为summary然后description空着
        // pin["description"] = n.summernote('code');
    } else {
        pin["name"] = document.getElementById("name").value;
        pin["tag_type"] = document.getElementById("iconSelect").value;
        // console.log(lon);
        pin["lon"] = parseFloat($("#longText")[0].value);
        pin["lat"] = parseFloat($("#latText")[0].value);
        pin["color"] = document.getElementById("colorSelectors").value;
        pin["uid"] = parseInt($("#pinId")[0].value);
        pin["description"] = $("#description")[0].value;
    }
    pin["time"] = 18711;
    return pin;
}


// refactor by Chen Xingyu, make it can use and make it more easy to use for users
function deletePin() {
    // eval($("#inDescription").innerText);
    let pin = $(".cesium-infoBox-iframe")[0].contentDocument.getElementById("inDescription");
    if (pin === null) {
        let e = $('#error-dialog');
        e.dialog({title: "No Select"});
        $('#error-dialog p').html("Please select a pin thank you!");
        e.dialog('open');
    }
    let id = pin.innerText;

    $.get("../api/delete?type=2&information=" + id, function (result) {
        console.log(result);
        viewer.entities.removeById(id);
    }).fail(function (xhr) {
        let e = $('#error-dialog');
        e.dialog({title: "Not login"});
        $('#error-dialog p').html("Please login thank you!");
        e.dialog('open');
    });
}

function submitPin() {
    if (temPin === null) {
        let pin = getallData();
        let pins = {};
        pins.data = "[" + JSON.stringify(pin) + "]";
        pins.num = 1;
        console.log("the pins" + pins.data);
        $.post("../api/update?type=2", pins, function () {
            console.log("posted");
            let p = $("#pin-dialog");
            if (p.length !== 0) {
                p.dialog('close');
            }
            viewer.entities.removeById("tem");
            temPin = null;
            loadpins();
        }).fail(function (r) {
            let e = $('#error-dialog');
            e.dialog({title: "Not login or other error"});
            $('#error-dialog p').html("Please login thank you or check other things!");
            e.dialog('open');
        });

    } else {
        let pin = getallData();
        let pins = {};
        pins.data = '[' + JSON.stringify(pin) + ']';
        console.log(pins);
        $.post("../api/addPins?num=1", pins, function () {
            console.log("posted");

            loadpins();
            temPin = null;
            viewer.entities.removeById("tem");
            $('#lon').text(0.0);
            $('#lat').text(0.0);
        }).fail(function (r) {
            let e = $('#error-dialog');
            e.dialog({title: "Not login or other error"});
            $('#error-dialog p').html("Please login thank you or check other things!");
            e.dialog('open');
        });
        let p = $("#pin-dialog");
        if (p.length !== 0) {
            p.dialog('close');
        }
    }
}

function toAdvancePins() {
    localStorage.setItem("editPin", JSON.stringify(getallData()));
    if (temPin !== null) {
        localStorage.setItem("mode", "add");
    } else {
        localStorage.setItem("mode", "update")
    }
    window.location.href="../html/settings.html#tabs-2";
}

function submitPinAd() {
    let combination = {};
    let pin = getallData();
    let pinId = null;
    let code = editor.summernote('code');
    // code = code.replace("<img", "<myImage");
    let help = $("<p></p>").append($.parseHTML(code));
    //console.log(help.html());
    let images = help.find("img");
    //console.log(images);
    let datas = [];
    for (let i = 0; i < images.length; i++) {
        let image = images[i];
        let filName = image.dataset.filename;
        let img = $(image);
        let src = img.attr("src");
        if (src.startsWith("data:")) {
            datas.push({"image":src.split(",")[1], "title":filName});
            img.attr("src", "../api/getImage?information=%x");
        }
    }
    //console.log(help.html());
    combination.images = JSON.stringify(datas);
    combination.imageNum = datas.length;
    combination.pins = JSON.stringify([pin,]);
    combination.articles = JSON.stringify([{"sum":"pin name:" + pin.name + " at longitude: " + pin.lon.toString().substring(0, 7) + ", latitude: " + pin.lat.toString().substring(0, 7)},]);
    combination.content = help.html();
    //console.log(help);
    // for (let i = 0; i < help.length; i++) {
    //     combination.content += $(help[i]).html();
    // }
    console.log(combination);
    $.post("../api/addPinWithArticle", combination, function (r) {
        localStorage.setItem("editPin", null);
        // console.log("success");
        window.location.href = "../html/home.html";
    }).fail(function (r) {
        console.log(r);
    })
}
