# 数据说明

> 数据共有6个文件，每个文件中每行都代表一条数据，每行中各列之间以Tab分隔

- papers.txt：每行代表一篇论文，各列信息分别为：paper_id title paper_publish_year conference_id，表示ID为'paper_id'的论文的题目为'title'，并于'paper_publish_year'年发表在'conference_id'这个会议上。
- authors.txt：每行代表一位学者，各列信息分别为：author_id author_name
- conferences.txt：每行代表一个会议，各列信息分别为：conference_id ConferenceName
- affiliations.txt：每行代表一个机构，各列信息分别为：affiliation_id AffiliationName
- paper_author_affiliation.txt：每行代表一条论文、学者、机构的关系，各列信息分别为：paper_id author_id affiliation_id author_sequence，表示ID为'author_id'的人在ID为'affiliation_id'的机构中以第'author_sequence'的作者次序发表了ID为'paper_id'的论文。
- paper_reference.txt：每行代表一条引用关系，各列信息分别为：paper_id reference_id，代表ID为'paper_id'的论文有一条指向ID为'reference_id'的论文的引用记录
