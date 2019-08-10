function loadImages(urls, images) {
    var obj = JSON.parse(urls);
    for (var one in images) {
        images.src = obj[one.id.substring(3)];
    }

}

function requestImageUrl() {
    var images = document.getElementsByClassName("DLImage");
    var nilImage = [];
    var ids = [];
    for (var i in images) {
        if (images.src.length === 0) {
            ids.push(i.id.substring(3));
            nilImage.push(i);
        }
    }
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (request.readyState === 4) {
            if (request.readyState === 200) {
                return loadImages(request.responseText, nilImage)
            }
        }
    };
    request.open("POST", "./api/imageUrls/");
    request.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    request.send(ids.join(";"))
}
