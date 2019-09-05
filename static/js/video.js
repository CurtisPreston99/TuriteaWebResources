var i = 0;
$.getJSON("getVideo", function (data) {
    $.each(data, function (name, value) {
        if (i % 4 === 0) {
            $("table").append("<tr><td></td></tr>");
        } else {
            $("tr:last").append("<td></td>");
        }
        $("td:last").append("<iframe src='" + value.url + "' allow='accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture' allowfullscreen></iframe>");
        $("td:last").append("<h4>" + value.title + "</h4>");
        $("td:last").append("<h5>" + value.author + "</h5>");
        $("td:last").append("<h5>" + value.date + "</h5>");
        $("td:last").append("<h6>" + value.description + "</h6>");
        i++;
    });
});
