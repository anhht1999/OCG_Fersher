import axios from "axios";
// import { API_DOMAIN } from "@/config";

const options = {};
options.baseURL = "https://online-shop-apis.herokuapp.com";

const instance = axios.create(options);

export default instance;