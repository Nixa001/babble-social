"use client";
import { followerHearder } from "../../components/sidebarRight/sidebar";
import { useContext, useEffect, useState } from "react";
import useSocket, { WebSocketContext } from "@/app/_lib/websocket";
import { getSessionUser } from "@/app/_lib/utils";

let idreceiver;
let idgroupreceiver;
const Messages = () => {
  const { sendMessageToServer, allMessages, resetAllMessages } =
    useContext(WebSocketContext);
  const [idUserReceiver, setIdUserReceiver] = useState(0);
  const [idGroupReceiver, setIdGroupReceiver] = useState(0);
  const [nameUser, setNameUser] = useState("");
  const [nameGroup, setNameGroup] = useState("");
  const [activeTab, setActiveTab] = useState("users");

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
    };
  }, [activeTab]);
  useEffect(() => {
    return () => {
      resetAllMessages();
    };
  }, [activeTab]);
  useEffect(() => {
    // console.log("all messages", allMessages);
  }, [allMessages]);
  useEffect(() => {
    idreceiver = idUserReceiver;
  }, [idUserReceiver]);
  useEffect(() => {
    idgroupreceiver = idGroupReceiver;
  }, [idGroupReceiver]);
  useEffect(() => {
    // console.log(nameUser);
  }, [nameUser]);
  useEffect(() => {
    // console.log(nameGroup);
  }, [nameGroup]);

  const handleUserClick = async (userId, name) => {
    // console.log("User clicked:", userId);
    setIdUserReceiver(userId);
    setNameUser(name);
    // Récupérer l'ID de l'utilisateur en session
    const sessionUser = await getSessionUser();
    const sessionUserId = sessionUser.id; // A
    // console.log("idreceiver", idreceiver);
    // console.log("Session user", sessionUserId);
    sendMessageToServer({
      type: "id-receiver-event",
      data: { clickedUserId: userId, sessionUserId: sessionUserId },
    });
  };

  //Traitement de message entre user et Groups

  const handleGroupClick = async (GroupId, nameGroup) => {
    // console.log("Group clicked:", GroupId);
    setIdGroupReceiver(GroupId);
    setNameGroup(nameGroup);
    const sessionUser = await getSessionUser();
    const sessionUserId = sessionUser.id; // A
    // console.log("Session user", sessionUserId);
    const token = localStorage.getItem("token");
    sendMessageToServer({
      type: "idGroup-receiver-event",
      data: { idgroup: GroupId, userID: sessionUserId },
    });
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
    // console.log(obj.message);
    if (obj.message.trim() !== "") {
      // console.log("idreceiver in message", idreceiver);
      // console.log("idgroupreceiver in message", idgroupreceiver);
      // Détermine le type de message en fonction de activeTab
      const messageType =
        activeTab === "group" ? "message-group-event" : "message-user-event";
      const receiverId =
        messageType === "message-group-event" ? idgroupreceiver : idreceiver;
      // Prépare les données du message
      const sessionUser = await getSessionUser();
      const sessionUserId = sessionUser.id; // A
      const messageData = {
        type: messageType,
        data: {
          message: obj.message,
          sendId: sessionUserId,
          receiverId: receiverId,
        },
      };
      // Envoie le message au serveur
      sendMessageToServer(messageData);
      e.target.reset();
    }
  };

  const displayTable = () => {
    if (activeTab === "users") {
      return display(users, handleUserClick);
    } else if (activeTab === "group") {
      return display(groups, handleGroupClick);
    }
    return null;
  };

  return (
    <div
      className="md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1200px] w-screen 
         flex flex-col sm:flex-row pr-5"
    >
      {/* Left sidebar */}
      <div className="w-[100%] sm:w-[30%] h-[100%]  ">
        <div className="flex flex-col items-start justify-start border border-gray-700 h-full sm:h-[700px] overflow-hidden overflow-y-scroll">
          {/* <div className='flex flex-col w-full bg-black justify-center'> */}
          <h2 className="text-2xl mb-4 text-center w-full  font-bold sm:block hidden">
            Messages
          </h2>
          <div className="flex gap-2 justify-evenly w-full ">
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
              src="/assets/profilibg.jpg"
              alt="Profile picture"
              className="profile_pic rounded-full cursor-pointer hover:opacity-60"
              width={50}
              height={50}
            />
            <div className="flex gap-2 items-center">
              <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                {activeTab === "group" ? nameGroup : nameUser}
              </h3>
              <span className="username_post italic text-primary">@Darze</span>
            </div>
          </div>
          <div className="h-full overflow-y-auto">
            {displayMessages(allMessages)}
          </div>
        </div>

        <form
          className="flex justify-end mt-4 gap-2"
          onSubmit={handleSendMessage}
        >
          <input
            type="text"
            name="message"
            className=" p-4 border border-gray-700 bg-transparent h-11 rounded-lg w-[90%] outline-none focus:ring-1 bg-primary focus:ring-primary"
            placeholder="Your message..."
          />
          <button
            type="submit"
            className="bg-primary  hover:bg-second font-bold px-4 rounded-lg"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
            </svg>
          </button>
        </form>
      </div>
    </div>
  );
};

