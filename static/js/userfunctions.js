function adduser() {
    user = {};
    user.name = document.getElementById("adduserName").value;

    user.role = document.getElementById("roleselect").value;

    console.log(user);

    $.post("../api/addUser", user, function (d) {
        console.log("posted");
        let data = JSON.parse(d);
        console.log(data);
        $("#userResult").removeClass("two_hidden");
        $("#userName").text(data["name"]);
        $("#password").text(data["password"]);
        if (data["role"] === 1) {
            $("#role").text("Researcher");
        } else if (data["role"] === 2) {
            $("#role").text("Administer");
        }
    }).fail(function () {

    });
}


function updatePassword() {
    passes = {};
    passes.old = calcMD5(document.getElementById("oldPassowrd").value);

    passes.new = calcMD5(document.getElementById("newPassword").value);

    // console.log(passes);

    $.post("../api/changePassword", passes, function (data) {
        // console.log("posted");
        // console.log(data);
        let v = $('#valid-dialog');
        v.dialog({title: "Change Successful"});
        $('#valid-dialog p').html("change password success");
        v.dialog('open');
    });
}

function removeuser(name) {

    $.get("../api/deleteUser?name=" + name);

}
