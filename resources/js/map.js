Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI1M2YwNTc4Ni0yNWYzLTQ2MTEtOGRkNC05OWFlODNlNTBkZWQiLCJpZCI6MTM5NTksInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQ0NzQwMTl9.X_iNRe8-4jhYrUyAh8QNt3d6aHAfysLye_m0zBHmuiM';
var viewer = new Cesium.Viewer('cesiumContainer', {timeline : false, animation : false});

var pinBuilder = new Cesium.PinBuilder();

//NOTE: Lance Gray added this
// this script needs to be added to allow scripts to execute in the pop-ups.
viewer.infoBox.frame.removeAttribute('sandbox');
var myObj, i, j=-1, z=[], url, s="";
myObj = {"pins":[
  { "lat":175.66622, "long":-40.46710, "name":"pin0", "label":"0" },
  { "lat":175.676554, "long":-40.432947, "name":"pin1", "label":"1" },
  { "lat":175.618652, "long":-40.386193, "name":"pin2", "label":"2" },
  { "lat":175.620858, "long":-40.383666, "name":"pin3", "label":"3" },
]};

for (i in myObj.pins) {
j ++;
s = "<img src='Turitea-Stream.jpg' alt='Turitea-Stream' height='100' width='100'> <a href='http://localhost/pin-"+j+"'target='_blank'>Wiki Page</a>";

z[j] = viewer.entities.add({
    name : myObj.pins[i].name,
    position : Cesium.Cartesian3.fromDegrees(myObj.pins[i].lat, myObj.pins[i].long),
	description: s,
    point : {
        pixelSize : 5,
        color : Cesium.Color.RED,
        outlineColor : Cesium.Color.WHITE,
        outlineWidth : 2
    },
    label : {
        text : myObj.pins[i].label,
        font : '14pt monospace',
        style: Cesium.LabelStyle.FILL_AND_OUTLINE,
        outlineWidth : 2,
        verticalOrigin : Cesium.VerticalOrigin.BOTTOM,
        pixelOffset : new Cesium.Cartesian2(0, -9)
    }
});
}


//url = Cesium.buildModuleUrl('Assets/Textures/maki/cross.png');
//z[j] = Cesium.when(pinBuilder.fromUrl(url, Cesium.Color.GREEN, 48), function(canvas) {
//    return viewer.entities.add({
//        name : j,
//	position : Cesium.Cartesian3.fromDegrees(myObj.pins[i].lat, myObj.pins[i].long),
//        description: myObj.pins[i].name,
//        billboard : {
//            image : pinBuilder.fromMakiIconId('cross', Cesium.Color.GREEN, 48),
//            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
//        }
//    });
//});
//}

//var wyoming = viewer.entities.add({
//  polygon : {
//    hierarchy : Cesium.Cartesian3.fromDegreesArray([
//                              -109.080842,45.002073,
//                              -105.91517,45.002073,
//                              -104.058488,44.996596,
//                              -104.053011,43.002989,
//                              -104.053011,41.003906,
//                              -105.728954,40.998429,
//                              -107.919731,41.003906,
//                              -109.04798,40.998429,
//                              -111.047063,40.998429,
//                              -111.047063,42.000709,
//                              -111.047063,44.476286,
//                              -111.05254,45.002073]),
//    height : 0,
//    material : Cesium.Color.RED.withAlpha(0.5),
//    outline : true,
//    outlineColor : Cesium.Color.BLACK
//  }
//});

//Since some of the pins are created asynchronously, wait for them all to load before zooming/
Cesium.when.all([z[0], z[1], z[2], z[3]], function(pins){
    viewer.zoomTo(pins);
});