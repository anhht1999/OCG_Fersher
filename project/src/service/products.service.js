import axios from "axios";
import { API_DOMAIN } from "@/config";

export default {
  async getProducts({ page, limit, sort, order, search, categoryId}) {
    let filterCategory = "";
    if (categoryId) filterCategory = "&categoryId=" + categoryId;
    return axios
      .get(
        `${API_DOMAIN}/product?page=${page}&limit=${limit}&sort=${sort}&order=${order}&q=${search}${filterCategory}`
      )
      .then((response) => {
        const products = response.data.map((product) => {
          product.Url = 'http://localhost:3000' + product.Url;
          return product;
        });

        return {
          totalItems: response.headers["x-total-count"],
          data: products,
        };
      });
  },

  async getCategories() {
    return axios.get(`${API_DOMAIN}/category`).then((response) => {
      return response.data;
    });
  },

  async getProductById(productId) {
    return axios.get(`${API_DOMAIN}/product/${productId}`).then((response) => {
      // console.log(response.data)
      return response.data;
    });
  },
};