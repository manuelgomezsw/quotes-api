SELECT quote_id,
       author,
       work,
       phrase,
       date_created
FROM quotes
WHERE author LIKE ?