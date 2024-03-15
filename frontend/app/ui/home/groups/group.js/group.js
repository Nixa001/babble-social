'use client'
import Image from 'next/image'
import React, { useState } from 'react'
import DisplayPost from '../../displayPost'
import { CreateEvent } from '@/app/ui/components/modals/createEvent'
import { CreatePost } from '@/app/ui/components/modals/createPost'
import { Suggest } from '@/app/ui/components/modals/suggest'
import { DisplayMembers } from '@/app/ui/components/modals/displayMembers'
// import { useState } from 'react'

const Group = () => {
    const [formCreateEv, setFormCreateEv] = useState(false)
    const [formCreateP, setFormCreateP] = useState(false)
    const [suggestFriend, setSuggestFriend] = useState(false)
    const [members, setMembers] = useState(false)


    return (
        <div className='md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1100px] w-screen h-full 
                        flex flex-col gap-2'>
            <div className="w-full h-60 mb-3" >
                <Image
                    src={"/assets/cover.webp"}
                    alt='cover'
                    width={1000}
                    height={1000}
                    className='w-full max-h-[250px] object-cover scale-100 hover:scale-105  rounded-sm  transition duration-300 ease-in shadow-lg'
                />
            </div>
            <div className='flex items-center justify-between 2xl:w-[90%] w-[95%]'>
                <div>
                    <h1 className='text-2xl font-bold underline underline-offset-4 mb-2 '>
                        EA Football 24
                    </h1>
                    <p className='text-lg'>
                        Un groupe pour les fans de football du monde entier
                    </p>
                </div>
                <div className='flex gap-2 lg:flex-row flex-col' >
                    <button
                        onClick={() => {
                            setFormCreateP(true)
                        }}
                        className="inline-flex items-center px-1 py-1 text-sm font-bold text-center max-h-[50px]  bg-gray-700 border border-gray-500 rounded-lg hover:bg-second">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path fillRule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25ZM12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25V9Z" clipRule="evenodd" />
                        </svg>
                        Post
                    </button>
                    <button onClick={() => {
                        setFormCreateEv(true)
                    }}
                        className="inline-flex items-center px-1 py-1 text-sm font-bold border border-gray-500 bg-gray-800 text-center max-h-[50px] rounded-lg hover:bg-second">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path d="M12.75 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM7.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM8.25 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM9.75 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM10.5 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM12.75 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM14.25 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 13.5a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" />
                            <path fillRule="evenodd" d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75Zm13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5Z" clipRule="evenodd" />
                        </svg>
                        Event
                    </button>
                    <button onClick={() => {
                        setSuggestFriend(true)
                    }}
                        className="inline-flex items-center px-1 py-1 text-sm font-bold text-center max-h-[50px]  bg-gray-900 border border-gray-500 rounded-lg hover:bg-second">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path d="M5.25 6.375a4.125 4.125 0 1 1 8.25 0 4.125 4.125 0 0 1-8.25 0ZM2.25 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63 13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63l-.001-.122ZM18.75 7.5a.75.75 0 0 0-1.5 0v2.25H15a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H21a.75.75 0 0 0 0-1.5h-2.25V7.5Z" />
                        </svg>
                        Suggest
                    </button>
                </div>
            </div>
            <p className='flex items-center gap-2 shadow w-fit font-bold rounded-md cursor-pointer py-2 px-3 border-gray-700' onClick={() => {
                setMembers(true)
            }}>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                    <path fillRule="evenodd" d="M8.25 6.75a3.75 3.75 0 1 1 7.5 0 3.75 3.75 0 0 1-7.5 0ZM15.75 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM2.25 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM6.31 15.117A6.745 6.745 0 0 1 12 12a6.745 6.745 0 0 1 6.709 7.498.75.75 0 0 1-.372.568A12.696 12.696 0 0 1 12 21.75c-2.305 0-4.47-.612-6.337-1.684a.75.75 0 0 1-.372-.568 6.787 6.787 0 0 1 1.019-4.38Z" clipRule="evenodd" />
                    <path d="M5.082 14.254a8.287 8.287 0 0 0-1.308 5.135 9.687 9.687 0 0 1-1.764-.44l-.115-.04a.563.563 0 0 1-.373-.487l-.01-.121a3.75 3.75 0 0 1 3.57-4.047ZM20.226 19.389a8.287 8.287 0 0 0-1.308-5.135 3.75 3.75 0 0 1 3.57 4.047l-.01.121a.563.563 0 0 1-.373.486l-.115.04c-.567.2-1.156.349-1.764.441Z" />
                </svg>

                Members: <span className='text-primary italic'>20k</span>

            </p>
            <div>
                <hr className="h-px m-2 bg-gray-700 border-0"></hr>
            </div>
            <div className='w-full flex justify-between'>
                <div className='w-[20%]'>
                    <h1 className=' text-xl flex items-center gap-2 font-extrabold w-fit border border-gray-700 shadow-lg px-2 rounded-md '>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5" />
                        </svg>
                        Joined Events
                    </h1>
                    {displayEvents(events)}
                    <h1 className=' text-xl flex items-center gap-2 font-extrabold w-fit border border-gray-700 shadow-lg px-2 rounded-md mt-4 '>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
                        </svg>
                        New Events
                    </h1>
                    {displayEvents(events)}
                </div>
                <div className='w-[75%] '>
                    <DisplayPost postData={postData2} onLikeClick={onLikeClick} onDislikeClick={onDislikeClick}
                        onCommentClick={onCommentClick} onProfileClick={onProfileClick}
                    />

                    <DisplayPost postData={postData2} onLikeClick={onLikeClick} onDislikeClick={onDislikeClick}
                        onCommentClick={onCommentClick} onProfileClick={onProfileClick}
                    />
                </div>
            </div>
            <CreateEvent isVisible={formCreateEv} onClose={() => setFormCreateEv(false)} />
            <CreatePost isVisible={formCreateP} onClose={() => setFormCreateP(false)} />
            <Suggest isVisible={suggestFriend} onClose={() => setSuggestFriend(false)} />
            <DisplayMembers isVisible={members} onClose={() => setMembers(false)} />

        </div>
    )
}

