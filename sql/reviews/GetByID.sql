SELECT review_id,
       title,
       review,
       date_created
FROM reviews
WHERE review_id = ?