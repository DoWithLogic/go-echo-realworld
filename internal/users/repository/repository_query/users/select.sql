SELECT 
    u.id,
    u.email,
    u.username,
    u.bio,
    u.image
FROM users u
WHERE u.id = ?
