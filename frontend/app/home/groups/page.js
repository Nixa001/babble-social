'use client'
import React from 'react'
import Groups from '@/app/ui/home/groups/groups'

import { QueryClient } from 'react-query';
import { QueryClientProvider } from 'react-query';
const queryClient = new QueryClient();

const page = () => {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <Groups />
      </QueryClientProvider>
    </>
  )
}

export default page