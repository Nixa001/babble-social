"use client"
import React from 'react'
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { GoHomeFill } from "react-icons/go";
import { AiFillMessage } from "react-icons/ai";
import { IoNotifications } from "react-icons/io5";
import { FaUserGroup } from "react-icons/fa6";
import { IoPersonCircleSharp } from "react-icons/io5";
import { IoLogOut } from "react-icons/io5";
import Image from 'next/image';
// import clsx from 'clsx';

const links = [
    { name: 'Home', href: '/home', icon: GoHomeFill },
    { name: 'Messages', href: '/home/messages', icon: AiFillMessage, },
    { name: 'Groups', href: '/home/groups', icon: FaUserGroup, },
    { name: 'Notifications', href: '/home/notifications', icon: IoNotifications },
    { name: 'Profile', href: '/home/profile', icon: IoPersonCircleSharp },
    { name: 'Logout', href: '/home/logout', icon: IoLogOut },
];

function Navbar() {
    const pathname = usePathname();
    return (
        <div className='shadowL navbar xl:before:w-72 before:w-48 z-0 xl:w-72 w-48 h-[700px] flex-col bg-text '>
            <div className='flex  relative z-0 flex-col w-full h-52 items-center justify-center'>
                <Image
                    src="/assets/profil.jpg" alt="logo"
                    width={80} height={80}
                    className='rounded-full z-10'
                />
                <h2 className='font-bold text-2xl text-center'>Nicolas Cor Faye</h2>
                <span className='text-xl italic text-primary'>@nixa</span>
            </div>
            <div className="navbar::before absolute content w-18 h-80 z-0 bg-other border-l-0 border-t-0 border-b-10 border-r-0 rounded-bl-0 rounded-tr-10 rounded-br-0 rounded-tl-0"></div>
            {/* <div className='flex-col h-full items-center justify-center'> */}

            {links.map((link) => {
                const LinkIcon = link.icon;
                return (
                    <Link key={link.name} href={link.href}
                        className=' shadowL flex h-[60px] grow items-center xl:w-72 w-48  gap-5 rounded-md mt-1 
                         font-bold hover:bg-primary hover:text-white md:p-2 md:px-3'
                    //     {
                    //         'bg-sky-100 text-blue-600': pathname === link.href,
                    //     ,
                    // }
                    >
                        <LinkIcon className="xl:text-5xl text-3xl" />
                        <p className="xl:text-xl hidden md:block">{link.name}</p>
                    </Link >
                );
            })}
            {/* </div> */}
        </div>
    );
}

export default Navbar