$(function () {
    $.get( "/prof-name", function( data ) {
        $( "#name-l,#name-r" ).html( data );
      });
});