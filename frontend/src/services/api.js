// src/services/api.js
import axios from 'axios';

const API_URL = 'http://127.0.0.1:3000'; // Replace with your API URL

const api = axios.create({
  baseURL: API_URL,
});

const username = "admin";
const password = "captainvader";

// Function to get captain data
export const getCaptain = async () => {
  try {
    const response = await api.get('/captain', {
        auth: {
          username: username,
          password: password
        }
      });
    return response.data[0];
  } catch (error) {
    console.error('Error fetching captain data:', error);
    throw error; // Propagate the error to the calling function
  }
};

// Function to get officer data
export const getOfficer = async () => {
    try {
        const response = await api.get('/officers', {
            auth: {
              username: username,
              password: password
            }
          });
        return response.data;
      } catch (error) {
        console.error('Error fetching officer data:', error);
        throw error; // Propagate the error to the calling function
      }
};

// Function to get crew data
export const getCrew = async () => {
    try {
        const response = await api.get('/crew', {
            auth: {
              username: username,
              password: password
            }
          });
        console.log('response in api', response)
        return response.data;
      } catch (error) {
        console.error('Error fetching officer data:', error);
        throw error; // Propagate the error to the calling function
      }
};
