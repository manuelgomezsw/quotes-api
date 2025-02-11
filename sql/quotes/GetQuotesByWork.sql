SELECT q.quote_id,
       q.author,
       q.work,
       q.phrase,
       IFNULL(GROUP_CONCAT(t.tag SEPARATOR ','), '') AS tags,
       q.date_created
FROM quotes q
         LEFT JOIN quotes.tags t ON q.quote_id = t.quote_id
WHERE q.work LIKE ?
GROUP BY q.quote_id
ORDER BY q.date_created DESC;