Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI3ODQyYjRlNS05NDg1LTQyM2YtOTJhOS0wODljNjM0MDIxMzIiLCJpZCI6MTM4NTcsInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQxNzgyNDB9.kVn00wi5JwJRS2XyAtYJX12x-jA4EapEscOzw2De16I';
var viewer = new Cesium.Viewer('cesiumContainer', {
    // imageryProvider : Cesium.createTileMapServiceImageryProvider({
    //   url : Cesium.buildModuleUrl('Assets/Textures/NaturalEarthII')
    // }),
    terrainProvider : new Cesium.CesiumTerrainProvider({
        url: Cesium.IonResource.fromAssetId(1)
    }),
    baseLayerPicker : false,
    geocoder : false,
    timeline	:false,
    animation:false
});


var west = 176.0;
var south = -40.8;
var east = -185;
var north = -40.275;

var rectangle = Cesium.Rectangle.fromDegrees(west, south, east, north);
viewer.camera.setView({
    destination: rectangle
});

var pinBuilder = new Cesium.PinBuilder();
var request = new XMLHttpRequest();

function AddPin(pin) {
    console.info(pin);
    Cesium.when(pinBuilder.fromMakiIconId(pin["tag_type"], Cesium.Color.RED, 48), function(canvas) {
        console.log("addPin");
        let x =viewer.entities.add({
            name : 'Hospital',
            position : Cesium.Cartesian3.fromDegrees(pin["lon"],pin["lat"],15.45),
            description:pin["description"],
            billboard : {
                image : canvas.toDataURL(),
                verticalOrigin : Cesium.VerticalOrigin.BOTTOM
            }
        });
        console.log(x)
    });
}

request.onreadystatechange = function() {
    if (request.readyState === 4) {
        if (request.status === 200) {
               var pins = JSON.parse(request.responseText);
               for (var i in pins) {
                   AddPin(pins[i]);
               }
               console.info("get finish")
        }
    }
    // console.info(request)
};

request.open("GET", "./api/pins",true);
console.log("before get");
request.send();

// var hospitalPin = Cesium.when(pinBuilder.fromMakiIconId('hospital', Cesium.Color.RED, 48), function(canvas) {
//     return viewer.entities.add({
//         name: 'Hospital',
//         position: Cesium.Cartesian3.fromDegrees(175.6223059, -40.3876416, 15.45),
//         description: "memes for dreams<h1> big memes</h1>",
//         billboard: {
//             image: canvas.toDataURL(),
//             verticalOrigin: Cesium.VerticalOrigin.BOTTOM
//         }
//     });
// });


//Since some of the pins are created asynchronously, wait for them all to load before zooming/
