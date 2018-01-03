<script src="/static/js/datepicker/js/bootstrap-datepicker.min.js"></script>
<script src="/static/js/datepicker/locales/bootstrap-datepicker.th.min.js" charset="UTF-8"></script>
<script src="/static/js/bootstrap-typeahead.js" charset="UTF-8"></script>
<script >
 $(function () {
     {{if .r}}

    {{else}}
    $('#req-date-event').datepicker({
        autoclose: true,
        language: 'th',
        todayBtn: true,
        orientation: "bottom auto",
        todayHighlight: true  ,
        format: 'dd-mm-yyyy',      
    });
    $('#req-date-event').datepicker('setDate', new Date(Date.parse("{{.currentDate}}"))) ;
    {{end}}

     $('#req-name').typeahead({
        ajax: '/service/userlist/json/?entity=users',
        display: 'Name',
        displayField: 'Name',
        valueField: 'ID',
        val: 'ID',
        onSelect:function(data){
            $('#req-name').addClass('check-text');
            $('#req-name-id').val(data.value); 
            if(data.value == 0){
                return;
            }       
            if($('#req-doc-id').val()!=''){
                return;
            }

            $.get( "/service/user/json/?query="+ data.value , function( userData ) {
                    
                    $('#req-tel').val(userData.Tel);
                    $('#req-tel').addClass('check-text');

                    $('#req-branch').val(userData.Branch.Name);
                    $('#req-branch').addClass('check-text');
                    $('#req-branch-id').val(userData.Branch.ID);  

                    $('#req-depart').val(userData.Depart.Name);        
                    $('#req-depart').addClass('check-text');
                    $('#req-depart-id').val(userData.Depart.ID);  

                    $('#req-building').val(userData.Building.Name);  
                    $('#req-building').addClass('check-text');      
                    $('#req-building-id').val(userData.Building.ID);      

                    $('#req-class').val(userData.Class.Name);     
                    $('#req-class').addClass('check-text');         
                    $('#req-class-id').val(userData.Class.ID);   
                          
                    $('#req-room').val(userData.Rooms.Name);        
                    $('#req-room').addClass('check-text');         
                    $('#req-room-id').val(userData.Rooms.ID);
                    
            });    
        }
    });

    $('#req-branch').typeahead({
        ajax: '/service/entitylist/json/?entity=branchs',
        display: 'Name',
        displayField: 'Name',
        valueField: 'ID',
        val: 'ID',
        onSelect:function(data){

            $('#req-branch').addClass('check-text');
            $('#req-branch-id').val(data.value);

            $('#req-building').val('');
            $('#req-building-id').val('');
            $('#req-building').removeClass('check-text');              

            $('#req-class').val('');
            $('#req-class-id').val('');
            $('#req-class').removeClass('check-text');

            $('#req-room').val('');
            $('#req-room-id').val('');
            $('#req-room').removeClass('check-text');  
        }
    });

    $('#req-building').typeahead({
                ajax: '/service/entitylist-p/json/?entity=buildings'  ,
                display: 'Name',
                displayField: 'Name',
                valueField: 'ID',
                val: 'ID',
                parent : $('#req-branch-id'), 
                onSelect:function(data){

                    $('#req-building').addClass('check-text');
                    $('#req-building-id').val(data.value); 

                    $('#req-class').val('');
                    $('#req-class-id').val('');
                    $('#req-class').removeClass('check-text');

                    $('#req-room').val('');
                    $('#req-room-id').val('');
                    $('#req-room').removeClass('check-text');

                },
                parent : 'req-branch-id',
                fixurl : '/service/entitylist-p/json/?entity=buildings'
            });
    $('#req-class').typeahead({
            ajax: '/service/entitylist-p/json/?entity=class'   ,
            display: 'Name',
            displayField: 'Name',
            valueField: 'ID',
            val: 'ID',
            onSelect:function(data){
                
                $('#req-class').addClass('check-text');
                $('#req-class-id').val(data.value);  
                
                $('#req-room').val('');
                $('#req-room-id').val('');
                $('#req-room').removeClass('check-text');
            },
            parent : 'req-building-id',
            fixurl : '/service/entitylist-p/json/?entity=class'
        });

    $('#req-room').typeahead({
            ajax: '/service/entitylist-p/json/?entity=rooms'   ,
            display: 'Name',
            displayField: 'Name',
            valueField: 'ID',
            val: 'ID',
            onSelect:function(data){
                $('#req-room').addClass('check-text');
                $('#req-room-id').val(data.value);           
            },
            parent : 'req-class-id',
            fixurl : '/service/entitylist-p/json/?entity=rooms'
        });

    $('#req-depart').typeahead({
                ajax: '/service/entitylist/json/?entity=departs'   ,
                display: 'Name',
                displayField: 'Name',
                valueField: 'ID',
                val: 'ID',
                onSelect:function(data){
                    $('#req-depart').addClass('check-text');
                    $('#req-depart-id').val(data.value); 
                }
            });

    $('#req-sn').typeahead({
        ajax: '/service/entitylist/json/?entity=sns',
        display: 'Name',
        displayField: 'Name',
        valueField: 'ID',
        val: 'ID',
        onSelect:function(data){
            $('#req-sn').addClass('check-text');
            $('#req-sn-id').val(data.value);            
        }
    });
   $('#req-tel').keyup(function(e){
       if($(this).val() == ""){
           $('#req-tel').removeClass('check-text');
       } else{
           $('#req-tel').addClass('check-text');
       }   
   });

    $('.has-query').keyup(function(e){
         switch (e.keyCode) {
                case 40: // down arrow
                case 38: // up arrow
                case 16: // shift
                case 17: // ctrl
                case 18: // alt
                    break
                case 9: // tab
                case 13: // enter
                    break
                case 27: // escape
                    break
                default:  { 
                    $(this).removeClass("check-text"); 
                    var name = $(this).attr("name"); 
                    $("#" + name+ "-id").val(''); 
                }  break
            }
    });
  
 });
 function Clear()
 {
                    $('#req-doc-id').val(''); 

                    $('#req-name').val(''); 
                    $('#req-name').removeClass('check-text');
                    $('#req-name-id').val(''); 

                    $('#req-tel').val('')
                    $('#req-tel').val('');
                    $('#req-tel').removeClass('check-text');

                    $('#req-branch').val('');
                    $('#req-branch').removeClass('check-text');
                    $('#req-branch-id').val('');  

                    $('#req-depart').val('');        
                    $('#req-depart').removeClass('check-text');
                    $('#req-depart-id').val('');  

                    $('#req-building').val('');  
                    $('#req-building').removeClass('check-text');      
                    $('#req-building-id').val('');      

                    $('#req-class').val('');     
                    $('#req-class').removeClass('check-text');         
                    $('#req-class-id').val('');   
                          
                    $('#req-room').val('');        
                    $('#req-room').removeClass('check-text');         
                    $('#req-room-id').val('');
                                              
                    $('#req-sn').val('');        
                    $('#req-sn').removeClass('check-text');         
                                              
                    $('#remark').val('');        
 }
 function Save(){
          hideTopAlert();
            url = $('#req-form').attr('action');
            $.post(url,$('#req-form').serialize(), function( data ) {
                $('#normal-code').removeClass('load-text');
                if(data.RetOK){
                    showTopAlert(data.RetData,"success")
                    setTimeout(function(){ window.location.href = '/request/list/' }, 2000);
                }else{
                    showTopAlert(data.RetData,"danger")
                }
            });
 }
</script>