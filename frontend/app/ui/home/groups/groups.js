'use client'
import React, { useState } from 'react';
import Image from 'next/image';
import Link from 'next/link';
import { useEffect } from 'react';
import { CreateGroup } from '../../components/modals/createGroup';

const Groups = () => {
    const [formCreateGr, setFormCreateGr] = useState(false)
    // const groupData = Data
    // const groupJoined = DataJoined

    const [groupData, setGroupData] = useState([]);
    const [groupJoined, setGroupJoined] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/groups')
            .then(response => response.json())
            .then(data => {
                setGroupData(data)
            })
            .catch(error => console.error('Erreur lors du chargement des données des groupes:', error));

        fetch('/api/groups/joined')
            .then(response => response.json())
            .then(data => setGroupJoined(data))
            .catch(error => console.error('Erreur lors du chargement des données des groupes rejoints:', error));
    }, []);


    return (
        <div className='md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1200px] w-screen h-full 
             flex flex-col  '>

            <div className='w-[90%] justify-between flex items-center mb-5 gap-5'>
                <h1 className='text-xl font-bold '>
                    My groups
                </h1>
                <button className="inline-flex items-center px-4 py-2 text-m font-semibold text-center 
                            text-white bg-primary rounded-lg hover:bg-second"
                    onClick={() => {
                        setFormCreateGr(true)
                    }}>
                    Create a new group
                </button>
            </div>
            <div className='w-full flex gap-3 flex-wrap '>

                {groupJoined.map((group) => (
                    <GroupCard key={group.id} isMember={true} {...group} />
                ))}

            </div>
            <h1 className='text-xl font-bold my-5'>
                Discover new groups
            </h1>

            <div className='w-full flex gap-3 flex-wrap pb-10 '>
                {groupData.map((group) => (
                    <GroupCard key={group.id} isMember={false} {...group} />
                ))}
            </div>
            <CreateGroup isVisible={formCreateGr} onClose={() => setFormCreateGr(false)} />
        </div>
    );
};

export default Groups;


const GroupCard = ({ isMember, image, name, description, href, functionOnclick }) => {
    description = description.slice(0, 50) + "..."
    return (
        <>
            {isMember ? (
                < Link href={href} className="inline-flex items-center text-m font-semibold text-center text-white rounded-l">
                    <div
                        className="w-[200px] border rounded-lg shadow-2xl bg-black bg-clip-padding backdrop-filter backdrop-blur-md bg-opacity-5 border-gray-700 hover:bg-opacity-15 cursor-pointer"
                    >
                        <div className="flex flex-col items-center py-3">
                            <Image
                                src={image}
                                alt={name}
                                width={500}
                                height={500}
                                className="w-24 h-24 mb-3 rounded-full hover:scale-110 transition duration-300 ease-in shadow-lg"
                            />
                            <h5 className="mb-1 text-xl font-medium text-white text-center">{name}</h5>
                            <span className="max-h-14 overflow-hidden text-sm text-gray-300 text-center">{description}</span>

                        </div>
                    </div>
                </Link >
            ) : (
                <div
                    className="w-[200px] border rounded-lg shadow-2xl bg-gray-700 bg-clip-padding backdrop-filter backdrop-blur-md bg-opacity-10 border-gray-700 hover:bg-opacity-25 cursor-pointer"
                >
                    <div className="flex flex-col items-center py-3">
                        <Image
                            src={`/assets/${image}`}
                            alt={name}
                            width={200}
                            height={200}
                            className="w-24 h-24 mb-3 rounded-full shadow-lg"
                        />
                        <h5 className="mb-1 text-xl font-medium text-white text-center">{name}</h5>
                        <span className="max-h-14 overflow-hidden text-sm text-gray-300 text-center">{description}</span>

                        <div className="flex mt-4 md:mt-6">
                            <button onClick={() => { functionOnclick() }} className="inline-flex items-center px-4 py-2 text-m font-semibold text-center text-white bg-primary rounded-lg hover:bg-second">
                                Join
                            </button>
                        </div>

                    </div>
                </div>
            )}


        </>

    );
};

function JoinGroup() {
    alert('Send join');
}

const Data = [
    {
        id: 1,
        image: "/assets/ea.jpg",
        name: "EA Football 24",
        description: "Un groupe pour les fans de football du monde entier",
        href: "/groups/join/EA Football 24",
        functionOnclick: JoinGroup,

    },
    {
        id: 2,
        image: "/assets/cod.jpg",
        name: "Call of duty",
        description: "Un groupe gamers du monde entier",
        href: "/groups/join/Call of duty",
        functionOnclick: JoinGroup,
    },
    {
        id: 3,
        image: "/assets/100daysofcode.jpg",
        name: "100 days of code",
        description: "100DaysOfCode is a coding challenge for developers, with thousands of members worldwide...",
        href: "/groups/join/EA Football 24",
        functionOnclick: JoinGroup,
    },
    {
        id: 4,
        image: "/assets/cod.jpg",
        name: "Call of duty",
        description: "Un groupe gamers du monde entier",
        href: "/groups/join/Call of duty",
        functionOnclick: JoinGroup,
    },
    {
        id: 5,
        image: "/assets/ea.jpg",
        name: "EA Football 24",
        description: "Un groupe pour les fans de football du monde entier",
        href: "/groups/join/EA Football 24",
        functionOnclick: JoinGroup,
    },
];
const DataJoined = [
    {
        id: 111,
        image: "/assets/ea.jpg",
        name: "EA Football 24",
        description: "Un groupe pour les fans de football du monde entier",
        href: "/home/groups/group/",
    },
    {
        id: 211,
        image: "/assets/cod.jpg",
        name: "Call of duty",
        description: "Un groupe gamers du monde entier",
        href: "/home/groups/group/",
    },
    {
        id: 311,
        image: "/assets/100daysofcode.jpg",
        name: "100 days of code",
        description: "100DaysOfCode is a coding challenge for developers, with thousands of members worldwide...",
        href: "/home/groups/group/",
    },
    {
        id: 411,
        image: "/assets/cod.jpg",
        name: "Call of duty",
        description: "Un groupe gamers du monde entier",
        href: "/home/groups/group/",
    },
    {
        id: 511,
        image: "/assets/ea.jpg",
        name: "EA Football 24",
        description: "Un groupe pour les fans de football du monde entier",
        href: "/home/groups/group/",
    },
];