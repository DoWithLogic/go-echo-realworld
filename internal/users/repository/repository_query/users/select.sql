SELECT 
    u.id,
    u.email,
    u.username,
    IFNULL(u.bio, ''),
    IFNULL(u.image, '')
FROM users u
WHERE u.id = ?
