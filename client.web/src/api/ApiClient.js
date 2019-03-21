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
        return response.data
    }

    async fetchConfig(id) {
        let response = await axios.get(this.apiGatewayUrl + "/config/" + id);
        return response.data;
    }

    async fetchConfigs(...ids) {
        let configs = [];

        for (let i = 0; i < ids.length; i++) {
            let conf = await this.fetchConfig(ids[i])
            configs.push(conf)
        }

        return configs
    }


}