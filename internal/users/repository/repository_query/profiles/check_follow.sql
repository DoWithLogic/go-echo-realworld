SELECT 
    p.user_id,
    p.follow_user_id
FROM profiles p 
WHERE p.user_id = ?
    AND p.follow_user_id = ?
    AND p.is_active IS NOT FALSE