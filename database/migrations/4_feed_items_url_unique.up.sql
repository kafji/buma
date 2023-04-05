WITH dupes AS (
  SELECT 
    id, 
    row_number() OVER w as n 
  FROM feed_items WINDOW w AS (
    PARTITION BY url
    ORDER BY id
  )
) DELETE FROM feed_items WHERE id IN (SELECT id FROM dupes WHERE dupes.n > 1);

ALTER TABLE feed_items ADD CONSTRAINT feed_items_url_key UNIQUE (url);
