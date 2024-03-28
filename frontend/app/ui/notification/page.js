import React, { useState } from "react";
// import Notification from './Notification';
const notif = (type, message) => {
  return (
    <div key={type} className="w-full" >
      <div className="w-full flex flex-row bg-black justify-evenly">
        <div>

        {type}
        </div>

        {message}
      </div>
    </div>
  );
};

const Notification = () => {
  const [notifications, setNotifications] = useState([]);
  const tab = Tabnotifications;
  console.log(tab);

  const addNotification = (type, message, dismissible = true, onDismiss) => {
    setNotifications((prevNotifications) => [
      ...prevNotifications,
      {
        id: Math.random().toString(36).substring(2, 15),
        type,
        message,
        dismissible,
        onDismiss,
      },
    ]);
  };

  const handleDismiss = (id) => {
    setNotifications((prevNotifications) =>
      prevNotifications.filter((notification) => notification.id !== id)
    );
  };

  return (
    <div className="w-[900px] flex flex-col space-y-2">
      {tab.map(
        (notification) =>
          // notif(type=notification.type, notification.type)
          // <notif  key={notification.type} type={notification.type} />
          notif(notification.type, notification.message)

        // dismissible={notification.dismissible}
        // onDismiss={() => handleDismiss(notification.id)} // Pass the notification ID to handleDismiss
      )}
    </div>
  );
};

export default Notification;

const Tabnotifications = [
  {
    id: 1,
    type: "Description event 1",
    message: "/assets/profilibg.jpg",
    dismissible: "profil",
  },
  {
    id: 2,
    type: "Description event 2",
    message: "/assets/profilibg.jpg",
    dismissible: "profil",
  },
  {
    id: 3,
    type: "Description event 3",
    message: "/assets/profilibg.jpg",
    dismissible: "profil",
  },
];



const tabb = {
  "From": "mail@a.com",
  "Type": "JoinGroup",
  "Data": {
      "id_group": 1,
      "button": "disable",
  },
  "To": "Adimine group"
}