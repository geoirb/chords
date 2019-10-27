import axios from 'axios';

const baseIP = '127.0.0.1:1323';
const baseURL =  '//' + baseIP;

const baseAPI = axios.create({
    baseURL: baseURL,
});

const getConfig = (params) => ({
    baseURL: baseURL,
    params: params,
});

const API = {
    GET: (path, params) => baseAPI.get(path, getConfig(params)),
    POST: (path, data) => baseAPI.post(path, data, getConfig()),
    PUT: (path, data) => baseAPI.put(path, data, getConfig()),
    DELETE: (path, params) => baseAPI.delete(path, getConfig(params)),
};

class GoService{
    getChords = (author, titel) => API.GET('/chord/' + author + '/' + titel);
};

// http://127.0.0.1:1323/chord/Dolly/Jolene

export default GoService;

