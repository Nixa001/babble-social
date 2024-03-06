export const displayFollowers = (data) => {
    return data.map((follower) => {
        return (
            <div key={follower.name} className="shadow flex items-center cursor-pointer justify-start gap-2 mt-1 mb-3 p-2 ">
                {/* <FaUserGroup className='border rounded-full p-2 w-10 h-10' /> */}

                <Image
                    className="rounded-full "
                    src={follower.src}
                    alt={follower.alt}
                    width={40}
                    height={40}
                />
                <h4 className="font-bold ">{follower.name}</h4>
            </div>
        );
    })
}