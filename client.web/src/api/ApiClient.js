import axios from "axios";
//TODO: Implement async/await in fetch functions
export default class ApiClient {

    constructor(apiGatewayUrl) {
        this.apiGatewayUrl = apiGatewayUrl;
    }

    async login(username, password) {
        let response = await axios.post(this.apiGatewayUrl + "/auth",
            {
                "username": username,
                "password": password
            })
        console.log(response)
        return response
    }

    async fetchConfigs(id) {
        let response = await axios.get(this.apiGatewayUrl + "/config/" + id);
        return response.data;
    }
}