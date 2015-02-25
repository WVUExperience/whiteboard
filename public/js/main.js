var NotLoggedIn = "You must be logged in to perform this action.";

$(function(){
    $('.votes > i').click(function(){
        var slug = $(this).attr("data-slug");
        console.log("upvoting " + slug + "...");
        $.post('/vote/' + slug, function(data){
            var result = JSON.parse(data);
            if(result.success) {
                var count = $('[data-slug="'+slug+'"]').next();
                $(count).text(Number.parseInt($(count).text())+1);
                $('.votes > i').css({"color":"orangered"});
            } else if(result.error == NotLoggedIn) {
                alert("Please log in to vote. You will now be redirected to login.\nPlease log in using your student MIX account.");
                document.location = "/login?returnUrl=" + encodeURIComponent(document.URL)
            } else {
                alert("Error: " + result.error + "\nIf you are unexpectedly receiving this error, please contact us.");
            }
        });
    });
});