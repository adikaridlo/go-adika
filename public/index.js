function removeFromDb(item){
    fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
 }
 
//  function updateDb(item) {
//     // let input = document.getElementById(item)
//     // let newitem = input.value
//     fetch(`/update?id=`+Number(item), {method: "GET"}).then(res =>{
//         if (res.status == 200){
//             window.location.pathname = "/update"
//         }
//     })
//  }

function updateDb(item){
    const params  = new URLSearchParams({
        id: item
    })

    const url = `/update?${params.toString()}`
    fetch(url,{ method: 'GET', redirect: 'follow'})
    .then(response => response.json())
    .then(result => {
        // var as = JSON.parse(result);
        // console.log(JSON.parse(result))
        var data = Object.entries(result);
        var res = data[1][1];
        console.log(data[1][1])
        // window.location.pathname = "/update"
    })
}