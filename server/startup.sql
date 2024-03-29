/*SELECT article_id, work_id, title, content, note, old_works FROM "Article"
    WHERE search @@ websearch_to_tsquery('Japan');

ALTER TABLE "Article" DROP search;
ALTER TABLE "Article" 
ADD search tsvector
    generated always as (
        to_tsvector('english', content)
    ) stored;

CREATE INDEX idx_search on "Article" using GIN(search);*/

-----------------------
-- sudo -u postgres psql postgres
DROP DATABASE marxists;
CREATE DATABASE marxists;

/*SELECT * FROM "Glossary" WHERE author_id is not NULL 
    and search @@ websearch_to_tsquery('David');


ALTER TABLE "Glossary" DROP search;*/
ALTER TABLE "Glossary" 
ADD search tsvector
    generated always as (
        setweight(to_tsvector('simple', name), 'A') 
            || ' ' || setweight(to_tsvector('english',description), 'B') :: tsvector
    ) stored;

CREATE INDEX idx_glossary_search on "Glossary" using GIN(search);



ALTER TABLE "Work"
ADD COLUMN search tsvector;
CREATE INDEX idx_work_search on "Work" using GIN(search);



////////
ALTER TABLE "Work" DROP search;
ALTER TABLE "Work" 
ADD search tsvector as (setweight(to_tsvector('simple', title), 'A') 
            || ' ' || setweight(to_tsvector('english', content), 'B') :: tsvector);


----------------------

SELECT * FROM "Glossary" WHERE author_id is not NULL 
    and (search @@ websearch_to_tsquery('english','David') 
    or search @@ websearch_to_tsquery('simple','David'))
    ORDER by ts_rank(search, websearch_to_tsquery('english','David')) +
    ts_rank(search, websearch_to_tsquery('simple','David')) DESC;



-=-=-=-=-=--=-=-=--=-=-=--=--
UPDATE "Glossary" 
	SET author_id = (SELECT author_id FROM "Author" where name ILIKE '%mao%') 
WHERE name ILIKE '%mao%';
