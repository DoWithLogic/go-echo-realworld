UPDATE users SET 
    username = CASE WHEN  ? != '' THEN ? ELSE username END,
    email = CASE WHEN  ? != '' THEN ? ELSE email END,
    password = CASE WHEN  ? != '' THEN ? ELSE password END,
    bio = CASE WHEN  ? != '' THEN ? ELSE bio END,
    image = CASE WHEN  ? != '' THEN ? ELSE image END,
    updated_at = ?,
    updated_by = ?
WHERE  id = ?