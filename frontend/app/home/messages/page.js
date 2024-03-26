'use client'
import { WebSocketProvider } from '@/app/_lib/websocket'
import Messages from '@/app/ui/home/messages/messages'
import React from 'react'

const page = () => {

  return (
    <WebSocketProvider>
      <div>
        <Messages />
      </div>
    </WebSocketProvider>
  )
}

export const displayFollowers = (data, handleUserClick) => {
  return data.map((follower) => {
    return (
      <div key={follower.name} className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 "
        onClick={() => handleUserClick(follower.id)}
      >
        <Image
          className="rounded-full "
          src={follower.src}
          alt={follower.alt}
          width={40}
          height={40}
        />
        <h4 className="font-bold" >{follower.name}</h4>
      </div>
    );
  })
}


export default page