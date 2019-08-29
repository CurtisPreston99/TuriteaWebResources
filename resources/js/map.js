Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI1M2YwNTc4Ni0yNWYzLTQ2MTEtOGRkNC05OWFlODNlNTBkZWQiLCJpZCI6MTM5NTksInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQ0NzQwMTl9.X_iNRe8-4jhYrUyAh8QNt3d6aHAfysLye_m0zBHmuiM';
var viewer = new Cesium.Viewer('cesiumContainer', {timeline: false, animation: false});

var pinBuilder = new Cesium.PinBuilder();
viewer.infoBox.frame.removeAttribute('sandbox');
var myObj, i, z = [], url, s = "", a;

$(document).ready(function () {
    $.getJSON("json/pin.json", function (data) {
        $.each(data.pins, function (name, value) {
            s = "<p>Latitude: " + value.lat
                    + "<br>Longitude: " + value.lon
                    + "<br>Description: " + value.description
                    + "<br>Source: wikipedia"
            $("#test").append(value.colour)
            z.push(viewer.entities.add({
                name: value.name,
                position: Cesium.Cartesian3.fromDegrees(value.lat, value.lon),
                description: s,
                point: {
                    pixelSize: 4,
                    color: Cesium.Color.BLACK,
                    outlineColor: Cesium.Color.fromCssColorString(value.colour),
                    outlineWidth: 6
                },
                label: {
                    text: value.label,
                    font: '16pt Arial',
                    fillColor: Cesium.Color.WHITE,
                    style: Cesium.LabelStyle.FILL,
                    verticalOrigin: Cesium.VerticalOrigin.BOTTOM,
                    pixelOffset: new Cesium.Cartesian2(0, -12)
                }/*,
                 billboard: {
                 image : pinBuilder.fromText('X', Cesium.Color.BLACK, 48).toDataURL(),
                 image: 'media/pin2.png',
                 scale: 0.025,
                 
                 }*/
            }));
        });
        Cesium.when.all(z, function (pins) {
            viewer.zoomTo(pins);
        });
    });
});