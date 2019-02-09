function submit() {
    console.log("Clicked submit button")
    let uid = localStorage.getItem("uid")
    console.log(uid)
    fetch("http://localhost:8081/api/config")
        .then(response => {
            return response.json()
        }).then(data => {
        console.log(data)
    }).catch(err => console.error(err))
}

