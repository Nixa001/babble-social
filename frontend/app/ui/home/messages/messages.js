'use client'
import { displayFollowers, followerHearder } from '../../components/sidebarRight/sidebar';
import { useState } from 'react';

const Messages = () => {
    const [activeTab, setActiveTab] = useState("users");
    const handleTabClick = (tab) => {
        setActiveTab(tab);
    };

    const handleSendMessage = (e) => {
        e.preventDefault()
    };
    const displayTable = () => {
        if (activeTab === "users") {
            return displayFollowers(users)
                ;
        } else if (activeTab === "group") {
            return displayFollowers(groups);
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
                        {followerHearder("Communities", "group", activeTab, handleTabClick)}
                        {/* <h2 className="text-xl font-semibold cursor-pointer hover:underline  mb-4">Users</h2>
                        <h2 className="text-xl font-semibold cursor-pointer hover:underline mb-4">Groups</h2> */}
                    </div>
                    {/* </div> */}
                    <ul className="list-none w-[100%] lg:px-5 md:px-0 px-5 sm:block flex overflow-y-scroll">
                        {/* {users.map((user) => (
                            <li key={user.id} className="mb-2 hover:opacity-70 cursor-pointer ">
                                <div className="flex flex-col flex-nowrap w-full items-center sm:flex-row ">
                                    <Image
                                        src={user.profilePicture}
                                        alt={user.username}
                                        width={40}
                                        height={40}
                                        className="rounded-full"
                                    />
                                    <span className="font-medium ml-2 text-sm lg:text-md w-full nowrap">{user.username}</span>
                                </div>
                            </li>
                        ))} */}


                        {displayTable()}
                    </ul>
                </div>
            </div>
            {/* Main content area */}
            <div className="md:h-[700px] h-full md:min-h-[600px] min-h-[500px] flex flex-col justify-between  md:w-full border-gray-700 p-2">
                <div>

                    <div className="info_post flex items-center gap-2 md:mb-5">
                        <img
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
                        {messages.map((message) => (
                            <div key={message.id} className="message-container flex items-end mb-4">
                                <div className="flex flex-col">
                                    <p className="text-sm font-semibold mb-1">{message.sender}</p>
                                    <div className=" font-semibold bg-primary p-4 rounded-lg">
                                        {message.content}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>

                <form className="flex justify-end mt-4 gap-2" onSubmit={handleSendMessage}>
                    <input
                        type="text"
                        className=" p-4 border border-gray-700 bg-transparent h-11 rounded-lg w-[90%] outline-none focus:ring-1 bg-primary focus:ring-primary"
                        placeholder="Your message..."
                    // value={message.content}
                    // onChange={(e) => setMessage(e.target.value)}
                    />
                    <button type="submit" className="bg-primary  hover:bg-second font-bold px-4 rounded-lg">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
                        </svg>

                        {/* <IoSend className='text-2xl text-center' /> */}
                    </button>



                </form>
            </div>
        </div>
    );
};

export default Messages;
const users = [
    { id: 1, first_name: 'Mamour', last_name: 'Drame', avatar: 'profilibg.jpg', alt: "profil" },
    { id: 2, first_name: 'Edouard', last_name: 'Mendy', avatar: 'profilibg.jpg', alt: "profil" },
    { id: 3, first_name: 'Vincent', last_name: 'Ndour', avatar: "profilibg.jpg", alt: "profil" },
];

const groups = [
    { id: 1, first_name: 'Call of duty', last_name: '', avatar: 'profilibg.jpg', alt: "profil" },
    { id: 2, first_name: 'Farcry 6 Team', last_name: '', avatar: 'profilibg.jpg', alt: "profil" },
    { id: 3, first_name: 'EA Fooball 24', last_name: '', avatar: "profilibg.jpg", alt: "profil" },
];

const messages = [
    {
        id: 1, sender: 'Mamou Drame', content: 'Hello everyone!', timestamp: '2023-11-16T12:00:00.000Z',
    },
    {
        id: 2, sender: 'Nicolas Faye', content: 'How are you all doing today?', timestamp: '2023-11-16T12:01:00.000Z',
    },
];


