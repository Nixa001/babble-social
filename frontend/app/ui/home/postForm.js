"use client"
import React from "react";
import { useState } from 'react';
import { FaImage } from "react-icons/fa6";

export const postForm = () => {
    return (
        <form
            className="flex flex-col gap-1  ml-2 mr-2"
            action=""
            method="POST"
            data-form="post"
            encType="multipart/form-data"
        >
            <input
                className="border focus:outline-none focus:border  text-bg  rounded-md p-1
        focus:ring-1 focus:border-primary focus:ring-primary"
                name="title_post"
                placeholder="Let's post something  ..."
                required
                defaultValue={""}
            />
            <textarea
                className="resize-none border focus:outline-none focus:border bg-transparent text-text rounded-md p-1
        focus:ring-1 focus:border-primary focus:ring-primary"
                name="title_post"
                placeholder="Description"
                required
                defaultValue={""}
            />
            <div className="flex items-start justify-end">
                <div className="flex gap-1 mr-2 mt-1 text-sm">
                    <div>
                        <input
                            type="checkbox"
                            id="check"
                            defaultValue="technologie"
                            name="techno"
                        />
                        <label htmlFor="check">Tech</label>
                    </div>
                    <div>
                        <input
                            type="checkbox"
                            id="checks"
                            defaultValue="sport"
                            name="sport"
                        />
                        <label htmlFor="checks">Sport</label>
                    </div>
                    <div>
                        <input
                            type="checkbox"
                            id="checkS"
                            defaultValue="sante"
                            name="sante"
                        />
                        <label htmlFor="checkS">Sante</label>
                    </div>
                    <div>
                        <input
                            type="checkbox"
                            id="checkM"
                            defaultValue="musique"
                            name="music"
                        />
                        <label htmlFor="checkM">Musique</label>
                    </div>
                    <div>
                        <input
                            type="checkbox"
                            id="checkN"
                            defaultValue="news"
                            name="news"
                        />
                        <label htmlFor="checkN">News</label>
                    </div>
                    <div>
                        <input
                            type="checkbox"
                            id="checko"
                            defaultValue="other"
                            name="other"
                            defaultChecked=""
                        />
                        <label htmlFor="checko">Other</label>
                    </div>
                </div>

                {PrivacySelect()}

                <label htmlFor="image_post" className=" cursor-pointer mr-2">
                    <FaImage className="w-8 h-8 " />
                </label>
                <input type="file" name="image_post" id="image_post" hidden />
                <input
                    className="bg-second text-lg font-bold pl-3 pr-3 rounded-sm cursor-pointer hover:bg-primary"
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


