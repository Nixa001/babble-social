-- database: social_networkiTEST.db
SELECT DISTINCT
    c.id,
    c.content,
    c.media,
    c.date,
    c.post_id,
    u.avatar,
    u.user_name,
    concat (u.first_name, " ", u.last_name) as full_name
FROM
    comment AS c
    LEFT JOIN users AS u on u.id = c.user_id
WHERE
    c.post_id = 1
ORDER BY c.id
    DESC;
