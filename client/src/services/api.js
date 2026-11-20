// api.js
import axios from 'axios';

const API_URL = 'http://localhost:3000';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// CRUD operations

export const getAllProducts = async () => {
  try {
    const response = await api.get('/products');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createProduct = async (productData) => {
  try {
    const payload = {
      ...productData,
      price: parseFloat(productData.price), // Ensure price is a number
    };
    const response = await api.post('/create-product', payload);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const updateProduct = async (productId, productData) => {
  try {
    const payload = {
      ...productData,
      price: parseFloat(productData.price), // Ensure price is a number
    };
    const response = await api.put(`/products/${productId}`, payload);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const deleteProduct = async (productId) => {
  try {
    const response = await api.delete(`/products/${productId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export default api;
