function adduser() {
    user = {};
    user.name = document.getElementById("adduserName").value;

    user.role = document.getElementById("roleselect").value;

    console.log(user);

    $.post("../api/addUser", user, function (data) {
        console.log("posted");
        console.log(data);
    });
}


function updatePassword() {
    passes = {};
    passes.old = calcMD5(document.getElementById("oldPassowrd").value);

    passes.new = calcMD5(document.getElementById("newPassword").value);

    console.log(passes);

    $.post("../api/changePassword", passes, function (data) {
        console.log("posted");
        console.log(data);
    });
}

function removeuser() {


    $.post("../api/changePassword", passes, function (data) {
        console.log("posted");
        console.log(data);
    });

}
