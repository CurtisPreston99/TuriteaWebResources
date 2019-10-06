function popupYN(I, fyes, fno) {
    $("#YN p").text(I);
    $("#YN").dialog({
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
