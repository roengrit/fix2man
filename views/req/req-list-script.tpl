<script src="/static/js/datepicker/js/bootstrap-datepicker.min.js"></script>
<script src="/static/js/datepicker/locales/bootstrap-datepicker.th.min.js" charset="UTF-8"></script>
<script >
 $(function () {
    $('#txt-date-begin,#txt-date-end').datepicker({
        autoclose: true,
        language: 'th',
        todayBtn: true,
        orientation: "bottom auto",
        todayHighlight: true  ,
        format: 'dd-mm-yyyy',      
    });
    $('#txt-date-begin').datepicker('setDate', new Date(Date.parse("{{.beginDate}}"))) ;
    $('#txt-date-end').datepicker('setDate', new Date(Date.parse("{{.endDate}}"))) ;
});
    loadNormalTable()
 </script>