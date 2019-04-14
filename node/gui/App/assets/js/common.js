var URL = "http://172.16.162.140:3001"

function delete_intrest(id) {
  console.log(id);
  $('#'+id).remove();
  $.ajax({
    url: URL+'/interest/'+id,
    type: 'DELETE',
    crossDomain: true,
    success: function(result) {
      console.log(result);
    }
  });
}
