"use client";
import { websocketProvider } from "@/app/home/page";
import React, { useCallback, useContext } from "react";
import { useState } from "react";
import { FaImage } from "react-icons/fa6";
import { IoSend } from "react-icons/io5";

export const postForm = () => {
  const handlePost = (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    console.log("my data => ", data);
    const options = {
      method: "POST",
      body: data,
    };
    fetch("http://localhost:8080/post", options).then(async (x) => {
      const retrieved = await x.json();
      if (retrieved.type != "success")
        alert(retrieved.StatusCode, retrieved.Msg);
      console.log("response", retrieved);
    });
  };

 // console.log("in postForm");
  return (
    <form
      //className="flex flex-col lg:w-[100%] 2xl-[80%] xl:w-[75%] w-[80%]  gap-1  "
      //action="http://localhost:8000/posts"
      method=""
      data-form="post"
      encType="multipart/form-data"
      onSubmit={handlePost}>
      <TextArea
        label="Post Title"
        name="content"
        placeholder="Let's post something"
        required
        defaultValue=""
      />
      <div className="flex items-start justify-end">
        <div className="flex gap-1 flex-wrap mr-2 mt-1 text-sm">
          <Checkbox label="Tech" value="techno" name="techno" />
          <Checkbox label="Sport" value="sport" name="sport" />
          <Checkbox label="Santé" value="health" name="health" />
          <Checkbox label="Musique" value="music" name="music" />
          <Checkbox label="News" value="news" name="news" />
          <Checkbox label="Other" value="other" name="other" true />
        </div>

        {PrivacySelect()}

        <label htmlFor="image_post" className=" cursor-pointer mr-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-8 h-8">
            <path
              fillRule="evenodd"
              d="M1.5 6a2.25 2.25 0 0 1 2.25-2.25h16.5A2.25 2.25 0 0 1 22.5 6v12a2.25 2.25 0 0 1-2.25 2.25H3.75A2.25 2.25 0 0 1 1.5 18V6ZM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0 0 21 18v-1.94l-2.69-2.689a1.5 1.5 0 0 0-2.12 0l-.88.879.97.97a.75.75 0 1 1-1.06 1.06l-5.16-5.159a1.5 1.5 0 0 0-2.12 0L3 16.061Zm10.125-7.81a1.125 1.125 0 1 1 2.25 0 1.125 1.125 0 0 1-2.25 0Z"
              clipRule="evenodd"
            />
          </svg>
        </label>
        <input type="file" name="image" id="image_post" hidden />
        {/* <input
                    className="bg-second text-lg font-bold pl-3 pr-3 rounded-lg cursor-pointer hover:bg-primary"
                    type="submit"
                    value="Post"
                /> */}
        <button
          type="submit"
          className="bg-second h-full text-lg font-bold pl-3 pr-3 rounded-lg cursor-pointer hover:bg-primary">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-6 h-6">
            <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
          </svg>
        </button>
        {/* < Button text="Log In" onClick={handlePost()} /> */}
      </div>
    </form>
  );
};

function PrivacySelect() {
  const [selectedValue, setSelectedValue] = useState("public");
  const [showUserList, setShowUserList] = useState(false);
  const [selectedUsers, setSelectedUsers] = useState([]);

  const handleChange = (event) => {
    setSelectedValue(event.target.value);
    setShowUserList(event.target.value === "almost");
  };

  const handleUserSelection = (userId) => {
    const isSelected = selectedUsers.includes(userId);
    setSelectedUsers(
      isSelected
        ? selectedUsers.filter((id) => id !== userId)
        : [...selectedUsers, userId]
    ); // Toggle user selection
  };

  return (
    <div className="flex flex-col items-center mr-2">
      <select
        id="privacy-select"
        value={selectedValue}
        name="privacy"
        onChange={handleChange}
        className="w-32 rounded-md px-2 py-1 font-bold outline-none focus:ring-1 bg-primary focus:ring-primary">
        <option value="public">Public</option>
        <option value="Private">Private</option>
        <option value="almost">Select users</option>
      </select>
      <input type="hidden" value={selectedUsers} name="viewers" />
      {showUserList && (
        <div className="mt-1 max-h-44  w-[300px] p-2 overflow-scroll border rounded-md">
          <ul className="flex flex-wrap gap-2">
            {followers.map((user) => (
              <li key={user.name}>
                <label className="flex gap-1 cursor-pointer">
                  <input
                    type="checkbox"
                    name={`view-${user.name}`}
                    //checked={selectedUsers.includes(user.name)}
                    value={user.id}
                    onChange={() => handleUserSelection(user.id)}
                  />
                  {user.name}
                </label>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

const followers = [
  { name: "Vindour", src: "/assets/profilibg.jpg", alt: "profil", id: 3 },
  { name: "ibg", src: "/assets/profilibg.jpg", alt: "profil", id: 2 },
  { name: "daniella", src: "/assets/profilibg.jpg", alt: "profil", id: 5 },
  { name: "Vindcour99", src: "/assets/profilibg.jpg", alt: "profil", id: 99 },
  { name: "nixa", src: "/assets/profilibg.jpg", alt: "profil", id: 4 },
  { name: "dickss", src: "/assets/profilibg.jpg", alt: "profil", id: 1 },
];

export function TextArea({
  label,
  name,
  placeholder,
  required,
  defaultValue,
  onChange,
}) {
  return (
    <div className="mb-2">
      {/* {<label className="block text-sm font-medium text-gray-700 mb-1" htmlFor={name}>{label}</label>} */}
      <textarea
        className="resize-none w-[100%] h-10 border border-gray-700 focus:outline-none focus:border bg-transparent text-text rounded-md p-1 focus:ring-1 focus:border-primary focus:ring-primary"
        name={name}
        placeholder={placeholder}
        required={required}
        defaultValue={defaultValue}
        onChange={onChange}
      />
    </div>
  );
}

export function Checkbox({ label, value, name, defaultChecked = false }) {
  return (
    <div className="checkbox-container">
      <input
        type="checkbox"
        id={label}
        value={value}
        name={name}
        defaultChecked={defaultChecked}
      />
      <label htmlFor={label}>{label}</label>
    </div>
  );
}
