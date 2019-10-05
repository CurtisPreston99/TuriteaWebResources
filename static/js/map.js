Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI1M2YwNTc4Ni0yNWYzLTQ2MTEtOGRkNC05OWFlODNlNTBkZWQiLCJpZCI6MTM5NTksInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQ0NzQwMTl9.X_iNRe8-4jhYrUyAh8QNt3d6aHAfysLye_m0zBHmuiM';


var home = window.location.origin;
var west = 175.60970056;
var south = -40.38724452;
var east = 175.63276182;
var north = -40.37842748;
var pinBuilder = new Cesium.PinBuilder();
var viewer;
var cesiumHandler;
var kmlmenu = [];
var description = "";
var image;
var ellipsoid;
var tag_types = ['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo'];

var scratchRectangle = new Cesium.Rectangle();

function loadMap() {
    viewer = new Cesium.Viewer('cesiumContainer', {
            // imageryProvider : Cesium.createTileMapServiceImageryProvider({
            //   url : Cesium.buildModuleUrl('Assets/Textures/NaturalEarthII')
            // }),
            terrainProvider: new Cesium.CesiumTerrainProvider({
                url: Cesium.IonResource.fromAssetId(1)
            }),
            baseLayerPicker: false,
            geocoder: false,
            timeline: false,
            animation: false,
            homeButton: false,

        }
    );

    // add
    ellipsoid = viewer.scene.globe.ellipsoid;
    cesiumHandler = new Cesium.ScreenSpaceEventHandler(viewer.canvas);

    viewer.infoBox.frame.removeAttribute('sandbox');

    // Add toolbar to Cesium map & load/hide pins


    // Add KML data to Cesium object
    // kmlmenu.push({
    //     text: "Remove all KML Data",
    //     onselect: function () {
    //         viewer.dataSources.removeAll();
    //     }
    // });
    // $.getJSON("getKML", function (data) {
    //     $.each(data, function (name, value) {
    //         var obj = {};
    //         var kml = new Cesium.KmlDataSource();
    //         kml.load(value.url);
    //         obj.text = value.name;
    //         obj.onselect = function () {
    //             viewer.dataSources.removeAll();
    //             viewer.dataSources.add(kml);
    //             viewer.zoomTo(kml);
    //         };
    //         kmlmenu.push(obj);
    //     });
    //     Sandcastle.addToolbarMenu(kmlmenu, 'toolbar');
    // });
    var rectangle = Cesium.Rectangle.fromDegrees(west, south, east, north);
    viewer.camera.setView({
        destination: rectangle
    });
    viewer.camera._changed.addEventListener(function () {
        loadpins()
    });
    loadpins();
}


function loadpins() {
    var rect = viewer.camera.computeViewRectangle(viewer.scene.globe.ellipsoid, scratchRectangle);
    var area = "north=" + Cesium.Math.toDegrees(rect.north).toFixed(8) +
        "&south=" + Cesium.Math.toDegrees(rect.south).toFixed(8) +
        "&east=" + Cesium.Math.toDegrees(rect.east).toFixed(8) +
        "&west=" + Cesium.Math.toDegrees(rect.west).toFixed(8) +
        "&timeBegin=0" +
        "&timeEnd=20000";

    console.log(area);
    $.getJSON("../api/getPins?" + area, function (data) {
        console.log(data);
        viewer.entities.removeAll();
        if (temPin !== null) {
            viewer.entities.add(temPin);
        }
        $.each(data, function (key, value) {
            description = "<p>Coordinates: (" + value.lon + ", " + value.lat + ")</p>"
                + "<hr>"
                + "<p style='display: none' id='inDescription'>"
                + value.uid.toString(16)
                + "</p>"
                + value.description;
            let tag_type;
            if (tag_types.includes(value.tag_type)) {
                tag_type = pinBuilder.fromMakiIconId(value.tag_type, Cesium.Color.fromCssColorString(value.color), 48);
            } else {
                tag_type = pinBuilder.fromColor(Cesium.Color.fromCssColorString(value.color), 48);
            }

            // Add entities
            viewer.entities.add({
                id: (value.uid).toString(16),
                name: value.name,
                position: Cesium.Cartesian3.fromDegrees(value.lon, value.lat),

                description: description,
                point: {
                    show: false, pixelSize: 4, color: Cesium.Color.BLACK,
                    outlineColor: Cesium.Color.fromCssColorString(value.color), outlineWidth: 6
                },
                label: {
                    show: false, text: (value.uid).toString(), font: '16pt Arial', fillColor: Cesium.Color.WHITE, style:
                    Cesium.LabelStyle.FILL, verticalOrigin:
                    Cesium.VerticalOrigin.BOTTOM, pixelOffset: new
                    Cesium.Cartesian2(0, -12)
                }, billboard: {
                    image: tag_type,
                    verticalOrigin: Cesium.VerticalOrigin.BOTTOM,
                    heightReference: Cesium.HeightReference.CLAMP_TO_GROUND,
                }, pin : value,
            });
        });
    });
}

