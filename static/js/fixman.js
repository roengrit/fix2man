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

function editNormal(id) {
    hideTopAlert();
    $.get("/normal/add/?entity="+ $("#entity").val() + "&id=" + id , function (data) {
        if (data.RetOK) {
            showGlobalSmallModal();
            $('#small-global-modal-content').html(data.RetData);
            
        } else {
            showGlobalSmallModal();
            $('#small-global-modal-content').html(data.RetData);
        }
    });
} 

function deleteNormal(id) {
    hideTopAlert();
    $.get("/normal/add/?entity="+ $("#entity").val() + "&del=1&id=" + id , function (data) {
        if (data.RetOK) {
            showGlobalSmallModal();
            $('#small-global-modal-content').html(data.RetData);
            
        } else {
            showGlobalSmallModal();
            $('#small-global-modal-content').html(data.RetData);
        }
    });
} 

function hideGlobalSmalModal()
{
    $('#small-global-modal').modal("hide");
}
function showGlobalSmallModal()
{
    $('#small-global-modal').modal("show");
}

function showTopAlert(alert,type)
{
    var html = `<div class="alert alert-`+type+` alert-dismissible"  >
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    `+ alert + `
    </div>`;
    $("#top-alert").html(html)
    $("#top-alert").fadeIn(500, function () {

    });
}

function hideTopAlert()
{
    $("#top-alert").fadeOut(500, function () {

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
           $('#top-search-label').html($(this).html());
           loadNormalTable()
      })
      $('#search-form').keyup(function(e){
        if (e.keyCode == 13) {
            e.preventDefault();
            loadNormalTable()
            return false;
        }
      })
      $('#normal-add').click(function(){
        hideTopAlert();
         $.get( "/normal/add/?entity="+ $(this).attr("entity") , function( data ) {
            $('#small-global-modal-content').html(data.RetData);
            showGlobalSmallModal();
          });
      })
     
});