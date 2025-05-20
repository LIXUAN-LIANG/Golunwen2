import axios from "axios";

let config={
    baseURL:"http://localhost:8080",
    timeout:10000,
    withCredentials:true,
}

const instance=axios.create(config);
export default instance