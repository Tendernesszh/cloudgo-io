$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.id').append(data.id);
       $('.password').append(data.password);
    });
});
