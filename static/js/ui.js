function popup(I) {
    document.getElementById('popup').innerHTML = I;

    $("#popup").dialog({
        autoOpen: true,
        modal: true,
        width: 500,
        buttons: {
            Done: function () {
                $(this).dialog("close");
            }
        }
    });
}

function popupYN(I, fyes, fno) {
    $("#valid-dialog p").text(I);
    $("#valid-dialog").dialog({
        autoOpen: true,
        modal: true,
        width: 500,
        title: "Right?",
        buttons: {
            Yes: function () {
                $(this).dialog("close");
                fyes();
            },
            No: function () {
                $(this).dialog("close");
                fno();
            }
        }
    });
}
