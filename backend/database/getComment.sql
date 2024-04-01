-- database: social_network.db
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
    c.post_id = "b01af696-f879-41a1-bfb0-70fa01852138"
ORDER BY c.id
    DESC;
