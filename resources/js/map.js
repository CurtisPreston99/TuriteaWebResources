Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI1M2YwNTc4Ni0yNWYzLTQ2MTEtOGRkNC05OWFlODNlNTBkZWQiLCJpZCI6MTM5NTksInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQ0NzQwMTl9.X_iNRe8-4jhYrUyAh8QNt3d6aHAfysLye_m0zBHmuiM';
var viewer = new Cesium.Viewer('cesiumContainer');
var pinBuilder = new Cesium.PinBuilder();
var entities = [];
var menuitems = [];
var description = "";
var image;
var tag_types = ['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo'];
viewer.infoBox.frame.removeAttribute('sandbox');
// Add entities to Cesium object
Sandcastle.addToolbarButton('Load Pins', function () {
    $('#map-dialog').dialog('open');
    ;
});
Sandcastle.addToolbarButton('Show Pins', function () {
    viewer.entities.removeAll();
    $.getJSON("json/pin.json", function (data) {
        $.each(data, function (key, value) {
            description = "<p>Coordinates: (" + value.lon + ", " + value.lat + ")"
                    + "<hr>"
                    + value.description;
            if (tag_types.includes(value.tag_type)) {
                tag_type = pinBuilder.fromMakiIconId(value.tag_type, Cesium.Color.fromCssColorString(value.colour), 48);
            } else {
                tag_type = pinBuilder.fromColor(Cesium.Color.fromCssColorString(value.colour), 48);
            }
            viewer.entities.add({
                id: (value.uid).toString(),
                name: value.name,
                position: Cesium.Cartesian3.fromDegrees(value.lat, value.lon),
                description: description,
                point: {
                    show: false,
                    pixelSize: 4,
                    color: Cesium.Color.BLACK,
                    outlineColor: Cesium.Color.fromCssColorString(value.colour),
                    outlineWidth: 6
                },
                label: {
                    show: false,
                    text: (value.uid).toString(),
                    font: '16pt Arial',
                    fillColor: Cesium.Color.WHITE,
                    style: Cesium.LabelStyle.FILL,
                    verticalOrigin: Cesium.VerticalOrigin.BOTTOM,
                    pixelOffset: new Cesium.Cartesian2(0, -12)
                },
                billboard: {
                    image: tag_type,
                    verticalOrigin: Cesium.VerticalOrigin.BOTTOM
                }
            });
        });
    });
    viewer.entities.add({
        name: 'Red Box Demonstration',
        position: Cesium.Cartesian3.fromDegrees(175.7, -40.4, 0),
        box: {
            dimensions: new Cesium.Cartesian3(30000.0, 30000.0, 0),
            material: Cesium.Color.RED.withAlpha(0.1)
        }
    });
    viewer.zoomTo(viewer.entities);
});
Sandcastle.addToolbarButton('Hide Pins', function () {
    viewer.entities.removeAll();
});
// Add KML data to Cesium object
menuitems.push({
    text: "Remove all KML Data",
    onselect: function () {
        viewer.dataSources.removeAll();
    }
});
$.getJSON("json/kml.json", function (data) {
    $.each(data, function (name, value) {
        var obj = {};
        var kml = new Cesium.KmlDataSource();
        kml.load(value.url);
        obj.text = value.name;
        obj.onselect = function () {
            viewer.dataSources.removeAll();
            viewer.dataSources.add(kml);
            viewer.zoomTo(kml);
        };
        menuitems.push(obj);
    });
    Sandcastle.addToolbarMenu(menuitems, 'toolbar');
});