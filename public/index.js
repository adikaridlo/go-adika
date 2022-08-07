function removeFromDb(item){
    fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
 }
//  1
//  function updateDb(item) {
//     // let input = document.getElementById(item)
//     // let newitem = input.value
//     fetch(`/update?id=${item}`, {method: "PUT"}).then(res =>{
//         if (res.status == 200){
//             const json = JSON.stringify(res);
//             $("#form-id").html(res.text);
//             console.log(res)
//             console.log(json)
//             // window.location.pathname = "/update"
//         }
//     })
//  }

//  2
// function updateDb(item){
//     const params  = new URLSearchParams({
//         id: item
//     })

//     const url = `/update?${params.toString()}`
//     fetch(url,{ method: 'GET', redirect: 'follow'})
//     .then(response => response.json())
//     .then(result => {
//         // var as = JSON.parse(result);
//         // console.log(JSON.parse(result))
//         var data = Object.entries(result);
//         var res = data[1][1];
//         console.log(data[1][1])
//         // window.location.pathname = "/update"
//     })
// }

// 3
function updateDb(item) {
    // let input = document.getElementById(item)
    // let newitem = input.value
    fetch(`/update?id=${item}`, {method: "PUT"}).then(function (response) {
        // The API call was successful!
        return response.text();
    }).then(function (html) {
        // This is the HTML from our response as a text string
        console.log(html);
        $("#form-id").html(html)
    }).catch(function (err) {
        // There was an error
        console.warn('Something went wrong.', err);
    })
 }

 $(document).ready(function () {
    $('#table-task').dataTable( {
        "autoWidth": false,
        // searching: false,
        // paging: false,
        // info: false,
        "autoWidth": false,
        "columns": [
            { "width": "20%" },
            { "width": "100%!important" },
            { "width": "100%!important" },
            { "width": "100%!important" },
            { "width": "100%!important" },
            { "width": "100%!important" },
            ]
    });
    $('.date').datepicker({
        autoclose: true
    });
});