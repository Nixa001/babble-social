"use client"
import React from "react";
import { useState } from 'react';
import { FaImage } from "react-icons/fa6";

export const postForm = () => {
    return (
        <form
            className="flex flex-col lg:w-[100%] 2xl-[80%] xl:w-[75%] w-[80%]  gap-1  "
            action=""
            method="POST"
            data-form="post"
            encType="multipart/form-data"
        >
            <TextArea
                label="Post Title"
                name="title_post"
                placeholder="Let's post something"
                required
                defaultValue=""
                onChange={(event) => console.log(event.target.value)} // Handle changes
            />
            <div className="flex items-start justify-end">
                <div className="flex gap-1 flex-wrap mr-2 mt-1 text-sm">
                    <Checkbox label="Tech" value="technologie" name="techno" />
                    <Checkbox label="Sport" value="sport" name="sport" />
                    <Checkbox label="Santé" value="sante" name="sante" />
                    <Checkbox label="Musique" value="musique" name="music" />
                    <Checkbox label="News" value="news" name="news" />
                    <Checkbox label="Other" value="other" name="other" defaultChecked />

                </div>

                {PrivacySelect()}

                <label htmlFor="image_post" className=" cursor-pointer mr-2">
                    <FaImage className="w-7 h-7 " />
                </label>
                <input type="file" name="image_post" id="image_post" hidden />
                <input
                    className="bg-second text-lg font-bold pl-3 pr-3 rounded-lg cursor-pointer hover:bg-primary"
                    type="submit"
                    value="Post"
                />
                {/* < Button text="Log In" onClick={handlePost()} /> */}
            </div>
        </form>
    );
};

function PrivacySelect() {
    const [selectedValue, setSelectedValue] = useState('Public');
    const [showUserList, setShowUserList] = useState(false);
    const [selectedUsers, setSelectedUsers] = useState([]);

    const handleChange = (event) => {
        setSelectedValue(event.target.value);
        setShowUserList(event.target.value === 'Select users');
    };


    const handleUserSelection = (userId) => {
        const isSelected = selectedUsers.includes(userId);
        setSelectedUsers(isSelected ? selectedUsers.filter((id) => id !== userId) : [...selectedUsers, userId]); // Toggle user selection
    };

    return (
        <div className="flex flex-col items-center mr-2">
            <select
                id="privacy-select"
                value={selectedValue}
                onChange={handleChange}
                className="w-32 rounded-md px-2 py-1 font-bold outline-none focus:ring-1 bg-primary focus:ring-primary"
            >
                <option value="Public">Public</option>
                <option value="Private">Private</option>
                <option value="Select users">Select users</option>
            </select>

            {showUserList && (
                <div className="mt-1 max-h-44  w-[300px] p-2 overflow-scroll border rounded-md">
                    <ul className="flex flex-wrap gap-2">
                        {followers.map((user) => (
                            <li key={user.name}>
                                <label className="flex gap-1 cursor-pointer">
                                    <input
                                        type="checkbox"
                                        checked={selectedUsers.includes(user.name)}
                                        onChange={() => handleUserSelection(user.name)}
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
    { name: 'Vindour', src: "/assets/profilibg.jpg", alt: "profil" },
    { name: 'ibg', src: "/assets/profilibg.jpg", alt: "profil", },
    { name: 'dicks', src: "/assets/profilibg.jpg", alt: "profil", },
    { name: 'Vindcour', src: "/assets/profilibg.jpg", alt: "profil" },
    { name: 'ibgs', src: "/assets/profilibg.jpg", alt: "profil", },
    { name: 'dickss', src: "/assets/profilibg.jpg", alt: "profil", },
];

export function TextArea({ label, name, placeholder, required, defaultValue, onChange }) {
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
};


export function Checkbox({ label, value, name, defaultChecked = false }) {
    return (
        <div className="checkbox-container">
            <input type="checkbox" id={label} value={value} name={name} defaultChecked={defaultChecked} />
            <label htmlFor={label}>{label}</label>
        </div>
    );
}

