'use client'
import Image from 'next/image';
import { display, displayFollowers, displayGroups, followerHearder } from '../../components/sidebarRight/sidebar';
import { useContext, useEffect, useState } from 'react';
import useSocket, { WebSocketContext } from '@/app/_lib/websocket';
import { getSessionUser } from '@/app/_lib/utils';
import { useRouter } from 'next/navigation';
import { userID } from '../../components/navbar/navbar';
import { loginUser } from '@/app/api/api';
import InputEmoji from "react-input-emoji"

let idreceiver;
let idgroupreceiver;
const fetchUserGroup = async () => {
    const token = localStorage.getItem("token") || null;
    const url = `http://localhost:8080/messages?token=${encodeURIComponent(
        token
    )}`;
    try {
        const response = await fetch(url, {

            method: "GET",
        });
        const data = response.json();
        // console.log(data);
        return data;
    } catch (error) {
        console.error("Erreur ", error);
    }
}



const Messages = () => {
    const { sendMessageToServer, allMessages, resetAllMessages } = useContext(WebSocketContext)
    const [idUserReceiver, setIdUserReceiver] = useState(0);
    const [idGroupReceiver, setIdGroupReceiver] = useState(0);
    const [nameUser, setNameUser] = useState("");
    const [nameGroup, setNameGroup] = useState("");
    const [activeTab, setActiveTab] = useState("users");
    const route = useRouter();
    const [text, setText] = useState('')
    const [selectedEmoji, setSelectedEmoji] = useState(null)

    const [data, setData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            const result = await fetchUserGroup();
            setData(result); // Mettre à jour l'état avec les données récupérées
        };

        fetchData();
    }, []);
    let OnlineUser = data[0]
    let Groups = data[1]


    const handleTabClick = (tab) => {
        setActiveTab(tab);
    };
    const reset = () => {
        setIdUserReceiver(0);
        setIdGroupReceiver(0);
        setNameUser("");
        setNameGroup("");
    };
    useEffect(() => {
        return () => {
            reset();
        }
    }, [activeTab])
    useEffect(() => {
        return () => {
            resetAllMessages();
        };
    }, [activeTab]);
    useEffect(() => {
        // console.log("all messages", allMessages);
    }, [allMessages])
    useEffect(() => {
        idreceiver = idUserReceiver
    }, [idUserReceiver]);
    useEffect(() => {
        idgroupreceiver = idGroupReceiver
    }, [idGroupReceiver]);
    useEffect(() => {
        // console.log(nameUser);
    }, [nameUser]);
    useEffect(() => {
        // console.log(nameGroup);
    }, [nameGroup]);

    const handleUserClick = async (userId, name) => {
        setIdUserReceiver(userId);
        setNameUser(name);
        const sessionUser = await getSessionUser();
        const sessionUserId = sessionUser.id; // A
        sendMessageToServer({ type: 'id-receiver-event', data: { clickedUserId: userId, sessionUserId: sessionUserId } });
        route.push("/home/messages?id=" + userId)
    };

    //Traitement de message entre user et Groups

    const handleGroupClick = async (GroupId, nameGroup) => {
        setIdGroupReceiver(GroupId);
        setNameGroup(nameGroup)
        const sessionUser = await getSessionUser();
        const sessionUserId = sessionUser.id; // A
        console.log("Session user", sessionUserId);
        const token = localStorage.getItem('token');
        sendMessageToServer({ type: 'idGroup-receiver-event', data: { idgroup: GroupId, userID: sessionUserId } });
        route.push("/home/messages?idgroup=" + GroupId)
    };

    //Traitement lors de l'envoie de messages
    const handleSendMessage = async (e) => {
        e.preventDefault();
        let data = new FormData(e.target);
        let obj = {};
        data.forEach((value, key) => {
            obj[key] = value;
        });


        obj.message = text;
        if (selectedEmoji) {
            obj.message = selectedEmoji.native + " " + obj.message;
        }
        if (obj.message?.trim() !== "") {
            console.log("idreceiver in message", idreceiver);
            console.log("idgroupreceiver in message", idgroupreceiver);
            // Détermine le type de message en fonction de activeTab
            const messageType = activeTab === "group" ? "message-group-event" : "message-user-event";
            const receiverId = messageType === "message-group-event" ? idgroupreceiver : idreceiver;
            // Prépare les données du message
            const sessionUser = await getSessionUser();
            const sessionUserId = sessionUser.id; // A
            const messageData = {
                type: messageType,
                data: {
                    message: obj.message,
                    sendId: sessionUserId,
                    receiverId: receiverId,
                }
            };
            sendMessageToServer(messageData);
            e.target.reset();
            setText("");
        }
    };

    const displayTable = () => {
        if (activeTab === "users") {
            return displayFollowers(OnlineUser, handleUserClick);
        } else if (activeTab === "group") {
            return displayGroups(Groups, handleGroupClick);
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
                            className="profile_pic rounded-full cursor-pointer hover:opacity-60"
                            width={50}
                            height={50}
                        />
                        <div className='flex gap-2 items-center'>
                            <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                                {activeTab === "group" ? nameGroup : nameUser}
                            </h3>
                            <span className="username_post italic text-primary">

                            </span>
                        </div>
                    </div>
                    <div className="h-full overflow-y-auto">
                        <DisplayMessages messages={allMessages} />
                    </div>
                </div>

                <form className="flex justify-end mt-4 gap-2" onSubmit={handleSendMessage}>
                    <InputEmoji
                        value={text}
                        onChange={(event) => setText(event)}
                        cleanOnEnter
                        placeholder='Your message...'
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


export const DisplayMessages = ({ messages }) => {
    // Vérifie si messages est null ou non un tableau
    if (!Array.isArray(messages) || messages.length === 0) {
        return <p>Aucun message à afficher.</p>;
    }
    // Si messages n'est pas vide, mappe sur les messages comme avant
    return messages.map((message) => (
        <div key={message.id} className="message-container flex items-end mb-4 ">
            <div className="flex flex-col w-[70%]">
                <p className="text-sm w-fit bg-black font-semibold mb-1">{message.first_name}</p>
                <div className="font-semibold bg-primary p-4 rounded-lg break-words text-wrap w-fit max-w-[40%]">
                    {message.message_content}
                </div>
            </div>
        </div>
    ));
};

export default Messages;
