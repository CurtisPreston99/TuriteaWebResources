


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


      var west = 176.0;
      var south = -40.8;
      var east = -185;
      var north = -40.275;

      var rectangle = Cesium.Rectangle.fromDegrees(west, south, east, north);
      viewer.camera.setView({
          destination: rectangle
      });

      var pinBuilder = new Cesium.PinBuilder();

    //   var hospitalPin = Cesium.when(pinBuilder.fromMakiIconId('hospital', Cesium.Color.RED, 48), function(canvas) {
    // viewer.entities.add({
    //     name : 'Hospital',
    //     position : Cesium.Cartesian3.fromDegrees(175.6223059,-40.3876416,15.45),
    //     description:"memes for dreams<h1> big memes</h1>",
    //     billboard : {
    //         image : canvas.toDataURL(),
    //         verticalOrigin : Cesium.VerticalOrigin.BOTTOM
    //     }
    // });
// });



$.get('./pins', {}, function(data){
   pins=JSON.parse(data);


   console.log(pins);

   for(let i=0;i<pins.length;i++){
      Cesium.when(pinBuilder.fromMakiIconId('hospital', Cesium.Color.RED, 48), function(canvas) {
   let x =viewer.entities.add({
         name : 'Hospital',
         position : Cesium.Cartesian3.fromDegrees(pins[i]["lon"],pins[i]["lat"],15.45),
         description:pins[i]["description"],
         billboard : {
             image : canvas.toDataURL(),
             verticalOrigin : Cesium.VerticalOrigin.BOTTOM
         }
     });
     console.log(x);
   })
   }

   console.log(data);
 });
