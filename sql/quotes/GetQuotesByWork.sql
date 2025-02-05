SELECT quote_id,
       author,
       work,
       phrase,
       date_created
FROM quotes
WHERE work LIKE ?
ORDER BY date_created DESC