let credentials;

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
        credentials = data;
        console.log(credentials)
    }).catch(err => console.error(err))
}


