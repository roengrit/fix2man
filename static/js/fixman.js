function loadNormalTable()
{   
    url = $('#search-form').attr('action');
    $('#txt-search').addClass('load-text');
    $.post(url,$('#search-form').serialize(), function( data ) {
        $('#retCount').html(data.RetCount);
        $('#retData').html(data.RetData);
        $('#txt-search').removeClass('load-text');
    });
}


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
      $('#btn-search-submit').click(function(){
        loadNormalTable()
      })
      $('.change-top').click(function(){
           $('#top').val($(this).attr("top"));
           loadNormalTable()
      })
      $('#search-form').keyup(function(e){
        if (e.keyCode == 13) {
            e.preventDefault();
            loadNormalTable()
            return false;
        }
      })
});