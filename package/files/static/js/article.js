var home = window.location.origin;

var editor;

function loadPinEditor() {
    editor = $("#summerNotePin");
    editor.summernote({
        height: 500,   //set editable area's height
        codemirror: { // codemirror options
            theme: 'monokai'
        },
        placeholder: "feel free to edit",
    });
    let pin = JSON.parse(localStorage.getItem("editPin"));
    if (pin) {
        let description = pin["description"];
        if (description.startsWith("<")){
            editor.summernote('code', description);
        } else {
            editor.summernote('code', "<p>"+description+"</p>");
        }
        $("#nameAd").val(pin["name"]);
        $("#latTextAd").val(pin["lat"]);
        $("#longTextAd").val(pin["lon"]);
        $("#iconSelectAd").val(pin["tag_type"]);
        $("#colorSelectorsAd").val(pin["color"]);
        console.log($(".note-codable"));
    }

    // console.log(editor);
}

function each(file) {
    let f = file;
    return function(readerEvt) {
        let binaryString = readerEvt.target.result;
        editor.summernote('insertImage', "data:" + f.type + ";base64," + btoa(binaryString));
    };
}
