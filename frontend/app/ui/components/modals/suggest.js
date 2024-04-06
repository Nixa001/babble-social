import React from 'react'

export const Suggest = ({ followers, isVisible, onClose, id_group, sendMessage }) => {
    if (!isVisible) return null;
    return (
        <div className='fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm 
        flex justify-center items-center'

        // onClick={() => onClose()}
        >
            <div className='w-[700px] h-[600px] pb-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
             backdrop-blur-md border border-gray-700 hover:bg-opacity-95' >
                <button className='w-full p-2 flex justify-end'
                    onClick={() => onClose()}>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"
                        className="w-8 h-8 hover:text-red-500 hover:rotate-90 transition duration-300 ease-in-out
                         place-self-end">
                        <path fillRule="evenodd" d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z" clipRule="evenodd" />
                    </svg>

                </button>
                <div className='flex flex-col items-center'>
                    <h1 className='text-2xl text-center font-bold underline underline-offset-8 mb-5'>
                        Suggest to your friends
                    </h1>
                    <div className="flex flex-col lg:w-[100%] 2xl-[80%] xl:w-[75%] w-[80%]  gap-1 ">
                        <input placeholder='search friend ...' className='h-8 bg-transparent border border-gray-700 rounded-md text-center focus:outline-none focus:border-primary' />
                        <div className="flex flex-col h-[400px] overflow-scroll">
                            {displaySuggestFriend(followers, id_group, sendMessage)}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


// const followers = [
//     { name: 'Vindour', src: "/assets/profilibg.jpg", alt: "profil" },
//     { name: 'ibg', src: "/assets/profilibg.jpg", alt: "profil", },
//     { name: 'dicks', src: "/assets/profilibg.jpg", alt: "profil", },
//     { name: 'Vindcour', src: "/assets/profilibg.jpg", alt: "profil" },
//     { name: 'ibgs', src: "/assets/profilibg.jpg", alt: "profil", },
//     { name: 'dickss', src: "/assets/profilibg.jpg", alt: "profil", },
// ];

export const displaySuggestFriend = (data, id_group, sendMessage) => {
    return data.map((follower) => {
        return (
            <div key={follower.name} className=" hover:opacity-90 w-[95%] flex items-center cursor-pointer justify-between gap-2 mt-1 mb-3 p-2  ">
                {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}
                <div className='flex items-center gap-2'>
                    <img
                        className="rounded-full "
                        src={`/assets/${follower.avatar}`}
                        alt={follower.alt}
                        width={40}
                        height={40}
                    />
                    <h4 className="font-bold ">{follower.first_name + " " + follower.last_name}</h4>
                </div>
                <button 
                onClick={()=>suggest(follower.id, id_group, sendMessage)}
                type="submit" className="bg-second h-7 text-lg font-bold px-2 rounded-md cursor-pointer hover:bg-primary">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                        <path d="M5.25 6.375a4.125 4.125 0 1 1 8.25 0 4.125 4.125 0 0 1-8.25 0ZM2.25 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63 13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63l-.001-.122ZM18.75 7.5a.75.75 0 0 0-1.5 0v2.25H15a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H21a.75.75 0 0 0 0-1.5h-2.25V7.5Z" />
                    </svg>
                </button>
            </div>
        );
    })
}

function suggest(id, id_group, sendMessage) {
    // alert(id)
    sendMessage({ type: "SuggestFriend", userId: id, id_group: id_group });
}