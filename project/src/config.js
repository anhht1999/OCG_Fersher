  
const isProductionEnv = process.env.NODE_ENV === "production";

export const API_DOMAIN = isProductionEnv
  ? "https://online-shop-apis.herokuapp.com" 
  : "http://localhost:3000/api";

export const BASE_URL = isProductionEnv ? "/project/" : "/";