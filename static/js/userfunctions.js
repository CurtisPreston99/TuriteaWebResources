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
    }).fail(function (r) {
        console.log(r.status);
        if (r.status === 500) {
            error("Data Validation Error", "The name is used by other account");
        } else if (r.status === 403) {
            error("Login Error", "please login first");
        }
    });
}


function updatePassword() {
    passes = {};
    passes.old = calcMD5(document.getElementById("oldPassowrd").value);

    passes.new = calcMD5(document.getElementById("newPassword").value);

    // console.log(passes);

    $.post("../api/changePassword", passes, function (data) {
        message("Change Successful", "change password success");
    }).fail(function () {
        error("Not login or other error", "Please login thank you or check other things!");
    });
}

function loadUsers() {
    $.get("../api/allUser", function (r) {
        let list = $("#userList");
        console.log(list);
        list.empty();
        let users = JSON.parse(r);
        let names = users["names"];
        let roles = users["roles"];
        let length = names.length;
        for (let i = 0; i < length; i++) {
            let row = `<tr class="user">
                <td><input type="checkbox" name="selectUsers" value="{0}"></td>
                <td>{0}</td>
                <td><button id="{0}" onclick="changeRole('{0}',{2})" >{1}</button></td>
            </tr>`.format(names[i], roles[i] === 1 ? "Researcher" : (roles[i] === 2 ? "Administer" : ""), roles[i]);
            list.append($(row));
        }
    }).fail(function () {
        error("Login Error", "Please login and retry");
    });
}

async function removeUsers(names) {
    // message("sorry", "processing...");
    for (let i = 0; i < names.length; i++) {
        $.get("../api/deleteUser?name=" + names[i], ).fail(
            await function () {
                i = names.length;
                error("Not login or other error", "Please login thank you or check other things!");
            }
        );
    }
    await loadUsers();
    // $("#valid-dialog").dialog("close");
}

function deleteUsers() {
    let selectUsers = $("input[name='selectUsers']:checked");
    let names = [];
    for (let i = 0; i < selectUsers.length; i++) {
        names.push(selectUsers[i].value);
    }
    removeUsers(names)
}

function changeRole(name, role) {
    if (role === 1) {
        $.post("../api/changeRole", {"name":name, "newRole":2},function () {
            loadUsers()
        }).fail(function () {
            error("Not login", "Please login first")
        });
    } else if (role === 2) {
        $.post("../api/changeRole", {"name":name, "newRole":1},function () {
            loadUsers()
        }).fail(function () {
            error("Not login", "Please login first")
        });
    }
}

