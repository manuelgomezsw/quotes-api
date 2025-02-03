SELECT author,
       work,
       phrase,
       date_created
FROM quotes
ORDER BY RAND()
LIMIT 1;