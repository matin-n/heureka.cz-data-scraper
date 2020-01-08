# heureka.cz-data-scraper
Scrapes every listing on heureka.cz

## Scraping Process
1. Iterates through pages 1 - 3434 on [https://obchody.heureka.cz/?f=1](https://obchody.heureka.cz/?f=1)
2. Utilizes traversal via jQuery style CSS selectors to find the listing URL and company URL
3. Scrapes & returns the listing url, company url, and company name

----
## Thanks
* [https://github.com/headzoo/surf](Surf)
