function getKMLList() {
    $.getJSON("../api/listKML", function (data) {
        console.log(data);
        let list = $("#kmlList");
        list.empty();
        for (let i = 0; i < data.length; i++) {
            list.append($(`<tr>
                                <td><button onclick="removeKML('{0}')">remove</button></td>
                                <td>{0}</td>
                           </tr>`.format(data[i])));
        }
    }).fail(function () {
        error("Error", "login first or this kml file<br/> has a same name with others in server")
    });
}


function removeKML(name) {
    console.log("removeing: " + name);
    $.get("../api/deleteKML?name=" + name, function () {
        message("success", "the kml has removed");
        getKMLList();
    }).fail(function () {
        error("Error", "Login first");
    });
}

