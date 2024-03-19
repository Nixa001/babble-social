-- database: social_networkiTEST.db
SELECT
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
    p.user_id AS post_user_id,
    u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(DISTINCT c.id) AS comment_count,
    GROUP_CONCAT(DISTINCT cat.category) AS categories
FROM
    posts AS p,
    users AS u
    LEFT JOIN comment AS c ON p.id = c.post_id,
    categories AS cat ON p.id = cat.post_id
WHERE
    (
        p.privacy = 'public'
        OR (
            p.privacy = 'private'
            AND (
                p.user_id = 1 --? Post creator
                OR EXISTS (
                    SELECT
                        1
                    FROM
                        users_followers
                    WHERE
                        user_id_followed = p.user_id
                        AND user_id_follower = 1 --?
                )
                OR EXISTS (
                    SELECT
                        1
                    FROM
                        users_followers
                    WHERE
                        user_id_follower = p.user_id
                        AND user_id_followed = 1 --?
                )
            )
        )
        OR (
            p.privacy = 'almost'
            AND EXISTS (
                SELECT
                    1
                FROM
                    viewers
                WHERE
                    user_id = 1 --?
                    AND post_id = p.id
            )
        )
    )
GROUP BY
    p.id,
    p.content,
    p.media,
    p.date,
    p.user_id;


-------getonePost

SELECT
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
    p.user_id AS  author_id,
    u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(DISTINCT c.id) AS comment_count,
    GROUP_CONCAT(DISTINCT cat.category) AS categories,
    CASE
    WHEN p.privacy ="public" THEN "public"
    WHEN p.privacy ="private" THEN (
         SELECT GROUP_CONCAT(user_id_follower) 
            FROM users_followers 
            WHERE user_id_followed = p.user_id
           -- OR user_id_follower = p.user_id  
    )
    WHEN p.privacy ="almost" THEN (
        SELECT GROUP_CONCAT(user_id)
            FROM viewers
            WHERE post_id = p.id
        )
        ELSE NULL
    END AS isPublic
FROM
    posts AS p
LEFT JOIN comment AS c ON p.id = c.post_id,
    categories AS cat ON p.id = cat.post_id,
    users AS u ON u.id =  p.user_id
WHERE p.id = 1;