export default Group

const postData2 = {
    profilePicture: "/assets/profilibg.jpg",
    userName: "Maurice Dassylva",
    userHandle: "@Maurice",
    timePosted: "2h",
    hashtags: ["Tech", "Sport"],
    title: "Ceci est mon titre",
    postImage: "/assets/imagepost2.jpg",
    likesCount: 19,
    dislikesCount: 20,
    commentsCount: 3,
}

const onLikeClick = () => {
    alert('like')
};

const onDislikeClick = () => {
    alert('dislike')

};

const onCommentClick = () => {
    alert('Comment disp')

};

const onProfileClick = () => {
    alert('profile disp')

};

export const displayEvents = (events) => {
    return events.map((event) => {
        return (
            <div key={event.id} className=" hover:text-primary flex items-center cursor-pointer justify-start gap-2 p-2 ">
                {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}

                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                    <path d="M12.75 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM7.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM8.25 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM9.75 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM10.5 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM12 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM12.75 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM14.25 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 17.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 15.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5ZM15 12.75a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0ZM16.5 13.5a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Z" />
                    <path fillRule="evenodd" d="M6.75 2.25A.75.75 0 0 1 7.5 3v1.5h9V3A.75.75 0 0 1 18 3v1.5h.75a3 3 0 0 1 3 3v11.25a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3V7.5a3 3 0 0 1 3-3H6V3a.75.75 0 0 1 .75-.75Zm13.5 9a1.5 1.5 0 0 0-1.5-1.5H5.25a1.5 1.5 0 0 0-1.5 1.5v7.5a1.5 1.5 0 0 0 1.5 1.5h13.5a1.5 1.5 0 0 0 1.5-1.5v-7.5Z" clipRule="evenodd" />
                </svg>

                <p className="font-semibold ">{event.description}</p>
            </div>
        );
    })
}

const events = [
    { id: 1, description: 'Description event 1', src: "/assets/profilibg.jpg", alt: "profil" },
    { id: 2, description: 'Description event 2', src: "/assets/profilibg.jpg", alt: "profil", },
    { id: 3, description: 'Description event 3', src: "/assets/profilibg.jpg", alt: "profil", },
];