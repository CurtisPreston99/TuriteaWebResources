
var mymap;
var theMarker ;


window.onload = function(){
  addIconOptions();
  document.getElementById("colorSelectors").value='#ff5555'
  //long lat text boxes
  document.getElementById('longLat').style.display='none';
  // initialize the map on the "map" div with a given center and zoom
 mymap = L.map('map').setView([-40.3994926,175.6390271], 13);



  L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token=pk.eyJ1Ijoid2l6b3JkIiwiYSI6ImNqeWgwZ3dpOTA2dHczbWxuNXh4NnRkOWsifQ.iw4-VDvNY0e3qtpb8_olGw', {
attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
maxZoom: 18,
id: 'mapbox.streets',
accessToken: 'your.mapbox.access.token'
}).addTo(mymap);

mymap.on('click',function(e){
  lat = e.latlng.lat;
  lon = e.latlng.lng;

      //Clear existing marker,

      if (theMarker) {
            mymap.removeLayer(theMarker);
      };

  //Add a marker to show where you clicked.
   theMarker = L.marker([lat,lon]).addTo(mymap);


       console.log(theMarker);
});


}


function addIconOptions(){
  Options=['airfield', 'airport', 'alcohol-shop', 'america-football', 'art-gallery', 'bakery', 'bank', 'bar', 'baseball', 'basketball', 'beer', 'bicycle', 'building', 'bus', 'cafe', 'camera', 'campsite', 'car', 'cemetery', 'cesium', 'chemist', 'cinema', 'circle', 'circle-stroked', 'city', 'clothing-store', 'college', 'commercial', 'cricket', 'cross', 'dam', 'danger', 'disability', 'dog-park', 'embassy', 'emergency-telephone', 'entrance', 'farm', 'fast-food', 'ferry', 'fire-station', 'fuel', 'garden', 'gift', 'golf', 'grocery', 'hairdresser', 'harbor', 'heart', 'heliport', 'hospital', 'ice-cream', 'industrial', 'land-use', 'laundry', 'library', 'lighthouse', 'lodging', 'logging', 'london-underground', 'marker', 'marker-stroked', 'minefield', 'mobilephone', 'monument', 'museum', 'music', 'oil-well', 'park2', 'parking-garage', 'parking', 'park', 'pharmacy', 'pitch', 'place-of-worship', 'playground', 'police', 'polling-place', 'post', 'prison', 'rail-above', 'rail-light', 'rail-metro', 'rail', 'rail-underground', 'religious-christian', 'religious-jewish', 'religious-muslim', 'restaurant', 'roadblock', 'rocket', 'school', 'scooter', 'shop', 'skiing', 'slaughterhouse', 'soccer', 'square', 'square-stroked', 'star', 'star-stroked', 'suitcase', 'swimming', 'telephone', 'tennis', 'theatre', 'toilets', 'town-hall', 'town', 'triangle', 'triangle-stroked', 'village', 'warehouse', 'waste-basket', 'water', 'wetland', 'zoo']
  console.log(Options);

  var html=""

  for(var s in Options){
    let line=  '<option value='+Options[s]+'>'+Options[s]+'</option>'
    html+=line;
  }
  console.log(html);
  document.getElementById('iconSelect').innerHTML=html;

}


function hidemap(){

  if(document.getElementById('map').style.display=='none'){
    document.getElementById('map').style.display=''
    console.log("show map");
    document.getElementById('longLat').style.display='none';
    document.getElementById('inputSwitch').innerText='manually enter cords';
  }else{
    document.getElementById('map').style.display='none';
    document.getElementById('long').innerText=theMarker.getLatLng().lng;
    document.getElementById('lat').innerText=theMarker.getLatLng().lat;
    console.log("hide map");
    document.getElementById('longLat').style.display='';
    document.getElementById('inputSwitch').innerText='use map to enter cords';
  }
  }


function getCords(){
  let cords={}
  if(document.getElementById('map').style.display=='none'){
    cords["lat"]=document.getElementById('lat').innerText;
    cords["lon"]=document.getElementById('long').innerText;
}else{
  if(theMarker){
    cords["lat"]=theMarker.getLatLng().lat;
    cords["lon"]=theMarker.getLatLng().lng;
}else{
cords["lat"]=0;
cords["lon"]=0;

}

}
return cords;
}

function getallData(){
  let cords =getCords();
  console.log(cords);
  var markupStr = $('#summernote').summernote('code');
  pin={}
  pin["name"]=document.getElementById('name').innerText;
  pin["type"]=document.getElementById("iconSelect").value;
  pin["lat"]=cords["lat"];
  pin["lon"]=cords["lon"];
  pin["color"]=document.getElementById("colorSelectors").value
  pin["description"]=$('#summernote').summernote('code');

  console.log(pin);
}
