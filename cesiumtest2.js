var viewer = new Cesium.Viewer('cesiumContainer', {timeline : false, animation : false});

var pinBuilder = new Cesium.PinBuilder();

var url = Cesium.buildModuleUrl('Assets/Textures/maki/cross.png');
var pin1 = Cesium.when(pinBuilder.fromUrl(url, Cesium.Color.GREEN, 48), function(canvas) {
    return viewer.entities.add({
        name : 'Turitea Dam',
        position : Cesium.Cartesian3.fromDegrees(175.66622,-40.46710),
        description: "<h1>Header 1 Text</h1> <h2>Header 2 Text</h2> <h3>Header 3 Text</h3> This is an example of text!",
        billboard : {
            image : pinBuilder.fromMakiIconId('cross', Cesium.Color.GREEN, 48),
            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
        }
    });
});


var url = Cesium.buildModuleUrl('Assets/Textures/maki/cross.png');
var pin2 = Cesium.when(pinBuilder.fromUrl(url, Cesium.Color.GREEN, 48), function(canvas) {
    return viewer.entities.add({
        name : 'Turitea Dam',
        position : Cesium.Cartesian3.fromDegrees(175.676554, -40.432947),
	description : "<img src='Turitea-Dam.jpg' alt='Massey-University' height='105' width='200'>",
        billboard : {
            image : pinBuilder.fromMakiIconId('cross', Cesium.Color.GREEN, 48),
            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
        }
    });
});

var url = Cesium.buildModuleUrl('Assets/Textures/maki/building.png');
var pin3 = Cesium.when(pinBuilder.fromUrl(url, Cesium.Color.GREEN, 48), function(canvas) {
    return viewer.entities.add({
        name : 'Massey University',
        position : Cesium.Cartesian3.fromDegrees(175.618652, -40.386193),
        description: "<img src='Massey-University.png' alt='Massey-University' height='50' width='50'> <a href='https://www.massey.ac.nz/' target='_blank'>Massey University</a>",
        billboard : {
            image : pinBuilder.fromMakiIconId('building', Cesium.Color.RED, 48),
            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
        }
    });
});

var url = Cesium.buildModuleUrl('Assets/Textures/maki/cross.png');
var pin4 = Cesium.when(pinBuilder.fromUrl(url, Cesium.Color.GREEN, 48), function(canvas) {
    return viewer.entities.add({
        name : 'Turitea Stream',
        position : Cesium.Cartesian3.fromDegrees(175.620858, -40.383666),
	description: "<img src='Turitea-Stream.jpg' alt='Turitea-Stream' height='100' width='100'>",
        billboard : {
            image : pinBuilder.fromMakiIconId('cross', Cesium.Color.GREEN, 48),
            verticalOrigin : Cesium.VerticalOrigin.BOTTOM
        }
    });
});


//Since some of the pins are created asynchronously, wait for them all to load before zooming/
Cesium.when.all([pin1, pin2, pin3, pin4], function(pins){
    viewer.zoomTo(pins);
});