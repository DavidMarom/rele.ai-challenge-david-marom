import Axios from 'axios';

var axios = Axios.create({
    withCredentials: true
});

export const httpService = {
    get(endpoint, data) {
        if (endpoint === 'undefined') { endpoint = '' }

        return ajax(endpoint, 'GET', data)
    }
}

async function ajax(endpoint, method = 'get', data = null) {
    try {
        const res = await axios({

            url: `//localhost:8081/${endpoint}`,
            method,
            data
        })
        return res.data;
    } catch (err) {
        console.log(`█Had Issues ${method}ing to the backend, endpoint: ${endpoint}, with data: ${data}`);
        console.dir(err);
        if (err.response && err.response.status === 401) {
            window.location.assign('/#/login');
        }
        throw err;
    }
}