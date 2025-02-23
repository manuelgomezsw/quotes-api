SELECT
    r.review_id,
    r.title,
    LEFT(review, 100) AS review,
    r.author,
    r.source,
    IFNULL(GROUP_CONCAT(t.tag SEPARATOR ','), '') AS tags,
    r.`column`,
    r.date_created
FROM reviews r
         LEFT JOIN quotes.tags t ON r.review_id = t.review_id
WHERE r.title LIKE ?
GROUP BY r.review_id
ORDER BY r.date_created DESC;