SELECT review_id,
       title,
       review,
       date_created
FROM reviews
WHERE title LIKE ?
ORDER BY date_created desc