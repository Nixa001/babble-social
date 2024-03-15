'use client'
import React from 'react'
import Image from 'next/image'
import { useState } from 'react';



const followers = [
    { name: 'Vincent Ndour', src: "/assets/profilibg.jpg", alt: "profil" },
    { name: 'Ibrahima Gueye', src: "/assets/profilibg.jpg", alt: "profil", },
    { name: 'Madike Yade', src: "/assets/profilibg.jpg", alt: "profil", },
];
const groups = [
    { name: 'Call of duty', src: "/assets/profilibg.jpg", alt: "profil", },
    { name: 'Farcry 6 Team', src: "/assets/profilibg.jpg", alt: "profil" },
    { name: 'EA Fooball 24', src: "/assets/profilibg.jpg", alt: "profil", },
];

function Sidebar() {
    const [activeTab, setActiveTab] = useState("followers");

    const handleTabClick = (tab) => {
        setActiveTab(tab);
    };

    const displayTable = () => {
        if (activeTab === "followers") {
            return displayFollowers(followers)
                ;
        } else if (activeTab === "group") {
            return displayFollowers(groups);
        }
        return null;
    };

    return (
        <div className='shadowL z-0 xl:w-64 w-52 h-[700px] md:block hidden flex-col bg-text '>
            <div className='flex justify-around gap-x-1 pt-3 pb-3'>
                {followerHearder("Followers", "followers", activeTab, handleTabClick)}
                {followerHearder("Groups", "group", activeTab, handleTabClick)}
            </div>
            <div className="overflow-y-scroll chrome overflow-x-hidden h-1/2">
                {displayTable()}
            </div>
            <hr />
            <h3 className='font-bold p-2 rounded-sm text-center underline underline-offset-4 '>
                Friends
            </h3>
            <div className="online"></div>
        </div>
    );
}


export const followerHearder = (text, state, activeTab, handleTabClick) => {

    return <h3
        className={`font-bold cursor-pointer  hover:text-primary  p-2 rounded-sm 
    ${activeTab === state ? "text-primary underline underline-offset-4" : "opacity-70"}`}
        onClick={() => { handleTabClick(state) }}
    >
        {text}
    </h3>
}

export const displayFollowers = (data) => {
    return data.map((follower) => {
        return (
            <div key={follower.name} className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 ">
                {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}

                <Image
                    className="rounded-full "
                    src={follower.src}
                    alt={follower.alt}
                    width={35}
                    height={35}
                />
                <h4 className="font-bold text-sm ">{follower.name}</h4>
            </div>
        );
    })
}



export default Sidebar;
