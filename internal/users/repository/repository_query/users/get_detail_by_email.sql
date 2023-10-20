SELECT 
    u.id,
    u.email,
    u.password,
    u.username,
    IFNULL(u.bio, ''),
    IFNULL(u.image, '')
FROM users u
WHERE u.email = ?