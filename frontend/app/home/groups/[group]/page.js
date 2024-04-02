"use client";
import { ApiProvider } from "@/app/_lib/utils";
import Group from "@/app/ui/home/groups/group.js/group";
import React from "react";
import { QueryClient } from "react-query";
import { QueryClientProvider } from "react-query";
const queryClient = new QueryClient();

const page = () => {
  return (
    <>
      <ApiProvider>
        <QueryClientProvider client={queryClient}>
          <Group />
        </QueryClientProvider>
      </ApiProvider>
    </>
  );
};

export default page;
