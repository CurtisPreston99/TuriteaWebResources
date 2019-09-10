Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI1M2YwNTc4Ni0yNWYzLTQ2MTEtOGRkNC05OWFlODNlNTBkZWQiLCJpZCI6MTM5NTksInNjb3BlcyI6WyJhc3IiLCJnYyJdLCJpYXQiOjE1NjQ0NzQwMTl9.X_iNRe8-4jhYrUyAh8QNt3d6aHAfysLye_m0zBHmuiM';


var home=window.location.origin;
var west =175.590700;
var south =-40.530029;
var east = 175.756629;
var north = -40.385832;
var pinBuilder = new Cesium.PinBuilder();
var viewer;
var entities = [];
var kmlmenu = [];
var description = "";
var image;
var tag_types = ['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo'];
var toolbar = document.getElementById('toolbar');
var scratchRectangle = new Cesium.Rectangle();

function loadMap(){
  viewer = new Cesium.Viewer('cesiumContainer', {
    // imageryProvider : Cesium.createTileMapServiceImageryProvider({
    //   url : Cesium.buildModuleUrl('Assets/Textures/NaturalEarthII')
    // }),
    terrainProvider : new Cesium.CesiumTerrainProvider({
        url: Cesium.IonResource.fromAssetId(1)
    }),
    baseLayerPicker : false,
    geocoder : false,
    timeline	:false,
    animation:false,
    homeButton:false,

  }
);



  viewer.infoBox.frame.removeAttribute('sandbox');

  // Add toolbar to Cesium map & load/hide pins


  Sandcastle.addToolbarButton('Hide Pins', function () {
      viewer.entities.removeAll();
  });

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
    viewer.camera._changed.addEventListener(function(){loadpins()});
    loadpins();
}


function loadpins(){
  var rect = viewer.camera.computeViewRectangle(viewer.scene.globe.ellipsoid, scratchRectangle);
var area = "north=" + Cesium.Math.toDegrees(rect.north).toFixed(8) +
        "&south=" + Cesium.Math.toDegrees(rect.south).toFixed(8) +
        "&east=" + Cesium.Math.toDegrees(rect.east).toFixed(8) +
        "&west=" + Cesium.Math.toDegrees(rect.west).toFixed(8) +
        "&timeBegin=0" +
        "&timeEnd=20000";

console.log(area);
$.getJSON(home+"/api/getPins?"+area, function (data) {
  console.log(data);
    viewer.entities.removeAll();
    $.each(data, function (key, value) {
        description = "<p>Coordinates: (" + value.lon + ", " + value.lat + ")"
                + "<hr>"
                + value.description;
        if (tag_types.includes(value.tag_type)) {
            tag_type = pinBuilder.fromMakiIconId(value.tag_type,Cesium.Color.fromCssColorString(value.color), 48);
        } else {
            tag_type = pinBuilder.fromColor(Cesium.Color.fromCssColorString(value.color), 48);
        }

        // Add entities
        viewer.entities.add({ id: (value.uid).toString(), name:
        value.name,
        position: Cesium.Cartesian3.fromDegrees(value.lon,
        value.lat), description: description, point: { show: false,
        pixelSize: 4, color: Cesium.Color.BLACK, outlineColor:
        Cesium.Color.fromCssColorString(value.color), outlineWidth: 6 },
        label: { show: false, text: (value.uid).toString(), font: '16pt Arial', fillColor: Cesium.Color.WHITE, style:
        Cesium.LabelStyle.FILL, verticalOrigin:
        Cesium.VerticalOrigin.BOTTOM, pixelOffset: new
        Cesium.Cartesian2(0, -12) }, billboard: { image: tag_type,
        verticalOrigin: Cesium.VerticalOrigin.BOTTOM } });
        });
         });
      }
