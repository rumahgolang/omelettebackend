<div class="content">
  <div class="form-group">
    <label>Search Place</label>
    <input type="text" class="form-control" id="keyword"
    placeholder="Type place name" name="keyword" value="{{.keyword}}"/>
  </div>
  <button type="submit" class="btn btn-info btn-fill pull-right" onclick="request()">Submit</button>
  <ul id="row"></ul>
</div>
<script>
function request() {
  var keyword = document.getElementById("keyword").value;

  $.ajax({
         url: "https://developers.zomato.com/api/v2.1/search?q=" + keyword,
         type: "GET",
         beforeSend: function(xhr){
           xhr.setRequestHeader('Accept', 'application/json');
           xhr.setRequestHeader('user-key', '8f579904c450d535003969488316607c');
         },
         success: function(response) {
           $("#row").html("");
           console.log(response.restaurants);
           for (i = 0; i < response.restaurants.length; i++) {
             var html =
             "<li>"
             + "<b><span>" + response.restaurants[i].restaurant.name + "</span></b><br/>"
             + "<span>" + response.restaurants[i].restaurant.location.address + "</span>"
             + "<p><a href=\"/merchant/" + response.restaurants[i].restaurant.id + "/add\" class=\"btn btn-info btn-fill\">Add to Omelette</a>"
             + "</li>";
             $("#row").append(html);
           }
         }
      });
}
</script>
