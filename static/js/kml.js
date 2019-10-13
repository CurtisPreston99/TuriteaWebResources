function getKMLList() {
    $.getJSON("../api/listKML", function (data) {
        let list = $("#kmlList");
        list.empty();
        for (let i = 0; i < data.length; i++) {
            list.append($(`<tr>
                                <td><button onclick="removeKML('{0}')">remove</button></td>
                                <td>{0}</td>
                           </tr>`.format(data[i])));
        }
    }).fail(function () {
        error("Error Message", "Sorry, something went wrong. Please try again.")
    });
}


function removeKML(name) {
    $.get("../api/deleteKML?name=" + name, function () {
        message("success", "the kml has removed");
        getKMLList();
    }).fail(function () {
        error("Error Message", "Sorry, something went wrong. Please try again.");
    });
}

