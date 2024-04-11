'use client'
import React from 'react'
import Image from 'next/image'
import { useState } from 'react';
import { userID } from '../navbar/navbar';



const followers = [
    { id: 1, name: 'Vincent Ndour', src: "/assets/profilibg.jpg", alt: "profil" },
    { id: 2, name: 'Ibrahima Gueye', src: "/assets/profilibg.jpg", alt: "profil", },
    { id: 3, name: 'Madike Yade', src: "/assets/profilibg.jpg", alt: "profil", },
];
const groups = [
    { id: 1, name: 'Call of duty', src: "/assets/profilibg.jpg", alt: "profil", },
    { id: 2, name: 'Farcry 6 Team', src: "/assets/profilibg.jpg", alt: "profil" },
    { id: 3, name: 'EA Fooball 24', src: "/assets/profilibg.jpg", alt: "profil", },
];

function Sidebar() {
    const [activeTab, setActiveTab] = useState("followers")

    const handleTabClick = (tab) => {
        setActiveTab(tab);
    };
    const handleSidebarUserClick = () => {
        alert("User clicked")
    }

    const displayTable = () => {
        if (activeTab === "followers") {
            return display(followers, handleSidebarUserClick)
                ;
        } else if (activeTab === "group") {
            return display(groups, handleSidebarUserClick);
        }
        return null;
    };

    return (
        <div className='shadowL z-0 xl:w-64 w-52 h-[700px] md:block hidden flex-col bg-text '>
            <div className='flex justify-around gap-x-1 pt-3 pb-3'>
                {followerHearder("Followers", "followers", activeTab, handleTabClick)}
                {followerHearder("Community", "group", activeTab, handleTabClick)}
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
        onClick={() => {
            handleTabClick(state)

        }}
    >
        {text}
    </h3>
}

export const displayFollowers = (data, handleUserClick) => {
    return data?.map((follower) => {
        if (!data || data.length === 0) {
            return <div>Vous n'avez encore de follower</div>;
        }
        if (follower.id != userID) {
            return (
                <div key={follower.id} className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 "
                    onClick={() => handleUserClick(follower.id, follower.first_name + ' ' + follower.last_name)}
                >
                    <Image
                        className="rounded-full "
                        // src={follower.src}
                        src='/assets/profilibg.jpg'
                        alt="missing"
                        width={40}
                        height={40}
                    />
                    <h4 className="font-bold" >{follower.first_name + ' ' + follower.last_name}</h4>
                </div>
            );
        }
        return null;
    }).filter(Boolean); // Filtre les éléments null pour ne garder que ceux qui doivent être affichés
}
export const displayGroups = (data, handleUserClick) => {
    if (!data || data.length === 0) {
        return <div>Vous n'avez encore de groupe</div>;
    }
    return data?.map((follower) => {
        return (
            <div
                key={follower.id}
                className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 "
                onClick={() => handleUserClick(follower.id, follower.name)}
            >
                <Image
                    className="rounded-full "
                    src='/assets/profilibg.jpg'
                    alt="missing"
                    width={35}
                    height={35}
                />
                <h4 className="font-bold">{follower.name}</h4>
            </div>
        );
    });
}

export const display = (data, handleUserClick) => {
    return data.map((follower) => {
        return (
            <div
                key={follower.id} 
                className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 "
                onClick={() => handleUserClick(follower.id, follower.name)}
            >
                <Image
                    className="rounded-full "
                    src={follower.src}
                    alt={follower.alt}
                    width={35}
                    height={35}
                />
                <h4 className="font-bold">{follower.name}</h4>
            </div>
        );
    });
}


export default Sidebar;
