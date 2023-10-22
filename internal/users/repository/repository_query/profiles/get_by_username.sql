SELECT 
    u.id,
    u.username,
    COALESCE(u.bio, ''),
    COALESCE(u.image, '')
FROM users u
WHERE u.username = ?