import React from 'react'
import Image from 'next/image';
import { Checkbox, TextArea } from '../../home/postForm';

export const DisplayComments = ({ isVisible, onClose }) => {
    const Comments = comments
    const handleCommentClick = () => {
        alert("comment");
    };
    if (!isVisible) return null;
    return (
        <div className='fixed z-40 inset-0 bg-bg bg-opacity-10 backdrop-blur-sm 
        flex justify-center items-center'>
            <div className='w-[700px] h-[90%] pb-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
             backdrop-blur-md border border-gray-700 hover:bg-opacity-95' >
                <button className='w-full p-2 flex justify-end'
                    onClick={() => onClose()}>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-8 h-8 hover:text-primary 
                   hover:rotate-90 transition duration-300 ease-in-out place-self-end">
                        <path fillRule="evenodd" d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z" clipRule="evenodd" />
                    </svg>

                </button>
                <div className='h-[95%] flex flex-col '>
                    <h1 className='text-2xl text-center font-bold underline underline-offset-8 mb-5'>
                        Comments
                    </h1>
                    {/* <div className="flex flex-col border border-gray-700 mx-5  gap-2 mb-9">
                        <div className=' flex items-center h-fit cursor-pointer justify-start gap-2 mt-1'>
                            <Image
                                className="rounded-full "
                                src="/assets/profil.jpg"
                                alt="img user"
                                width={35}
                                height={35}
                            />
                            <h4 className="font-bold text-sm ">Masseck</h4>
                        </div>
                        <p className=''>Comment ici</p>
                    </div> */}
                    <div className='h-[90%] overflow-x-scroll '>

                        {printComment(Comments)}
                    </div>


                    <form
                        className="flex items-center justify-center w-[100%] gap-1 
                         "
                        action=""
                        method="POST"
                        data-form="post"
                        encType="multipart/form-data">
                        <input type='text' placeholder='Your comment ...' className='bg-transparent border border-gray-700 w-[80%] h-10 px-3 rounded-lg ' />

                        <label htmlFor="image_post" className=" cursor-pointer mr-2">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-8 h-8">
                                <path fillRule="evenodd" d="M1.5 6a2.25 2.25 0 0 1 2.25-2.25h16.5A2.25 2.25 0 0 1 22.5 6v12a2.25 2.25 0 0 1-2.25 2.25H3.75A2.25 2.25 0 0 1 1.5 18V6ZM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0 0 21 18v-1.94l-2.69-2.689a1.5 1.5 0 0 0-2.12 0l-.88.879.97.97a.75.75 0 1 1-1.06 1.06l-5.16-5.159a1.5 1.5 0 0 0-2.12 0L3 16.061Zm10.125-7.81a1.125 1.125 0 1 1 2.25 0 1.125 1.125 0 0 1-2.25 0Z" clipRule="evenodd" />
                            </svg>

                        </label>
                        <input type="file" name="image_post" id="image_post" hidden />

                        <button type="submit" onClick={handleCommentClick}  className="bg-second h-10 text-lg font-bold pl-3 pr-3 rounded-lg cursor-pointer flex items-center hover:bg-primary">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                                <path d="M3.478 2.404a.75.75 0 0 0-.926.941l2.432 7.905H13.5a.75.75 0 0 1 0 1.5H4.984l-2.432 7.905a.75.75 0 0 0 .926.94 60.519 60.519 0 0 0 18.445-8.986.75.75 0 0 0 0-1.218A60.517 60.517 0 0 0 3.478 2.404Z" />
                            </svg>
                        </button>
                    </form>
                </div>

            </div>
        </div>
    )
}

const comments = [
    { id: 1, name: "Nicolas", avatar: '/assets/profil.jpg', content: "Test", imageComment: "/assets/cod.jpg" },
    { id: 2, name: "Nicolas", avatar: '/assets/profil.jpg', content: "Test2", imageComment: "" },
]

const printComment = (comments) => {
    return comments.map((comment) => {
        const hasImage = comment.imageComment !== ""

        return (
            <div key={comment.id} className="flex flex-col border border-gray-700 mx-5  gap-2 mb-9">
                {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}
                <div className=' flex items-center h-fit cursor-pointer justify-start gap-2 mt-1'>
                    <Image
                        className="rounded-full "
                        src={comment.avatar}
                        alt="img user"
                        width={35}
                        height={35}
                    />
                    <h4 className="font-bold text-sm ">{comment.name}</h4>
                </div>
                <p className=''>{comment.content}</p>
                {hasImage ? (
                    <Image
                        className=" "
                        src={comment.imageComment}
                        alt="img comment"
                        width={300}
                        height={300}
                    />
                ) : (
                    ""
                )

                }
            </div>
        )
    })
}