"use client";
import { useEffect } from "react";
import { useState } from "react";
import { Header } from "../ui/components/header/header";
import Navbar from "../ui/components/navbar/navbar";
import Sidebar from "../ui/components/sidebarRight/sidebar";
import { displayFollowers } from "./utils";
import { ToastContainer } from "react-toastify";

export default function Layout({ children }) {
  const { userData, followers, isLoading, error } = useFetchData();

  return (
    <div className="h-screen">
      <div className="fixed">
        <Header />
      </div>
      <div
        className="md:flex md:flex-row flex flex-col-reverse 
      justify-between h-[99%] md:justify-between md:h-full  overflow-hidden">
        <div className="md:mt-20">
          <Navbar user={userData} />
          <ToastContainer />
        </div>{" "}
        {/* Enveloppez le contenu avec WebSocketProvider */}
        <div className="mt-20 overflow-x-hidden overflow-y-scroll chrome pl-3 mr-3">
          {children}
        </div>
        <div className="md:mt-20 hidden md:block">
          <Sidebar followers={followers} />
        </div>
      </div>
    </div>
  );
}

function useFetchData() {
  const [userData, setUserData] = useState({});
  const [followers, setFollowers] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchUserData = async () => {
      setIsLoading(true);
      setError(null);

      try {
        const url = `http://localhost:8080/userInfo`;
        const response = await fetch(url, { method: "GET" });
        const data = await response.json();

        setUserData(data.user);
        setFollowers(data.followers);
      } catch (err) {
        setError(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, []);
  return { userData, followers, isLoading, error };
}
