import React from 'react'

export const CreateEvent = ({ isVisible, onClose }) => {
    if (!isVisible) return null;
    return (
        <div className='fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm 
        flex justify-center items-center'
        
        // onClick={() => onClose()}
        >
            <div className='w-[700px] pb-5 rounded-lg shadow-2xl bg-bg bg-clip-padding backdrop-filter
             backdrop-blur-md border border-gray-700 hover:bg-opacity-95' >
                <button className='w-full p-2 flex justify-end'
                    onClick={() => onClose()}>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-8 h-8 hover:text-red-500 
                   hover:rotate-90 transition duration-300 ease-in-out place-self-end">
                        <path fillRule="evenodd" d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z" clipRule="evenodd" />
                    </svg>

                </button>
                <div>
                    <h1 className='text-2xl text-center font-bold underline underline-offset-8 mb-5'>
                        Create an Event
                    </h1>
                    <form className='flex flex-col gap-4 px-5'>
                        
                        <textarea placeholder='Event description ...' className='bg-transparent h-[100px] border rounded-md border-gray-700 resize-none
                        focus:outline-none  p-1 focus:ring-1 focus:ring-primary
                        '>
                        </textarea>
                        <label htmlFor='date' className='text-gray-300'>
                           Choose a date
                        </label>
                        <input type='date' name='date' className='border border-gray-700 bg-transparent focus:outline-none p-1 focus:ring-1 focus:ring-primary' />
                        {/* <input type='file' className='bg-transparent' /> */}
                        <input type='submit' className='bg-primary rounded-md border border-gray-700 h-[50px] cursor-pointer hover:bg-second text-lg font-bold '  value={"Create Event"} />

                    </form>


                </div>

            </div>
        </div>
    )
}