import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";


import Layout from "./components/Layout";
import Home from "./pages/Home";
import CreateTask from "./pages/CreateTask";
import EditTask from "./pages/EditTask";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="create" element={<CreateTask />} />
          <Route path="edit/:id" element={<EditTask />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;
