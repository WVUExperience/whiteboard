$(function(){
    $('.votes > i').click(function(){
        var slug = $(this).attr("data-slug");
        console.log("upvoting " + slug + "...");
        $.post('/vote/' + slug, function(data){
            var result = JSON.parse(data);
            if(result.success) {
                $(this).css({"color":"orangered"});
            } else {
                alert("Error: " + result.error);
            }
        });
    });
});