SELECT q.quote_id,
       q.author,
       q.work,
       q.phrase,
       q.date_created
FROM quotes q
         JOIN tags t ON q.quote_id = t.quote_id
WHERE t.tag LIKE ?