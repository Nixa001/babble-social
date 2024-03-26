'use client';
import React from 'react'
import HomePage from '../ui/home/page'
import { WebSocketProvider } from '../_lib/websocket'

const Page = () => {
  return (
    <WebSocketProvider> 
      <div className=''>
        <HomePage />
      </div>
    </WebSocketProvider>

  )
}

export default Page
