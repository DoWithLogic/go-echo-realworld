SELECT 
    u.id,
    u.email,
    u.password,
    u.username,
    u.bio,
    u.image
FROM users u
WHERE u.email = ?