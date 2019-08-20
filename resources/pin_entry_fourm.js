
var mymap;
var theMarker ;
var mapshown=false;
//runs when window is loaded
window.onload = function(){
  addIconOptions();
  //sets color of pin
  document.getElementById("colorSelectors").value='#ff5555'
  //hodes long lat text boxes
  document.getElementById('longLat').style.display='none';
  // initialize the map on the "map" div with a given center and zoom
 mymap = L.map('map').setView([-40.3994926,175.6390271], 15);


 //adding the tiles to the maps
  L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token=pk.eyJ1Ijoid2l6b3JkIiwiYSI6ImNqeWgwZ3dpOTA2dHczbWxuNXh4NnRkOWsifQ.iw4-VDvNY0e3qtpb8_olGw', {
attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
maxZoom: 18,
id: 'mapbox.streets',
accessToken: 'your.mapbox.access.token'
}).addTo(mymap);
//function that is run when the map is clickeds
mymap.on('click',function(e){

  //get where user clicked
  lat = e.latlng.lat;
  lon = e.latlng.lng;

  //update marker function
  updateMarker(lon,lat);
});
//setting up summer note
$('#summernote').summernote({
height: 400   //set editable area's height

});

}
//to update the marker
function updateMarker(lon,lat){
  //remove marker if it exitis
  if (theMarker) {
        mymap.removeLayer(theMarker);
  };

  //Add a marker to show where you clicked.
  theMarker = L.marker([lat,lon]).addTo(mymap);


   console.log(theMarker);

}

//fills the icon selection screen
function addIconOptions(){
  // TODO:
  // change this array to get from server
  Options=['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo']
  console.log(Options);

  var html=""
  //makes html
  for(var s in Options){
    let line=  '<option value='+Options[s]+'>'+Options[s]+'</option>'
    html+=line;
  }
  console.log(html);
  document.getElementById('iconSelect').innerHTML=html;

}

//hides/shows map and text boxes
function hidemap(){

  if(mapshown){

    //update the map marker
    updateMarker(getCords().lon,getCords().lat);

    document.getElementById('mapDiv').style.display='';
    console.log("show map");
    document.getElementById('longLat').style.display='none';
    document.getElementById('inputSwitch').innerText='manually enter cords';
    mapshown=false;
  }else{

    //update text boxes
    document.getElementById('longText').value=getCords().lon;
    document.getElementById('latText').value=getCords().lat;
    document.getElementById('mapDiv').style.display='none';

    console.log("hide map");
    document.getElementById('longLat').style.display='';
    document.getElementById('inputSwitch').innerText='use map to enter cords';
    mapshown=true;
  }
  }

//returns the selected cords
function getCords(){
  let cords={}
  if(mapshown){
    cords["lat"]= document.getElementById('latText').value;
    cords["lon"]= document.getElementById('longText').value;
}else{
  if(theMarker){
    console.log("marker");

    cords["lat"]=theMarker.getLatLng().lat;
    cords["lon"]=theMarker.getLatLng().lng;
  }

}

console.log(cords);
return cords;
}

//compiles pin data one object
function getallData(){
  let cords =getCords();
  var markupStr = $('#summernote').summernote('code');
  pin={}
  pin["name"]=document.getElementById('name').innerText;
  pin["type"]=document.getElementById("iconSelect").value;
  pin["lat"]=cords["lat"];
  pin["lon"]=cords["lon"];
  pin["color"]=document.getElementById("colorSelectors").value
  pin["description"]=$('#summernote').summernote('code');

  console.log(pin);

  return pin;
}
