$(function () {
    $.get( "/prof-name", function( data ) {
        if(data.length>0)
        {
            $( "#name-l" ).html( data.substring(0,20) );    
            $( "#name-r" ).html( data  );           
        }else{
             $( "#name-l,#name-r" ).html( data );
        }
      });
});