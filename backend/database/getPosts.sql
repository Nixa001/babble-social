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
    p.user_id
