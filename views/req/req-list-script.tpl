<script src="/static/js/datepicker/js/bootstrap-datepicker.min.js"></script>
<script src="/static/js/datepicker/locales/bootstrap-datepicker.th.min.js" charset="UTF-8"></script>
<script >
 $(function () {
    $('.date').datepicker({
        autoclose: true,
        language: 'th',
        todayBtn: true,
        orientation: "bottom auto",
        todayHighlight: true  ,
        format: 'dd-mm-yyyy',      
    });
    $('.date').datepicker('setDate', new Date(Date.parse("{{.currentDate}}"))) ;
});
    loadNormalTable()
 </script>