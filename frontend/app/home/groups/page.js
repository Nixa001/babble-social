'use client'
import { WebSocketProvider } from '@/app/_lib/websocket'
import React from 'react'

const Groups = () => {
  return (
    <WebSocketProvider> 
      <div className=''>groups</div>
    </WebSocketProvider>
  )
}

export default Groups