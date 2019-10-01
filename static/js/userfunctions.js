


var url=window.location.origin;
var selectedUsed="";


function getallUsers(){
  userlist=document.getElementById('userlist');
  $.getJSON(home+"/api/allUser", function (users) {
    console.log(users);
    table="<table><tr><th>removeUser</th><th>userName</th></tr>"
    for (let i=0; i<users.names.length; i++){
      table=table+"<tr>"
      table=table+"<td><button onclick=\'removeuserCheck(\""+users.names[i]+"\")\'>remove</button></td>"
      table=table+"<td>"+users.names[i]+"</td>"
      if(users.roles[i]==1){
// normal
table=table+"<td>normal</td>"

      }
      if(users.roles[i]==2){
// super
table=table+"<td>super</td>"
      }

      table=table+"</tr>"
    }

    userlist.innerHTML=table;
    console.log(table);
  });

}

function removeuserCheck(S) {
  selectedUsed=S;
  popupYN("are you sure you wish to delete User :"+S,removeuserbyName,getallUsers);


}




function removeuserbyName() {
console.log(selectedUsed);
user={}
user.name=selectedUsed;
$.post(url+"/api/deleteUser",user,function(data){
  console.log("posted");
  console.log(data);
   getallUsers()
});

}


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
