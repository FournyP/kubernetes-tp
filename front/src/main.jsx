import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import App from "./App";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
    <Router>
      <Routes>
        <Route exact path="/" element={<App />} />
        <Route path="/:name" element={<App />} />
      </Routes>
    </Router>
);
