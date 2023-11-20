SELECT article_id, work_id, title, content, note, old_works FROM "Article"
    WHERE search @@ websearch_to_tsquery('Japan');

ALTER TABLE "Article" DROP search;
ALTER TABLE "Article" 
ADD search tsvector
    generated always as (
        to_tsvector('english', content)
    ) stored;

CREATE INDEX idx_search on "Article" using GIN(search);