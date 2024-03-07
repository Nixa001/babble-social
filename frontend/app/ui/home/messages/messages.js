'use client'
import Image from 'next/image';

const Messages = () => {

    const handleSendMessage = (e) => {
        e.preventDefault()
    };

    return (
        <div className="md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1200px] w-screen 
     
      flex flex-col sm:flex-row">
            {/* Left sidebar */}
            <div className="w-[100%] sm:w-[30%] h-[100%]  ">
                <div className="flex flex-col items-start justify-center h-full">
                    {/* <div className='flex flex-col w-full bg-black justify-center'> */}
                    <h2 className="text-2xl mb-4 text-center w-full  font-bold">Messages</h2>
                    <div className='flex gap-2 justify-center w-full underline '>
                        <h2 className="text-xl font-semibold mb-4">Users</h2>
                        <h2 className="text-xl font-semibold mb-4">Groups</h2>
                    </div>
                    {/* </div> */}
                    <ul className="list-none">
                        {users.map((user) => (
                            <li key={user.id} className="mb-2 hover:opacity-70 cursor-pointer">
                                <div className="flex items-center">
                                    <Image
                                        src={user.profilePicture}
                                        alt={user.username}
                                        width={40}
                                        height={40}
                                        className="rounded-full"
                                    />
                                    <span className=" font-medium ml-2">{user.username}</span>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>
            {/* Main content area */}
            <div className="flex-grow h-1/2 flex flex-col justify-between border border-gray-700 p-2">
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
                <form className="flex mt-4 gap-2" onSubmit={handleSendMessage}>
                    <input
                        type="text"
                        className=" p-4 text-bg h-11 rounded-lg w-[70%]"
                        placeholder="Your message..."
                    // value={message.content}
                    // onChange={(e) => setMessage(e.target.value)}
                    />
                    <button type="submit" className="bg-primary font-bold px-5 rounded-lg">
                        Send
                    </button>
                </form>
            </div>
        </div>
    );
};

export default Messages;
const users = [
    { id: 1, username: 'Mamou Drame', profilePicture: '/assets/profilibg.jpg' },
    { id: 2, username: 'Edouard Mendy', profilePicture: '/assets/profilibg.jpg' },
];

const messages = [
    {
        id: 1,
        sender: 'Mamou Drame',
        content: 'Hello everyone!',
        timestamp: '2023-11-16T12:00:00.000Z',
    },
    {
        id: 2,
        sender: 'Edouard Mendy',
        content: 'How are you all doing today?',
        timestamp: '2023-11-16T12:01:00.000Z',
    },
];