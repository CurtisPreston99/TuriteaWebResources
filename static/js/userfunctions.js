


var url=window.location.origin;

function adduser() {
  user={}
  user.name=document.getElementById("adduserName").value;

  user.role=document.getElementById("roleselect").value;

  console.log(user);

  $.post(url+"/api/addUser",pins,function(data){
    console.log("posted");
    console.log(data);
  });
}
