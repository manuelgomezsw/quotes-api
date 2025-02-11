SELECT
    r.review_id,
    r.title,
    review,
    r.author,
    r.source,
    IFNULL(GROUP_CONCAT(t.tag SEPARATOR ','), '') AS tags,
    r.date_created
FROM reviews r
         LEFT JOIN quotes.tags t ON r.review_id = t.review_id
WHERE r.review_id = ?
GROUP BY r.review_id
ORDER BY r.date_created DESC;