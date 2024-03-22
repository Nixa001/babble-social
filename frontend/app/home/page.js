"use client";
import React from "react";
import HomePage from "../ui/home/page";
import {QueryClientProvider } from "react-query";
import { queryClient } from "./groups/page";

const Page = () => {
  return (
    <div className="">
      <QueryClientProvider client={queryClient}>
        <HomePage/>
      </QueryClientProvider>
    </div>
  );
};

export default Page;
