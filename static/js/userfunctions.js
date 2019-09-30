


var url=window.location.origin;

function adduser() {
  user={}
  user.name=document.getElementById("adduserName").value;

  user.role=document.getElementById("roleselect").value;

  console.log(user);

  $.post(url+"/api/addUser",user,function(data){
    console.log("posted");
    console.log(data);
  });
}




function updatePassword() {
  passes={}
  passes.old=calcMD5(document.getElementById("oldPassowrd").value);

  passes.new=calcMD5(document.getElementById("newPassword").value);

  console.log(passes);

  $.post(url+"/api/changePassword",passes,function(data){
    console.log("posted");
    console.log(data);
  });
}

function removeuser() {


  $.post(url+"/api/changePassword",passes,function(data){
    console.log("posted");
    console.log(data);
  });

}
