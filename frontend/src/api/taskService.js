import axios from "axios";

const api = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL || "/api",
  headers: { "Content-Type": "application/json" },
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error("API Error:", error.response?.data || error.message);
    return Promise.reject(
      error.response?.data?.message || "Error while requesting server"
    );
  }
);

export const taskService = {
  async getAll() {
    const res = await api.get("/tasks");
    return res.data;
  },

  async getById(id) {
    const res = await api.get(`/tasks/${id}`);
    return res.data;
  },

  async create(task) {
    const res = await api.post("/tasks", task);
    return res.data;
  },

  async update(id, task) {
    const res = await api.put(`/tasks/${id}`, task);
    return res.data;
  },

  async remove(id) {
    const res = await api.delete(`/tasks/${id}`);
    return res.data;
  },
};