const displayMessages = (messages) => {
  // Vérifie si messages est null ou non un tableau
  if (!Array.isArray(messages) || messages.length === 0) {
    return <p>Aucun message à afficher.</p>;
  }
  // Si messages n'est pas vide, mappe sur les messages comme avant
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

export const display = (data, handleUserClick) => {
  return data.map((follower) => {
    return (
      <div
        key={follower.id}
        className=" hover:opacity-60 flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 "
        onClick={() => handleUserClick(follower.id, follower.name)}
      >
        <img
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
};

export default Messages;
// const users = [
//     { id: 1, first_name: 'Mamour', last_name: 'Drame', avatar: 'profilibg.jpg', alt: "profil" },
//     { id: 2, first_name: 'Edouard', last_name: 'Mendy', avatar: 'profilibg.jpg', alt: "profil" },
//     { id: 3, first_name: 'Vincent', last_name: 'Ndour', avatar: "profilibg.jpg", alt: "profil" },
// ];

// const groups = [
//     { id: 1, first_name: 'Call of duty', last_name: '', avatar: 'profilibg.jpg', alt: "profil" },
//     { id: 2, first_name: 'Farcry 6 Team', last_name: '', avatar: 'profilibg.jpg', alt: "profil" },
//     { id: 3, first_name: 'EA Fooball 24', last_name: '', avatar: "profilibg.jpg", alt: "profil" },
// ];

const messages = [
  {
    id: 1,
    first_name: "Mamou Drame",
    message_content: "Hello everyone!",
    timestamp: "2023-11-16T12:00:00.000Z",
  },
  {
    id: 2,
    first_name: "Nicolas Faye",
    message_content: "How are you all doing today?",
    timestamp: "2023-11-16T12:01:00.000Z",
  },
];

const users = [
  { id: 1, name: "Mamour Drame", src: "/assets/profilibg.jpg", alt: "profil" },
  { id: 2, name: "Edouard Mendy", src: "/assets/profilibg.jpg", alt: "profil" },
  { id: 3, name: "Vincent Ndour", src: "/assets/profilibg.jpg", alt: "profil" },
  {
    id: 4,
    name: "Ibrahima Gueye",
    src: "/assets/profilibg.jpg",
    alt: "profil",
  },
  { id: 5, name: "Madike Yade", src: "/assets/profilibg.jpg", alt: "profil" },
];

// const messages = [
//     {
//         id: 1, first_name: 'Mamou Drame', message_content: 'Hello everyone!', timestamp: '2023-11-16T12:00:00.000Z',
//     },
//     {
//         id: 2, first_name: 'Nicolas Faye', message_content: 'How are you all doing today?', timestamp: '2023-11-16T12:01:00.000Z',
//     },
// ];

const groups = [
  { id: 1, name: "Call of duty", src: "/assets/profilibg.jpg", alt: "profil" },
  { id: 2, name: "Farcry 6 Team", src: "/assets/profilibg.jpg", alt: "profil" },
  { id: 3, name: "EA Fooball 24", src: "/assets/profilibg.jpg", alt: "profil" },
];
