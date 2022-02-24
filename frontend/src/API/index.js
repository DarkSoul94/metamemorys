import axios from "axios";
import { LOGIN_ROUTE } from "../utils/consts";
import { useNavigate } from "react-router-dom";

const $host = axios.create({
    baseURL: process.env.REACT_APP_API_URL
})

const $authHost = axios.create({
    baseURL: process.env.REACT_APP_API_URL
})

const authInterceptor = config => {
    const headers = {
        "authorization" : localStorage.getItem("token"),
    }
    config.headers = headers
    return config
};

$authHost.interceptors.request.use(authInterceptor)

$authHost.interceptors.response.use(function (config){
    return config
}, function (error){
    if (error.response.status === 401){
        const navigator = useNavigate();
        navigator(LOGIN_ROUTE)
    }

    return Promise.reject(error)
});

export {
    $host,
    $authHost
}