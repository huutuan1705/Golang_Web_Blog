import React from "react";
import Register from "./components/register";
import SnackbarProvider from "react-simple-snackbar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./components/Login";
import Home from "./components/Home";
import CreateBlog from "./components/CreateBlog";
import Navbar from "./components/Navbar";
import BlogDetail from "./components/BlogDetails";
import PersonalBlog from "./components/PersonalBlog";
import EditPost from "./components/EditPost";
function App() {
  return (
    <div className="">
      <SnackbarProvider>
        <Navbar />
        <Routes>
          <Route exact path="/" element={<Home />} />
          <Route exact path="/register" element={<Register />} />
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/create" element={<CreateBlog />} />
          <Route exact path="/detail/:id" element={<BlogDetail />} />
          <Route exact path="/personal" element={<PersonalBlog />} />
          <Route exact path="/edit/:id" element={<EditPost />} />
        </Routes>
      </SnackbarProvider>
    </div>
  );
}

export default App;