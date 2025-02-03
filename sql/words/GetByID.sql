SELECT word_id,
       word,
       meaning,
       date_created
FROM words
WHERE word_id = ?