var temPin = null;
function addTemporaryPin(lon, lat) {
    if (temPin === null) {
        temPin = {
            id: "tem",
            name: "newPin",
            position: Cesium.Cartesian3.fromDegrees(lon, lat),
            point: {
                show: false, pixelSize: 4, color: Cesium.Color.BLACK,
                outlineColor: Cesium.Color.fromCssColorString("#ff0000"), outlineWidth: 6
            },
            label: {
                show: false, text: "tem", font: '16pt Arial', fillColor: Cesium.Color.WHITE,
                style: Cesium.LabelStyle.FILL,
                verticalOrigin: Cesium.VerticalOrigin.BOTTOM,
                pixelOffset: new Cesium.Cartesian2(0, -12)
            }, billboard: {
                image: pinBuilder.fromColor(Cesium.Color.fromCssColorString("#ff0000"), 48),
                verticalOrigin: Cesium.VerticalOrigin.BOTTOM,
                heightReference: Cesium.HeightReference.CLAMP_TO_GROUND,
            },
            pin: {
                "lat" : lat,
                "lon" : lon,
                "time" : 18711,
            },
        }
    } else {
        temPin.position = Cesium.Cartesian3.fromDegrees(lon, lat);
    }
    viewer.entities.removeById("tem");
    console.log(viewer.entities.removed);
    $('#lon').text(lon);
    $('#lat').text(lat);
    viewer.entities.add(temPin);

}

function addPinPrepare() {
    let b = $('#addP');
    b.text("select");
    b.off("click", addPinPrepare);
    b.on("click", selectAndCreate);
    // b.click(selectAndCreate);
    cesiumHandler.setInputAction(selectPosition, Cesium.ScreenSpaceEventType.LEFT_CLICK);
    $("#cancelAdd").removeClass("two_hidden");
}

function cancelAdd() {
    $("#cancelAdd").addClass("two_hidden");
    let b = $('#addP');
    b.text("add pin");
    b.off("click", selectAndCreate);
    b.on("click", addPinPrepare);
    temPin = null;
    viewer.entities.removeById("tem");
    cesiumHandler.removeInputAction(Cesium.ScreenSpaceEventType.LEFT_CLICK);
    $('#lon').text(0.0);
    $('#lat').text(0.0);
}

function selectAndCreate() {
    if (temPin === null) {
        console.log("in");
        let e = $('#error-dialog');
        e.dialog({title: "error"});
        $('#error-dialog p').html("Please select a position!");
        e.dialog('open');
        return;
    }
    let pinDetail = $("#pin-dialog");
    pinDetail.dialog({title:"edit pin"});
    pinDetail.dialog('open');
    $("#advance").show();
    $("#longText").val(temPin.pin["lon"]);
    $("#latText").val(temPin.pin["lat"]);
    // todo 创建一个 dialog 去更新点或者是创建点，并提供一个跳转到扩展模式的按钮
    //temPin = null;
}

function hidePins() {
    viewer.entities.removeAll();
    // console.log(viewer.entities.removed);
    // viewer.entities.
}

function selectPosition(event) {
    var earthPosition  = viewer.camera.pickEllipsoid(event.position, ellipsoid);
    if (Cesium.defined(earthPosition)) {
        let cartographic = ellipsoid.cartesianToCartographic(earthPosition);
        //将弧度转为度的十进制度表示
        let longitude = parseFloat(Cesium.Math.toDegrees(cartographic.longitude));
        let latitude = parseFloat(Cesium.Math.toDegrees(cartographic.latitude));
        addTemporaryPin(longitude, latitude);
        console.log(longitude);
        console.log(latitude);
    }
}

function updatePin() {
    let node = $(".cesium-infoBox-iframe")[0].contentDocument.getElementById("inDescription");
    if (node === null) {
        let e = $('#error-dialog');
        e.dialog({title: "No Select"});
        $('#error-dialog p').html("Please select a pin thank you!");
        e.dialog('open');
    }
    let id = node.innerText;
    let mark = viewer.entities.getById(id);
    let pin = mark.pin;
    let pinDetail = $("#pin-dialog");
    pinDetail.dialog({title:"edit pin"});
    pinDetail.dialog('open');
    $("#advance").hide();
    $("#longText").val(pin["lon"]);
    $("#latText").val(pin["lat"]);
    $("#description").val(pin["description"]);
    $("#name").val(pin["name"]);
    $("#color").val(pin["color"]);
    $("#iconSelect").val(pin["tag_type"]);
    $("#pinId").val(pin["uid"]);
}