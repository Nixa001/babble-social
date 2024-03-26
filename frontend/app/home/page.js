"use client";
import React from "react";
import HomePage from "../ui/home/page";
import {QueryClientProvider } from "react-query";
import { queryClient } from "./groups/page";
import { ToastContainer } from "react-toastify";

const Page = () => {
  return (
    <div className="">
      <QueryClientProvider client={queryClient}>
        <HomePage/>
      </QueryClientProvider>
        <ToastContainer/>
    </div>
  );
};

export default Page;
