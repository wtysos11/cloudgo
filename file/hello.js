$(document).ready(function() {
    $.ajax({
        url: "/api/123/456"
    }).then(function(data) {
       $('.greeting-id').append(data.id);
       $('.greeting-content').append(data.content);
    });
});
