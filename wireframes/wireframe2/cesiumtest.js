

// NOTE: Viewer constructed after default view is set.
//loading and making map
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
    //change view to be centered on the Turitea

      var west = 176.0;
      var south = -40.8;
      var east = -185;
      var north = -40.275;

      var rectangle = Cesium.Rectangle.fromDegrees(west, south, east, north);
      viewer.camera.setView({
          destination: rectangle
      });

      //how to add pins
      var pinBuilder = new Cesium.PinBuilder();

      var hospitalPin = Cesium.when(pinBuilder.fromMakiIconId('hospital', Cesium.Color.RED, 48), function(canvas) {
    return viewer.entities.add({
        name : 'Hospital',
        position : Cesium.Cartesian3.fromDegrees(175.6223059,-40.3876416,getHight(175.6223059,-40.3876416)),
        description:"memes for dreams<h1> big memes</h1>",
        billboard : {
            image : canvas.toDataURL(),
            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
        }
    });
});

viewer.dataSources.add(Cesium.KmlDataSource.load('./Turitea Pitfall Points.kmz',
     {
          camera: viewer.scene.camera,
          canvas: viewer.scene.canvas
     })
);



function getHight(long,lat){

  // Query the terrain height of two Cartographic positions
var terrainProvider = viewer.terrainProvider;
var positions = [
    Cesium.Cartographic.fromDegrees(long, lat),
];
var promise = Cesium.sampleTerrain(terrainProvider, 11, positions);
Cesium.when(promise, function(updatedPositions) {
    console.log(updatedPositions[0].height);
    return updatedPositions[0].height;
});
}

viewer.fullscreenButton.viewModel.command.afterExecute.addEventListener(function() {
  const canvas = viewer.canvas;
  if ('webkitRequestFullscreen' in canvas) {
    // chrome
    canvas.webkitRequestFullscreen();
  } else if ('requestFullScreen' in canvas) {
      // other browsers
      canvas.requestFullScreen();
  }
});
