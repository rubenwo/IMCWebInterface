function login() {
    const authURL = "http://localhost/auth ";

    let username = document.getElementById("username").value;
    let password = document.getElementById("password").value;
    let data = {
        "username": username,
        "password": password
    };
    console.log(data)
    fetch(authURL, {
        method: 'post',
        body: JSON.stringify(data)
    }).then(response => {
        return response.json()
    }).then(data => {
        console.log(data.uid)
        sessionStorage.setItem("uid", data.uid)
        console.log(data.devices)
        sessionStorage.setItem("devices", JSON.stringify(data.devices))
        window.location.href = "configurator.html"
    }).catch(err => console.error(err))
    return false;
}


