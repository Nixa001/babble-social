'use client'
import Image from 'next/image';
import { IoSend } from "react-icons/io5";
import { displayFollowers, followerHearder } from '../../components/sidebarRight/sidebar';
import { useContext, useEffect, useState } from 'react';
import { postData } from '@/app/lib/utils';
import useSocket, { WebSocketContext } from '@/app/_lib/websocket';

const Messages = () => {
    const [activeTab, setActiveTab] = useState("users");
    const {sendMessageToServer, allMessages } = useContext(WebSocketContext)
    // const {sendMessage} = useSocket('ws://localhost:8080/ws');

    const handleTabClick = (tab) => {
        setActiveTab(tab);
    };


    //Traitement des message entre user
    useEffect(() => {
        console.log("all messages", allMessages);
    }, [allMessages])

    const handleUserClick = (userId) => {
        console.log("User clicked:", userId);
        sendMessageToServer({ type: 'id-receiver-event', data: userId });
        // console.log("all messages", allMessages);
        // Ici, vous pouvez ajouter la logique pour gérer l'ID de l'utilisateur cliqué


    };

    //Traitement de message entre user et Groups

    const handleGroupClick = (GroupId) => {
        console.log("Group clicked:", GroupId);
        sendMessageToServer({ type: 'idGroup-receiver-event', data: GroupId });
        // Ici, vous pouvez ajouter la logique pour gérer l'ID de l'utilisateur cliqué
    };

    //Traitement lors de l'envoie de messages
    const handleSendMessage = async (e) => {
        e.preventDefault();
        let data = new FormData(e.target);
        let obj = {};
        data.forEach((value, key) => {
            obj[key] = value;
        });
        console.log(obj.message);
        if (obj.message.trim() !== "") {
            // Détermine le type de message en fonction de activeTab
            const messageType = activeTab === "group" ? "message-group-event" : "message-user-event";

            // Prépare les données du message
            const messageData = {
                type: messageType,
                data: obj.message
            };

            // Envoie le message au serveur
            sendMessageToServer(messageData);
            e.target.reset();
        }
    };


    const displayTable = () => {
        if (activeTab === "users") {
            return displayFollowers(users, handleUserClick);

        } else if (activeTab === "group") {
            return displayFollowers(groups, handleGroupClick);
        }
        return null;
    };


    return (
        <div className="md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1200px] w-screen 
         flex flex-col sm:flex-row pr-5">
            {/* Left sidebar */}
            <div className="w-[100%] sm:w-[30%] h-[100%]  ">
                <div className="flex flex-col items-start justify-start border border-gray-700 h-full sm:h-[700px] overflow-hidden overflow-y-scroll">
                    {/* <div className='flex flex-col w-full bg-black justify-center'> */}
                    <h2 className="text-2xl mb-4 text-center w-full  font-bold sm:block hidden">Messages</h2>
                    <div className='flex gap-2 justify-evenly w-full '>

                        {followerHearder("Users", "users", activeTab, handleTabClick)}
                        {followerHearder("Community", "group", activeTab, handleTabClick)}

                    </div>
                    {/* </div> */}

                    <ul className="list-none w-[100%] lg:px-5 md:px-0 px-5 sm:block flex overflow-y-scroll">
                        {displayTable()}
                    </ul>

                </div>
            </div>
            {/* Main content area */}
            <div className="md:h-[700px] h-full md:min-h-[600px] min-h-[500px] flex flex-col justify-between  md:w-full border-gray-700 p-2">
                <div>

                    <div className="info_post flex items-center gap-2 md:mb-5">
                        <Image
                            src='/assets/profilibg.jpg'
                            alt="Profile picture"
                            //   onClick={handleProfileClick}
                            className="profile_pic rounded-full cursor-pointer hover:opacity-60"
                            width={50}
                            height={50}
                        />
                        <div className='flex gap-2 items-center'>
                            <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                                Mamour Drame
                            </h3>
                            <span className="username_post italic text-primary">
                                @Darze
                            </span>
                        </div>
                    </div>
                    <div className="h-full overflow-y-auto">
                        {displayMessages(allMessages.length > 0 ? allMessages[0] : messages)}
                    </div>
                </div>

                <form className="flex justify-end mt-4 gap-2" onSubmit={handleSendMessage}>
                    <input
                        type="text"
                        name="message"
                        className=" p-4 border border-gray-700 bg-transparent h-11 rounded-lg w-[90%] outline-none focus:ring-1 bg-primary focus:ring-primary"
                        placeholder="Your message..."
                    />
                    <button type="submit" className="bg-primary  hover:bg-second font-bold px-4 rounded-lg">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
                        </svg>
                    </button>
                </form>
            </div>
        </div>
    );
};


const displayMessages = (messages) => {
    return messages.map((message) => (
        <div key={message.id} className="message-container flex items-end mb-4">
            <div className="flex flex-col">
                <p className="text-sm font-semibold mb-1">{message.first_name}</p>
                <div className="font-semibold bg-primary p-4 rounded-lg">
                    {message.message_content}
                </div>
            </div>
        </div>
    ));
};
export default Messages;
const users = [
    { id: 1, name: 'Mamour Drame', src: '/assets/profilibg.jpg', alt: "profil" },
    { id: 2, name: 'Edouard Mendy', src: '/assets/profilibg.jpg', alt: "profil" },
    { id: 3, name: 'Vincent Ndour', src: "/assets/profilibg.jpg", alt: "profil" },
    { id: 4, name: 'Ibrahima Gueye', src: "/assets/profilibg.jpg", alt: "profil", },
    { id: 5, name: 'Madike Yade', src: "/assets/profilibg.jpg", alt: "profil", },
];

const messages = [
    {
        id: 1, first_name: 'Mamou Drame', message_content: 'Hello everyone!', timestamp: '2023-11-16T12:00:00.000Z',
    },
    {
        id: 2, first_name: 'Nicolas Faye', message_content: 'How are you all doing today?', timestamp: '2023-11-16T12:01:00.000Z',
    },
];


const groups = [
    { id: 1, name: 'Call of duty', src: "/assets/profilibg.jpg", alt: "profil", },
    { id: 2, name: 'Farcry 6 Team', src: "/assets/profilibg.jpg", alt: "profil" },
    { id: 3, name: 'EA Fooball 24', src: "/assets/profilibg.jpg", alt: "profil", },
];

