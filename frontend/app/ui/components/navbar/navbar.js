"use client"
import React from 'react'
import { useState } from "react";
import { useEffect } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { GoHomeFill } from "react-icons/go";
import { AiFillMessage } from "react-icons/ai";
import { IoNotifications } from "react-icons/io5";
import { FaUserGroup } from "react-icons/fa6";
import { IoPersonCircleSharp } from "react-icons/io5";
import { IoLogOut } from "react-icons/io5";
import Image from 'next/image';

// import { GoHomeFill, AiFillMessage, FaUserGroup, IoNotifications, IoPersonCircleSharp, IoLogOut } from 'react-icons/all';
// import clsx from 'clsx';

const links = [
    { name: 'Home', href: '/home', icon: GoHomeFill },
    { name: 'Messages', href: '/home/messages', icon: AiFillMessage, },
    { name: 'Communities', href: '/home/groups', icon: FaUserGroup, },
    { name: 'Notifications', href: '/home/notifications', icon: IoNotifications },
    { name: 'Profile', href: '/home/profile', icon: IoPersonCircleSharp },
    { name: 'Logout', href: '/home/logout', icon: IoLogOut },
];

function Navbar({ userData }) {
    // const location = useLocation();
    const pathname = usePathname();

    // const { userData, followers, isLoading, error } = useFetchData();

    return (
        <div className='shadowL  md:navbar xl:before:w-72 before:w-48 z-0 xl:w-60 md:block md:h-[700px] flex-col'>
            <div className='md:flex hidden relative z-0 flex-col w-full h-52 items-center justify-center'>
                {userData.avatar ?
                    <Image
                        src={`/assets/${userData.avatar}`}
                        alt="logo"
                        width={80} height={80}
                        className='rounded-full z-10'
                    />
                    : ("")}

                <h2 className='font-semibold text-xl text-center'>
                    {userData.first_name ? userData.first_name + " " + userData.last_name : ('')}
                </h2>
                <span className='text-md italic text-primary'>
                    {userData.user_name ? "@" + userData.user_name : ("")}
                </span>
            </div>

            <div className=" absolute content w-18 h-80 z-0 bg-other border-l-0 border-t-0 border-b-10 border-r-0 rounded-bl-0 rounded-tr-10 rounded-br-0 rounded-tl-0"></div>
            <div className='md:block flex gap-1 justify-center'>

                {links.map((link) => {
                    const LinkIcon = link.icon;
                    const isActive = (pathname === link.href) || pathname.includes(link.href) && link.href.length > 10;
                    // const isActive = pathname.includes(link.href)
                    // alert(pathname)

                    return (
                        <Link key={link.name} href={link.href}
                            className={` flex  h-[60px] items-center md:justify-start justify-center xl:w-72 md:w-48 gap-2 rounded-md mt-1
                         font-bold  hover:text-primary md:p-2 w-16 md:px-3 ${isActive ? 'isActive' : ''}`}
                        >
                            <LinkIcon className="xl:text-3xl text-xl" />
                            <p className="xl:text-md hidden md:block">{link.name}</p>
                        </Link >
                    );
                })}
            </div>
        </div>

    );
}

export default Navbar
