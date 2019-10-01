



function popup(I){
  document.getElementById('popup').innerHTML=I;

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



function popupYN(I,fyes,fno){
  document.getElementById('popup').innerHTML=I;

  $("#popup").dialog({
      autoOpen: true,
      modal: true,
      width: 500,
      buttons: {
          Yes: function () {
              fyes();
              $(this).dialog("close");
          },
          No: function () {
              fno();
              $(this).dialog("close");
          }
      }
  });
}
