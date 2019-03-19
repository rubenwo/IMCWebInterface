let state = {
    ALARM: 1,
    RADIO_STATUS: 2,
    NETWORK: 3,
    LANGUAGE_NAME: 4,
    HELP: 5
};

let currentState;

let device_ids = JSON.parse(sessionStorage.getItem("devices"))
console.log(device_ids)

for (let i = 0; i < device_ids.lenght; i++) {
    console.log(device_ids[i])
}



function alarm() {
    console.log("Clicked Alarm button")
    let html = "ALARM";
    document.getElementById("main_content").innerHTML = html;
    currentState = device_ids//state.ALARM
}

function radio_status() {
    console.log("Clicked Radio Status button")
    document.getElementById("main_content").innerHTML = "RADIO STATUS"
    currentState = state.RADIO_STATUS
}

function network() {
    console.log("Clicked Network button")
    document.getElementById("main_content").innerHTML = "NETWORK"
    currentState = state.NETWORK
}

function languageNtime() {
    console.log("Clicked Language & Time button")
    document.getElementById("main_content").innerHTML = "LANGUAGE & TIME"
    currentState = state.LANGUAGE_NAME
}

function help() {
    console.log("Clicked Help button")
    document.getElementById("main_content").innerHTML = "HELP"
    currentState = state.HELP
}

function apply() {
    console.log("Clicked Apply button")
    let uid = sessionStorage.getItem("uid")
    console.log(uid)
    fetch("http://disclaimerszijneending.nl/config/123")
        .then(response => {
            return response.json()
        }).then(data => {
            console.log(data)
        }).catch(err => console.error(err))
    document.getElementById("main_content").innerHTML = "APPLY"
    switch (currentState) {
        case state.ALARM:
            break;
        case state.RADIO_STATUS:
            break;
        case state.NETWORK:
            break;
        case state.LANGUAGE_NAME:
            break;
        case state.HELP:
            break;
    }
}

function cancel() {
    console.log("Clicked Cancel button")
    document.getElementById("main_content").innerHTML = "CANCEL"
}




