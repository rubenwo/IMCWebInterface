function login() {
    const authURL = "http://localhost/login";

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
        localStorage.setItem("uid", data.uid)
        window.location.href = "configPage.html"
    }).catch(err => console.error(err))
}


