"use client";
import React from "react";
import {QueryClientProvider } from "react-query";
import { queryClient } from "./groups/page";
import { ToastContainer } from "react-toastify";
import HomePage from "../ui/home/page";

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
