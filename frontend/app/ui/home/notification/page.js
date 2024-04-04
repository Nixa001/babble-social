"use client";
import React, { useState } from "react";
// import Notification from './Notification';
const notif = (type, message) => {
  return (
    <div key={type} className="w-full">
      {/* <div className="w-full flex flex-row bg-black justify-evenly">
        <div>

        {type}
        </div>

        {message}
      </div> */}

      {displayNotification(notifications)}
    </div>
  );
};

export const displayNotification = (notifications, sendMessage) => {
  // const { sendMessage } = useApi();
  return notifications.map((notification) => {
    return (
      <div
        key={notification.id}
        className="  flex flex-col items-start border border-gray-700 rounded  justify-start gap-2 p-1 mt-1 "
      >
        {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}
        <div className="flex">
          {notification.type === "SuggestFriend" ? (notification.from + " Suggests you join this group " +  notification.name_group.toUpperCase()) : ""}
          {notification.type === "JoinGroup" ? (notification.to + " want to join  " +   notification.name_group.toUpperCase() ) : ""}
          {notification.type === "Follow" ? (notification.to + " want to follow you ") : ""}
          {/* <p className="font-semibold ">{  notification.from}</p> */}
        </div>
        <div className="flex gap-1">
          {/* <button>Going</button> */}
          <button
            className="bg-primary hover:bg-gray-600 text-sm text-white py-1 px-2 rounded flex items-center"
            onClick={() => {
              going(notification.id, sendMessage, notification.group_id);
            }}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M19.916 4.626a.75.75 0 0 1 .208 1.04l-9 13.5a.75.75 0 0 1-1.154.114l-6-6a.75.75 0 0 1 1.06-1.06l5.353 5.353 8.493-12.74a.75.75 0 0 1 1.04-.207Z"
                clipRule="evenodd"
              />
            </svg>
            Accept
          </button>
          <button
            onClick={() => {
              notGoing(notification.id, sendMessage, notification.group_id);
            }}
            className="bg-red-400 hover:bg-gray-600 text-sm  text-white py-1 px-2 rounded flex items-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-6 h-6"
            >
              <path
                fillRule="evenodd"
                d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z"
                clipRule="evenodd"
              />
            </svg>
            Decline
          </button>
        </div>
      </div>
    );
    // }
  });
};

const notifications = [
  {
    id: 1,
    type: "JoinGroup",
    from: "Nicolas Faye",
    to: "Madike Yade",
    id_group: 1,
    name_group: "Group1",
  },
  {
    id: 2,
    type: "SuggestFriend",
    from: "Mike",
    to: "Madike Yade",
    id_group: 2,
    name_group: "Group2",
  },
  {
    id: 3,
    type: "SuggestFriend",
    from: "Nicolas Faye",
    to: "Madike Yade",
    id_group: 3,
    name_group: "Group3",
  },
  {
    id: 4,
    type: "Follow",
    from: "Nicolas Faye",
    to: "Madike Yade",
    id_group: 3,
    name_group: "Group3",
  },
];

const Notification = () => {
  return (
    <div className="w-[900px] flex flex-col space-y-2">
      {displayNotification(notifications)}
    </div>
  );
};

export default Notification;

const tabb = {
  From: "mail@a.com",
  Type: "JoinGroup",
  Data: {
    id_group: 1,
    button: "disable",
  },
  To: "Adimine group",
};
